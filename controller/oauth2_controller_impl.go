package controller

import (
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"net/url"
	"simple-oauth-service/constanta"
	errors2 "simple-oauth-service/errors"
	"simple-oauth-service/helper"
	"simple-oauth-service/model/request"
	"simple-oauth-service/model/response"
	"simple-oauth-service/service"
	"simple-oauth-service/templates"
	"strconv"
)

type OAuth2ControllerImpl struct {
	Oauth2Service service.OAuth2Service
}

func NewOauth2Controller(oauth2Service service.OAuth2Service) OAuth2Controller {
	return &OAuth2ControllerImpl{Oauth2Service: oauth2Service}
}

func (controller *OAuth2ControllerImpl) Authorize(w http.ResponseWriter, r *http.Request) {

	u, err := url.Parse(r.URL.String())
	helper.PanicIfError(err)

	clientId, err := strconv.ParseInt(u.Query()["client_id"][0], 10, 64)
	helper.PanicIfError(err)
	redirectUrl := u.Query()["redirect_url"][0]
	state := u.Query()["state"][0]

	authorizeRequest := request.AuthorizeRequest{
		ClientId:    clientId,
		RedirectUrl: redirectUrl,
	}

	controller.Oauth2Service.Authorize(r.Context(), authorizeRequest)

	loginUrl := fmt.Sprintf("/oauth/login?client_id=%d&redirect_url=%s&state=%s", clientId, redirectUrl, state)
	http.Redirect(w, r, loginUrl, http.StatusPermanentRedirect)
}

func (controller *OAuth2ControllerImpl) Login(w http.ResponseWriter, r *http.Request) {

	_, isValidToken := helper.ValidateRefreshToken(r)

	u, err := url.Parse(r.URL.String())
	helper.PanicIfError(err)

	clientId, err := strconv.ParseInt(u.Query()["client_id"][0], 10, 64)
	helper.PanicIfError(err)
	redirectUrl := u.Query()["redirect_url"][0]
	state := u.Query()["state"][0]

	email := r.FormValue("email")
	password := r.FormValue("password")

	if email == "" && password == "" && !isValidToken {
		t, err2 := template.ParseFS(templates.Templates, "*.gohtml")
		helper.PanicIfError(err2)

		loginUrl := fmt.Sprintf("/oauth/login?client_id=%d&redirect_url=%s&state=%s", clientId, redirectUrl, state)

		err2 = t.ExecuteTemplate(w, "login.gohtml", loginUrl)
		helper.PanicIfError(err2)
		return
	} else if isValidToken {
		http.Redirect(w, r, "/oauth/refresh_token", http.StatusPermanentRedirect)
		return
	}

	credential := request.LoginRequest{
		Email:    email,
		Password: password,
	}

	codeResponse := controller.Oauth2Service.Login(r.Context(), credential)

	callbackUrl := fmt.Sprintf("%s?code=%s&state=%s", redirectUrl, codeResponse.AuthorizationCode, state)

	http.Redirect(w, r, callbackUrl, http.StatusPermanentRedirect)
}

func (controller *OAuth2ControllerImpl) AccessToken(w http.ResponseWriter, r *http.Request) {
	var accessTokenRequest request.AccessTokenRequest
	helper.ReadFromRequestBody(r, &accessTokenRequest)

	accessTokenResponse := controller.Oauth2Service.AccessToken(r.Context(), accessTokenRequest)

	webResponse := response.WebResponse{
		Code:   http.StatusOK,
		Status: constanta.Status200,
		Data:   accessTokenResponse,
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (controller *OAuth2ControllerImpl) RefreshToken(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("jid")
	if err != nil && errors.Is(err, http.ErrNoCookie) {
		panic(errors2.NewUnauthorizedError(constanta.StatusUnauthorized))
	} else if err != nil {
		helper.PanicIfError(err)
	}

	accessTokenResponse, cookie := controller.Oauth2Service.RefreshToken(r.Context(), c)

	c.MaxAge = -1

	webResponse := response.WebResponse{
		Code:   http.StatusOK,
		Status: constanta.Status200,
		Data:   accessTokenResponse,
	}

	helper.WriteToResponseBodyWithCookie(w, cookie, webResponse)
}

func (controller *OAuth2ControllerImpl) RevokeRefreshToken(w http.ResponseWriter, r *http.Request) {
	var revokeRefreshTokenRequest request.RevokeRefreshTokenRequest
	helper.ReadFromRequestBody(r, &revokeRefreshTokenRequest)

	controller.Oauth2Service.RevokeRefreshToken(r.Context(), revokeRefreshTokenRequest)

	webResponse := response.WebResponse{
		Code:   http.StatusOK,
		Status: constanta.Status200,
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (controller *OAuth2ControllerImpl) InternalLogin(w http.ResponseWriter, r *http.Request) {
	_, isValidToken := helper.ValidateRefreshToken(r)

	if r.Method != http.MethodPost && !isValidToken {
		t, err2 := template.ParseFS(templates.Templates, "*.gohtml")
		helper.PanicIfError(err2)

		err2 = t.ExecuteTemplate(w, "login.gohtml", "/login")
		helper.PanicIfError(err2)
		return
	} else if isValidToken {
		http.Redirect(w, r, "/oauth/refresh_token", http.StatusPermanentRedirect)
		return
	}

	credential := request.LoginRequest{
		Email:    r.FormValue("email"),
		Password: r.FormValue("password"),
	}

	accessTokenResponse, cookie := controller.Oauth2Service.InternalLogin(r.Context(), credential)

	webResponse := response.WebResponse{
		Code:   http.StatusOK,
		Status: constanta.Status200,
		Data:   accessTokenResponse,
	}

	helper.WriteToResponseBodyWithCookie(w, cookie, webResponse)
}

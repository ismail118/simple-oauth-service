package controller

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
	"simple-oauth-service/constanta"
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
	var authorizeRequest request.AuthorizeRequest
	helper.ReadFromRequestBody(r, &authorizeRequest)

	controller.Oauth2Service.Authorize(r.Context(), authorizeRequest)

	redirectUrl := fmt.Sprintf("/login/%v/%v", authorizeRequest.ClientId, authorizeRequest.RedirectUrl)

	req, err := http.NewRequest(http.MethodGet, redirectUrl, nil)
	helper.PanicIfError(err)
	req.Header.Set("Content-Type", "application/json")

	http.Redirect(w, req, redirectUrl, http.StatusPermanentRedirect)
}

func (controller *OAuth2ControllerImpl) Login(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	clientId, err := strconv.ParseInt(vars["clientId"], 10, 64)
	helper.PanicIfError(err)
	redirectUrl := vars["redirectUrl"]

	if r.Method != http.MethodPost {
		t, err2 := template.ParseFS(templates.Templates, "*.gohtml")
		helper.PanicIfError(err2)

		err2 = t.ExecuteTemplate(w, "login.gohtml", fmt.Sprintf("/oauth/login/%d/%s", clientId, redirectUrl))
		helper.PanicIfError(err2)
		return
	}

	credential := request.LoginRequest{
		Email:    r.FormValue("email"),
		Password: r.FormValue("password"),
	}

	codeResponse := controller.Oauth2Service.Login(r.Context(), credential)

	webResponse := response.WebResponse{
		Code:   http.StatusPermanentRedirect,
		Status: constanta.Status308,
		Data:   codeResponse,
	}

	body, err := json.Marshal(webResponse)
	helper.PanicIfError(err)

	req, err := http.NewRequest(http.MethodPost, redirectUrl, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	http.Redirect(w, req, redirectUrl, http.StatusPermanentRedirect)
}

func (controller *OAuth2ControllerImpl) AccessToken(w http.ResponseWriter, r *http.Request) {
	var accessTokenRequest request.AccessTokenRequest
	helper.ReadFromRequestBody(r, &accessTokenRequest)

	accessTokenResponse, c := controller.Oauth2Service.AccessToken(r.Context(), accessTokenRequest)

	webResponse := response.WebResponse{
		Code:   http.StatusOK,
		Status: constanta.Status200,
		Data:   accessTokenResponse,
	}

	helper.WriteToResponseBodyWithCookie(w, c, webResponse)
}

func (controller *OAuth2ControllerImpl) RefreshToken(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("jid")
	helper.PanicIfError(err)

	accessTokenResponse, cookie := controller.Oauth2Service.RefreshToken(r.Context(), c)

	webResponse := response.WebResponse{
		Code:   http.StatusOK,
		Status: constanta.Status200,
		Data:   accessTokenResponse,
	}
	context.Background()

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

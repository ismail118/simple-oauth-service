package middleware

import (
	"net/http"
	"simple-oauth-service/constanta"
	"simple-oauth-service/errors"
	"simple-oauth-service/helper"
	"strings"
)

type UrlData struct {
	Method      string
	Url         string
	IsUrlPrefix bool
}

func AuthorizationHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		noNeedAuthUrls := []UrlData{
			{
				Method:      "GET",
				Url:         "/oauth/authorize",
				IsUrlPrefix: false,
			},
			{
				Method:      "POST",
				Url:         "/oauth/access_token",
				IsUrlPrefix: false,
			},
			{
				Method:      "POST",
				Url:         "/oauth/refresh_token",
				IsUrlPrefix: false,
			},
			{
				Method:      "GET",
				Url:         "/oauth/refresh_token",
				IsUrlPrefix: false,
			},
			{
				Method:      "POST",
				Url:         "/api/user",
				IsUrlPrefix: false,
			},
			{
				Method:      "POST",
				Url:         "/api/user/validate",
				IsUrlPrefix: false,
			},
			{
				Method:      "POST",
				Url:         "/login",
				IsUrlPrefix: true,
			},
			{
				Method:      "GET",
				Url:         "/login",
				IsUrlPrefix: true,
			},
			{
				Method:      "POST",
				Url:         "/oauth/login",
				IsUrlPrefix: false,
			},
			{
				Method:      "GET",
				Url:         "/oauth/login",
				IsUrlPrefix: false,
			},
			{
				Method:      "POST",
				Url:         "/test/",
				IsUrlPrefix: true,
			},
			{
				Method:      "GET",
				Url:         "/test/",
				IsUrlPrefix: true,
			},
		}
		if isUrlNoNeedAuth(r, noNeedAuthUrls...) {
			next.ServeHTTP(w, r)
			return
		}

		xAuthorizationHeader := r.Header.Get(constanta.XAuthorizationKey)
		if xAuthorizationHeader == "" {
			helper.PanicIfError(errors.UnauthorizedError{constanta.StatusUnauthorized})
		}

		accessTokenClaim, err := helper.ParseJwtTokenToClaims(xAuthorizationHeader, constanta.SecretKey)
		helper.PanicIfError(err)

		if accessTokenClaim.RegisteredClaims.Subject != constanta.AccessToken {
			panic(errors.NewForbiddenError(constanta.InvalidToken))
		}
		accessTokenClaim.Context.Context = r.Context()

		r = r.WithContext(accessTokenClaim.Context)

		next.ServeHTTP(w, r)
	})
}

func isUrlNoNeedAuth(r *http.Request, listUrl ...UrlData) bool {
	for _, each := range listUrl {
		if !each.IsUrlPrefix {
			if r.URL.Path == each.Url && r.Method == each.Method {
				return true
			}
		} else {
			if strings.HasPrefix(r.URL.Path, each.Url) && r.Method == each.Method {
				return true
			}
		}
	}

	return false
}

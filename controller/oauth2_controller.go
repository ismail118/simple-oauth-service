package controller

import (
	"net/http"
)

type OAuth2Controller interface {
	Authorize(w http.ResponseWriter, r *http.Request)
	Login(w http.ResponseWriter, r *http.Request)
	AccessToken(w http.ResponseWriter, r *http.Request)
	RefreshToken(w http.ResponseWriter, r *http.Request)
	RevokeRefreshToken(w http.ResponseWriter, r *http.Request)
}

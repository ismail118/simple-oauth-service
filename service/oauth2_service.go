package service

import (
	"context"
	"net/http"
	"simple-oauth-service/model/request"
	"simple-oauth-service/model/response"
)

type OAuth2Service interface {
	Authorize(ctx context.Context, request request.AuthorizeRequest)
	Login(ctx context.Context, request request.LoginRequest) response.LoginResponse
	AccessToken(ctx context.Context, request request.AccessTokenRequest) response.AccessTokenResponse
	RefreshToken(ctx context.Context, c *http.Cookie) (response.AccessTokenResponse, *http.Cookie)
	RevokeRefreshToken(ctx context.Context, request request.RevokeRefreshTokenRequest)
	InternalLogin(ctx context.Context, request request.LoginRequest) (response.AccessTokenResponse, *http.Cookie)
}

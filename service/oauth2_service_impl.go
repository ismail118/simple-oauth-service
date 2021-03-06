package service

import (
	"context"
	"database/sql"
	"encoding/json"
	"github.com/go-playground/validator"
	"github.com/go-redis/redis/v8"
	"net/http"
	"simple-oauth-service/constanta"
	ctx2 "simple-oauth-service/ctx"
	"simple-oauth-service/errors"
	errors2 "simple-oauth-service/errors"
	"simple-oauth-service/helper"
	"simple-oauth-service/model/request"
	"simple-oauth-service/model/response"
	"simple-oauth-service/repository"
	"time"
)

type OAuth2ServiceImpl struct {
	OauthRepository repository.Oauth2Repository
	DB              *sql.DB
	RDB             *redis.Client
	Validator       *validator.Validate
}

func NewOauth2Service(oauth2Repository repository.Oauth2Repository, DB *sql.DB, RDB *redis.Client, validator *validator.Validate) OAuth2Service {
	return &OAuth2ServiceImpl{
		OauthRepository: oauth2Repository,
		DB:              DB,
		RDB:             RDB,
		Validator:       validator,
	}
}

func (service *OAuth2ServiceImpl) Authorize(ctx context.Context, request request.AuthorizeRequest) {
	err := service.Validator.Struct(request)
	helper.PanicIfError(err)

	_, err = service.OauthRepository.FindClientById(ctx, service.DB, request.ClientId)
	helper.PanicIfError(err)
}

func (service *OAuth2ServiceImpl) Login(ctx context.Context, request request.LoginRequest) response.LoginResponse {
	err := service.Validator.Struct(request)
	helper.PanicIfError(err)

	user, err := service.OauthRepository.FindUserByEmail(ctx, service.DB, request.Email)
	if err != nil {
		panic(errors.NewValidationErrors(constanta.UserNotFound))
	}

	if !user.IsVerified.Bool {
		panic(errors.NewForbiddenError(constanta.UserInactive))
	}

	if err = helper.CompareHasPassword(user.Password.String, request.Password); err != nil {
		panic(errors.NewValidationErrors(constanta.WrongPassword))
	}

	code := helper.RandStringBytes(12)

	dataContextModel, err := service.OauthRepository.FindDataContextByUserId(ctx, service.DB, user.Id.Int64)
	if err != nil {
		panic(errors.NewNotFoundError(constanta.DataContextNotFound))
	}

	userResponse := helper.ToUserResponse(dataContextModel.User)
	userRoleResponse := helper.ToUserRoleResponse(dataContextModel.UserRole)
	dataScopesResponse := helper.ToDataScopeResponses(dataContextModel.DataScopes)

	dataContext, err := ctx2.NewContext(userResponse, userRoleResponse, dataScopesResponse)
	helper.PanicIfError(err)

	accessTokenClaims := helper.NewMyCustomClaims(dataContext, user.Email.String, constanta.AccessToken, time.Duration(24)*time.Hour)
	refreshTokenClaims := helper.NewMyCustomClaims(dataContext, user.Email.String, constanta.RefreshToken, time.Duration(24*7)*time.Hour)

	accessTokenStr, err := helper.GenerateJwtToken(accessTokenClaims, constanta.SecretKey)
	helper.PanicIfError(err)
	refreshTokenStr, err := helper.GenerateJwtToken(refreshTokenClaims, constanta.SecretKey)
	helper.PanicIfError(err)

	accessTokenResponse := response.AccessAndRefreshTokenResponse{
		AccessToken:  accessTokenStr,
		RefreshToken: refreshTokenStr,
	}

	accessTokenResponseByte, err := json.Marshal(accessTokenResponse)
	helper.PanicIfError(err)

	err = helper.SetWithExp(ctx, service.RDB, code, string(accessTokenResponseByte), 5)
	helper.PanicIfError(err)

	return response.LoginResponse{
		AuthorizationCode: code,
	}
}

func (service *OAuth2ServiceImpl) AccessToken(ctx context.Context, request request.AccessTokenRequest) response.AccessTokenResponse {
	err := service.Validator.Struct(request)
	helper.PanicIfError(err)

	client, err := service.OauthRepository.FindClientById(ctx, service.DB, request.ClientId)
	if err != nil {
		helper.PanicIfError(errors.NewValidationErrors(constanta.ClientNotFound))
	}

	if client.ClientSecret.String != request.ClientSecret {
		helper.PanicIfError(errors.NewUnauthorizedError(constanta.WrongClientSecret))
	}

	accessAndRefreshTokenResponseStr, err := helper.Get(ctx, service.RDB, request.Code)
	helper.PanicIfError(err)

	var accessAndRefreshTokenResponse response.AccessAndRefreshTokenResponse
	err = json.Unmarshal([]byte(accessAndRefreshTokenResponseStr), &accessAndRefreshTokenResponse)
	helper.PanicIfError(err)

	helper.Delete(ctx, service.RDB, request.Code)

	return response.AccessTokenResponse{
		AccessToken:  accessAndRefreshTokenResponse.AccessToken,
		RefreshToken: accessAndRefreshTokenResponse.RefreshToken,
	}
}

func (service *OAuth2ServiceImpl) RefreshToken(ctx context.Context, c *http.Cookie) (response.AccessTokenResponse, *http.Cookie) {
	if c.Value == "" {
		helper.PanicIfError(errors2.NewUnauthorizedError(constanta.StatusUnauthorized))
	}

	refreshTokenClaim, err := helper.ParseJwtTokenToClaims(c.Value, constanta.SecretKey)
	helper.PanicIfError(err)

	c.MaxAge = -1

	if refreshTokenClaim.RegisteredClaims.Subject != constanta.RefreshToken {
		panic(errors.NewForbiddenError(constanta.InvalidToken))
	}

	refreshTokenClaim.Context.Context = context.Background()

	user, err := service.OauthRepository.FindUserById(refreshTokenClaim.Context, service.DB, refreshTokenClaim.Context.User.Id)
	helper.PanicIfError(err)

	if user.TokenVersion.Int64 != refreshTokenClaim.Context.User.TokenVersion {
		panic(errors.NewUnauthorizedError(constanta.InvalidToken))
	}

	accessTokenClaims := helper.NewMyCustomClaims(refreshTokenClaim.Context, user.Email.String, constanta.AccessToken, time.Duration(24)*time.Hour)
	refreshTokenClaims := helper.NewMyCustomClaims(refreshTokenClaim.Context, user.Email.String, constanta.RefreshToken, time.Duration(24*7)*time.Hour)

	accessTokenStr, err := helper.GenerateJwtToken(accessTokenClaims, constanta.SecretKey)
	helper.PanicIfError(err)
	refreshTokenStr, err := helper.GenerateJwtToken(refreshTokenClaims, constanta.SecretKey)
	helper.PanicIfError(err)

	cookie := &http.Cookie{
		Name:  constanta.JID,
		Value: refreshTokenStr,
		Path:  "/",
	}

	return response.AccessTokenResponse{
		AccessToken: accessTokenStr,
	}, cookie
}

func (service *OAuth2ServiceImpl) RevokeRefreshToken(ctx context.Context, request request.RevokeRefreshTokenRequest) {
	err := service.Validator.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user, err := service.OauthRepository.FindUserById(ctx, service.DB, request.UserId)
	if err != nil {
		panic(errors.NewNotFoundError(constanta.UserNotFound))
	}

	helper.IncreaseUserTokenVersion(&user)

	service.OauthRepository.UpdateUserTokenVersion(ctx, tx, request.UserId, user.TokenVersion.Int64)
}

func (service *OAuth2ServiceImpl) InternalLogin(ctx context.Context, request request.LoginRequest) (response.AccessTokenResponse, *http.Cookie) {
	err := service.Validator.Struct(request)
	helper.PanicIfError(err)

	user, err := service.OauthRepository.FindUserByEmail(ctx, service.DB, request.Email)
	if err != nil {
		panic(errors.NewValidationErrors(constanta.UserNotFound))
	}

	if !user.IsVerified.Bool {
		panic(errors.NewForbiddenError(constanta.UserInactive))
	}

	if errs := helper.CompareHasPassword(user.Password.String, request.Password); errs != nil {
		panic(errors.NewValidationErrors(constanta.WrongPassword))
	}

	dataContextModel, err := service.OauthRepository.FindDataContextByUserId(ctx, service.DB, user.Id.Int64)
	if err != nil {
		panic(errors.NewNotFoundError(constanta.DataContextNotFound))
	}

	userResponse := helper.ToUserResponse(dataContextModel.User)
	userRoleResponse := helper.ToUserRoleResponse(dataContextModel.UserRole)
	dataScopesResponse := helper.ToDataScopeResponses(dataContextModel.DataScopes)

	dataContext, err := ctx2.NewContext(userResponse, userRoleResponse, dataScopesResponse)
	helper.PanicIfError(err)

	accessTokenClaims := helper.NewMyCustomClaims(dataContext, user.Email.String, constanta.AccessToken, time.Duration(24)*time.Hour)
	refreshTokenClaims := helper.NewMyCustomClaims(dataContext, user.Email.String, constanta.RefreshToken, time.Duration(24*7)*time.Hour)

	accessTokenStr, err := helper.GenerateJwtToken(accessTokenClaims, constanta.SecretKey)
	helper.PanicIfError(err)
	refreshTokenStr, err := helper.GenerateJwtToken(refreshTokenClaims, constanta.SecretKey)
	helper.PanicIfError(err)

	cookie := &http.Cookie{
		Name:  constanta.JID,
		Value: refreshTokenStr,
		Path:  "/",
	}

	return response.AccessTokenResponse{
		AccessToken: accessTokenStr,
	}, cookie
}

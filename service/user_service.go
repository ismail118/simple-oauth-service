package service

import (
	"simple-oauth-service/ctx"
	"simple-oauth-service/model/request"
	"simple-oauth-service/model/response"
)

type UserService interface {
	FindAll(ctx ctx.Context, roles ...string) []response.UserResponse
	FindById(ctx ctx.Context, userId int64, roles ...string) response.UserResponse
	Create(ctx ctx.Context, request request.UserCreatedRequest, roles ...string) response.UserResponse
	Update(ctx ctx.Context, request request.UserUpdateRequest, roles ...string) response.UserResponse
	Delete(ctx ctx.Context, userId int64, roles ...string)
}

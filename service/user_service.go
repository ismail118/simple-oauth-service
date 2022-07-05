package service

import (
	"context"
	"simple-oauth-service/ctx"
	"simple-oauth-service/model/request"
	"simple-oauth-service/model/response"
)

type UserService interface {
	FindAll(ctx ctx.Context, roles ...string) []response.UserResponse
	FindById(ctx ctx.Context, userId int64, roles ...string) response.UserResponse
	Create(ctx context.Context, request request.UserCreatedRequest) response.MessageResponse
	Update(ctx ctx.Context, request request.UserUpdateRequest, roles ...string) response.UserResponse
	Delete(ctx ctx.Context, userId int64, roles ...string)
	Validate(ctx context.Context, request request.ValidateRequest) response.MessageResponse
	ChangePassword(ctx ctx.Context, request request.ChangePasswordRequest, roles ...string) response.MessageResponse
}

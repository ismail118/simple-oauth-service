package service

import (
	"simple-oauth-service/ctx"
	"simple-oauth-service/model/request"
	"simple-oauth-service/model/response"
)

type UserRoleService interface {
	FindAll(ctx ctx.Context, roles ...string) []response.UserRoleResponse
	FindById(ctx ctx.Context, userRoleId int64, roles ...string) response.UserRoleResponse
	Create(ctx ctx.Context, request request.UserRoleCreateRequest, roles ...string) response.UserRoleResponse
	Update(ctx ctx.Context, request request.UserRoleUpdateRequest, roles ...string) response.UserRoleResponse
	Delete(ctx ctx.Context, userRoleId int64, roles ...string)
}

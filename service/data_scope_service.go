package service

import (
	"simple-oauth-service/ctx"
	"simple-oauth-service/model/request"
	"simple-oauth-service/model/response"
)

type DataScopeService interface {
	FindAll(ctx ctx.Context, roles ...string) []response.DataScopeResponse
	FindById(ctx ctx.Context, dataScopeId int64, roles ...string) response.DataScopeResponse
	Create(ctx ctx.Context, request request.DataScopeCreateRequest, roles ...string) response.DataScopeResponse
	Update(ctx ctx.Context, request request.DataScopeUpdateRequest, roles ...string) response.DataScopeResponse
	Delete(ctx ctx.Context, dataScopeId int64, roles ...string)
}

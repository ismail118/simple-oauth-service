package service

import (
	"simple-oauth-service/ctx"
	"simple-oauth-service/model/request"
	"simple-oauth-service/model/response"
)

type ClientService interface {
	FindAll(ctx ctx.Context, roles ...string) []response.ClientResponse
	FindById(ctx ctx.Context, clientId int64, roles ...string) response.ClientResponse
	Create(ctx ctx.Context, request request.ClientCreateRequest, roles ...string) response.ClientResponse
	Update(ctx ctx.Context, request request.ClientUpdateRequest, roles ...string) response.ClientResponse
	Delete(ctx ctx.Context, clientId int64, roles ...string)
}

package controller

import (
	"github.com/gorilla/mux"
	"net/http"
	"simple-oauth-service/constanta"
	ctx2 "simple-oauth-service/ctx"
	"simple-oauth-service/helper"
	"simple-oauth-service/model/request"
	"simple-oauth-service/model/response"
	"simple-oauth-service/service"
	"strconv"
)

type ClientControllerImpl struct {
	ClientService service.ClientService
}

func NewClientController(clientService service.ClientService) ClientController {
	return &ClientControllerImpl{
		ClientService: clientService,
	}
}

func (controller *ClientControllerImpl) FindAll(w http.ResponseWriter, r *http.Request) {
	ctx, err := ctx2.ToCtxContext(r.Context())
	helper.PanicIfError(err)

	clientResponses := controller.ClientService.FindAll(ctx, constanta.RoleAdmin, constanta.RolePrincipal, constanta.RoleDistributor, constanta.RoleBuyer)

	webResponse := response.WebResponse{
		Code:   http.StatusOK,
		Status: constanta.Status200,
		Data:   clientResponses,
	}

	helper.WriteToResponseBody(w, webResponse)
}
func (controller *ClientControllerImpl) FindById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	clientId, err := strconv.ParseInt(vars[constanta.ClientId], 10, 64)
	helper.PanicIfError(err)

	ctx, err := ctx2.ToCtxContext(r.Context())
	helper.PanicIfError(err)

	clientResponse := controller.ClientService.FindById(ctx, clientId, constanta.RoleAdmin, constanta.RolePrincipal, constanta.RoleDistributor, constanta.RoleBuyer)

	webResponse := response.WebResponse{
		Code:   http.StatusOK,
		Status: constanta.Status200,
		Data:   clientResponse,
	}

	helper.WriteToResponseBody(w, webResponse)
}
func (controller *ClientControllerImpl) Create(w http.ResponseWriter, r *http.Request) {
	var clientCreateRequest request.ClientCreateRequest
	helper.ReadFromRequestBody(r, &clientCreateRequest)

	ctx, err := ctx2.ToCtxContext(r.Context())
	helper.PanicIfError(err)

	clientResponse := controller.ClientService.Create(ctx, clientCreateRequest, constanta.RoleAdmin, constanta.RolePrincipal, constanta.RoleDistributor, constanta.RoleBuyer)

	webResponse := response.WebResponse{
		Code:   http.StatusOK,
		Status: constanta.Status200,
		Data:   clientResponse,
	}

	helper.WriteToResponseBody(w, webResponse)
}
func (controller *ClientControllerImpl) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	clientId, err := strconv.ParseInt(vars[constanta.ClientId], 10, 64)
	helper.PanicIfError(err)

	var clientUpdateRequest request.ClientUpdateRequest
	helper.ReadFromRequestBody(r, &clientUpdateRequest)

	clientUpdateRequest.Id = clientId

	ctx, err := ctx2.ToCtxContext(r.Context())
	helper.PanicIfError(err)

	clientResponse := controller.ClientService.Update(ctx, clientUpdateRequest, constanta.RoleAdmin, constanta.RolePrincipal, constanta.RoleDistributor, constanta.RoleBuyer)

	webResponse := response.WebResponse{
		Code:   http.StatusOK,
		Status: constanta.Status200,
		Data:   clientResponse,
	}

	helper.WriteToResponseBody(w, webResponse)
}
func (controller *ClientControllerImpl) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	clientId, err := strconv.ParseInt(vars[constanta.ClientId], 10, 64)
	helper.PanicIfError(err)

	ctx, err := ctx2.ToCtxContext(r.Context())
	helper.PanicIfError(err)

	controller.ClientService.Delete(ctx, clientId, constanta.RoleAdmin, constanta.RolePrincipal, constanta.RoleDistributor, constanta.RoleBuyer)

	webResponse := response.WebResponse{
		Code:   http.StatusOK,
		Status: constanta.Status200,
	}

	helper.WriteToResponseBody(w, webResponse)
}

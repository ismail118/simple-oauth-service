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

type DataScopeControllerImpl struct {
	DataScopeService service.DataScopeService
}

func NewDataScopeController(dataScopeService service.DataScopeService) DataScopeController {
	return &DataScopeControllerImpl{
		DataScopeService: dataScopeService,
	}
}

func (controller *DataScopeControllerImpl) FindAll(w http.ResponseWriter, r *http.Request) {
	ctx, err := ctx2.ToCtxContext(r.Context())
	helper.PanicIfError(err)

	dataScopeResponses := controller.DataScopeService.FindAll(ctx, constanta.RoleAdmin, constanta.RolePrincipal, constanta.RoleDistributor, constanta.RoleBuyer)

	webResponse := response.WebResponse{
		Code:   http.StatusOK,
		Status: constanta.Status200,
		Data:   dataScopeResponses,
	}

	helper.WriteToResponseBody(w, webResponse)
}
func (controller *DataScopeControllerImpl) FindById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	dataScopeId, err := strconv.ParseInt(vars[constanta.DataScopeId], 10, 64)
	helper.PanicIfError(err)

	ctx, err := ctx2.ToCtxContext(r.Context())
	helper.PanicIfError(err)

	dataScopeResponse := controller.DataScopeService.FindById(ctx, dataScopeId, constanta.RoleAdmin, constanta.RolePrincipal, constanta.RoleDistributor, constanta.RoleBuyer)

	webResponse := response.WebResponse{
		Code:   http.StatusOK,
		Status: constanta.Status200,
		Data:   dataScopeResponse,
	}

	helper.WriteToResponseBody(w, webResponse)
}
func (controller *DataScopeControllerImpl) Create(w http.ResponseWriter, r *http.Request) {
	var dataScopeCreateRequest request.DataScopeCreateRequest
	helper.ReadFromRequestBody(r, &dataScopeCreateRequest)

	ctx, err := ctx2.ToCtxContext(r.Context())
	helper.PanicIfError(err)

	dataScopeResponse := controller.DataScopeService.Create(ctx, dataScopeCreateRequest, constanta.RoleAdmin, constanta.RolePrincipal, constanta.RoleDistributor, constanta.RoleBuyer)

	webResponse := response.WebResponse{
		Code:   http.StatusOK,
		Status: constanta.Status200,
		Data:   dataScopeResponse,
	}

	helper.WriteToResponseBody(w, webResponse)
}
func (controller *DataScopeControllerImpl) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	dataScopeId, err := strconv.ParseInt(vars[constanta.DataScopeId], 10, 64)
	helper.PanicIfError(err)

	var dataScopeUpdateRequest request.DataScopeUpdateRequest
	helper.ReadFromRequestBody(r, &dataScopeUpdateRequest)

	dataScopeUpdateRequest.Id = dataScopeId

	ctx, err := ctx2.ToCtxContext(r.Context())
	helper.PanicIfError(err)

	dataScopeResponse := controller.DataScopeService.Update(ctx, dataScopeUpdateRequest, constanta.RoleAdmin, constanta.RolePrincipal, constanta.RoleDistributor, constanta.RoleBuyer)

	webResponse := response.WebResponse{
		Code:   http.StatusOK,
		Status: constanta.Status200,
		Data:   dataScopeResponse,
	}

	helper.WriteToResponseBody(w, webResponse)
}
func (controller *DataScopeControllerImpl) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	dataScopeId, err := strconv.ParseInt(vars[constanta.DataScopeId], 10, 64)
	helper.PanicIfError(err)

	ctx, err := ctx2.ToCtxContext(r.Context())
	helper.PanicIfError(err)

	controller.DataScopeService.Delete(ctx, dataScopeId, constanta.RoleAdmin, constanta.RolePrincipal, constanta.RoleDistributor, constanta.RoleBuyer)

	webResponse := response.WebResponse{
		Code:   http.StatusOK,
		Status: constanta.Status200,
	}

	helper.WriteToResponseBody(w, webResponse)
}

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

type UserRoleControllerImpl struct {
	UserRoleService service.UserRoleService
}

func NewUserRoleController(userRoleService service.UserRoleService) UserRoleController {
	return &UserRoleControllerImpl{
		UserRoleService: userRoleService,
	}
}

func (controller *UserRoleControllerImpl) FindAll(w http.ResponseWriter, r *http.Request) {
	ctx, err := ctx2.ToCtxContext(r.Context())
	helper.PanicIfError(err)

	userRoleResponses := controller.UserRoleService.FindAll(ctx, constanta.RoleAdmin, constanta.RolePrincipal, constanta.RoleDistributor, constanta.RoleBuyer)

	webResponse := response.WebResponse{
		Code:   http.StatusOK,
		Status: constanta.Status200,
		Data:   userRoleResponses,
	}

	helper.WriteToResponseBody(w, webResponse)
}
func (controller *UserRoleControllerImpl) FindById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	userRoleId, err := strconv.ParseInt(vars["userRoleId"], 10, 64)
	helper.PanicIfError(err)

	ctx, err := ctx2.ToCtxContext(r.Context())
	helper.PanicIfError(err)

	userRoleResponse := controller.UserRoleService.FindById(ctx, userRoleId, constanta.RoleAdmin, constanta.RolePrincipal, constanta.RoleDistributor, constanta.RoleBuyer)

	webResponse := response.WebResponse{
		Code:   http.StatusOK,
		Status: constanta.Status200,
		Data:   userRoleResponse,
	}

	helper.WriteToResponseBody(w, webResponse)
}
func (controller *UserRoleControllerImpl) Create(w http.ResponseWriter, r *http.Request) {
	var userRoleCreateRequest request.UserRoleCreateRequest
	helper.ReadFromRequestBody(r, &userRoleCreateRequest)

	ctx, err := ctx2.ToCtxContext(r.Context())
	helper.PanicIfError(err)

	userRoleResponse := controller.UserRoleService.Create(ctx, userRoleCreateRequest, constanta.RoleAdmin, constanta.RolePrincipal, constanta.RoleDistributor, constanta.RoleBuyer)

	webResponse := response.WebResponse{
		Code:   http.StatusOK,
		Status: constanta.Status200,
		Data:   userRoleResponse,
	}

	helper.WriteToResponseBody(w, webResponse)
}
func (controller *UserRoleControllerImpl) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	userRoleId, err := strconv.ParseInt(vars["userRoleId"], 10, 64)
	helper.PanicIfError(err)

	var userRoleUpdateRequest request.UserRoleUpdateRequest
	helper.ReadFromRequestBody(r, &userRoleUpdateRequest)

	userRoleUpdateRequest.Id = userRoleId

	ctx, err := ctx2.ToCtxContext(r.Context())
	helper.PanicIfError(err)

	userRoleResponse := controller.UserRoleService.Update(ctx, userRoleUpdateRequest, constanta.RoleAdmin, constanta.RolePrincipal, constanta.RoleDistributor, constanta.RoleBuyer)

	webResponse := response.WebResponse{
		Code:   http.StatusOK,
		Status: constanta.Status200,
		Data:   userRoleResponse,
	}

	helper.WriteToResponseBody(w, webResponse)
}
func (controller *UserRoleControllerImpl) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	userRoleId, err := strconv.ParseInt(vars["userRoleId"], 10, 64)
	helper.PanicIfError(err)

	ctx, err := ctx2.ToCtxContext(r.Context())
	helper.PanicIfError(err)

	controller.UserRoleService.Delete(ctx, userRoleId, constanta.RoleAdmin, constanta.RolePrincipal, constanta.RoleDistributor, constanta.RoleBuyer)

	webResponse := response.WebResponse{
		Code:   http.StatusOK,
		Status: constanta.Status200,
	}

	helper.WriteToResponseBody(w, webResponse)
}

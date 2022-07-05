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

type UserControllerImpl struct {
	UserService service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return &UserControllerImpl{
		UserService: userService,
	}
}

func (controller *UserControllerImpl) FindAll(w http.ResponseWriter, r *http.Request) {
	ctx, err := ctx2.ToCtxContext(r.Context())
	helper.PanicIfError(err)

	userResponses := controller.UserService.FindAll(ctx, constanta.RoleAdmin, constanta.RolePrincipal, constanta.RoleDistributor, constanta.RoleDistributor)

	webResponse := response.WebResponse{
		Code:   http.StatusOK,
		Status: constanta.Status200,
		Data:   userResponses,
	}

	helper.WriteToResponseBody(w, webResponse)
}
func (controller *UserControllerImpl) FindById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	userId, err := strconv.ParseInt(vars["userId"], 10, 64)
	helper.PanicIfError(err)

	ctx, err := ctx2.ToCtxContext(r.Context())
	helper.PanicIfError(err)

	userResponse := controller.UserService.FindById(ctx, userId, constanta.RoleAdmin, constanta.RolePrincipal, constanta.RoleDistributor, constanta.RoleDistributor)

	webResponse := response.WebResponse{
		Code:   http.StatusOK,
		Status: constanta.Status200,
		Data:   userResponse,
	}

	helper.WriteToResponseBody(w, webResponse)
}
func (controller *UserControllerImpl) Create(w http.ResponseWriter, r *http.Request) {
	var userCreateRequest request.UserCreatedRequest
	helper.ReadFromRequestBody(r, &userCreateRequest)

	messageResponse := controller.UserService.Create(r.Context(), userCreateRequest)

	webResponse := response.WebResponse{
		Code:   http.StatusOK,
		Status: constanta.Status200,
		Data:   messageResponse,
	}

	helper.WriteToResponseBody(w, webResponse)
}
func (controller *UserControllerImpl) Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	userId, err := strconv.ParseInt(vars["userId"], 10, 64)
	helper.PanicIfError(err)

	var userUpdateRequest request.UserUpdateRequest
	helper.ReadFromRequestBody(r, &userUpdateRequest)

	userUpdateRequest.Id = userId

	ctx, err := ctx2.ToCtxContext(r.Context())
	helper.PanicIfError(err)

	userResponse := controller.UserService.Update(ctx, userUpdateRequest, constanta.RoleAdmin, constanta.RolePrincipal, constanta.RoleDistributor, constanta.RoleDistributor)

	webResponse := response.WebResponse{
		Code:   http.StatusOK,
		Status: constanta.Status200,
		Data:   userResponse,
	}

	helper.WriteToResponseBody(w, webResponse)
}
func (controller *UserControllerImpl) Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	userId, err := strconv.ParseInt(vars["userId"], 10, 64)
	helper.PanicIfError(err)

	ctx, err := ctx2.ToCtxContext(r.Context())
	helper.PanicIfError(err)

	controller.UserService.Delete(ctx, userId, constanta.RoleAdmin, constanta.RolePrincipal, constanta.RoleDistributor, constanta.RoleDistributor)

	webResponse := response.WebResponse{
		Code:   http.StatusOK,
		Status: constanta.Status200,
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (controller *UserControllerImpl) Validate(w http.ResponseWriter, r *http.Request) {
	var validateRequest request.ValidateRequest
	helper.ReadFromRequestBody(r, &validateRequest)

	messageResponse := controller.UserService.Validate(r.Context(), validateRequest)

	webResponse := response.WebResponse{
		Code:   http.StatusOK,
		Status: constanta.Status200,
		Data:   messageResponse,
	}

	helper.WriteToResponseBody(w, webResponse)
}

func (controller *UserControllerImpl) ChangePassword(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	userId, err := strconv.ParseInt(vars["userId"], 10, 64)
	helper.PanicIfError(err)

	var changePasswordRequest request.ChangePasswordRequest
	helper.ReadFromRequestBody(r, &changePasswordRequest)

	changePasswordRequest.Id = userId

	ctx, err := ctx2.ToCtxContext(r.Context())
	helper.PanicIfError(err)

	messageResponse := controller.UserService.ChangePassword(ctx, changePasswordRequest)

	webResponse := response.WebResponse{
		Code:   http.StatusOK,
		Status: constanta.Status200,
		Data:   messageResponse,
	}

	helper.WriteToResponseBody(w, webResponse)
}

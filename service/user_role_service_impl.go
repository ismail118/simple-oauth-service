package service

import (
	"database/sql"
	"github.com/go-playground/validator"
	"simple-oauth-service/constanta"
	"simple-oauth-service/ctx"
	"simple-oauth-service/errors"
	"simple-oauth-service/helper"
	"simple-oauth-service/model/domain"
	"simple-oauth-service/model/request"
	"simple-oauth-service/model/response"
	"simple-oauth-service/repository"
	"time"
)

type UserRoleServiceImpl struct {
	UserRoleRepository repository.UserRoleRepository
	DB                 *sql.DB
	Validator          *validator.Validate
}

func NewUserRoleService(userRoleRepository repository.UserRoleRepository, DB *sql.DB, validator *validator.Validate) UserRoleService {
	return &UserRoleServiceImpl{
		UserRoleRepository: userRoleRepository,
		DB:                 DB,
		Validator:          validator,
	}
}

func (service *UserRoleServiceImpl) FindAll(ctx ctx.Context, roles ...string) []response.UserRoleResponse {
	err := helper.CheckRoles(ctx, roles...)
	helper.PanicIfError(err)

	userRoles := service.UserRoleRepository.FindAll(ctx, service.DB)

	return helper.ToUserRoleResponses(userRoles)
}

func (service *UserRoleServiceImpl) FindById(ctx ctx.Context, userRoleId int64, roles ...string) response.UserRoleResponse {
	err := helper.CheckRoles(ctx, roles...)
	helper.PanicIfError(err)

	userRole, err := service.UserRoleRepository.FindById(ctx, service.DB, userRoleId)
	if err != nil {
		panic(errors.NewNotFoundError(constanta.UserRoleNotFound))
	}

	return helper.ToUserRoleResponse(userRole)
}

func (service *UserRoleServiceImpl) Create(ctx ctx.Context, request request.UserRoleCreateRequest, roles ...string) response.UserRoleResponse {
	err := helper.CheckRoles(ctx, roles...)
	helper.PanicIfError(err)

	err = service.Validator.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	userRole := domain.UserRoleModel{
		Role:      sql.NullString{String: request.Role, Valid: true},
		CreatedAt: sql.NullTime{Time: time.Now(), Valid: true},
	}

	userRole = service.UserRoleRepository.Save(ctx, tx, userRole)

	return helper.ToUserRoleResponse(userRole)
}

func (service *UserRoleServiceImpl) Update(ctx ctx.Context, request request.UserRoleUpdateRequest, roles ...string) response.UserRoleResponse {
	err := helper.CheckRoles(ctx, roles...)
	helper.PanicIfError(err)

	err = service.Validator.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	userRole, err := service.UserRoleRepository.FindById(ctx, service.DB, request.Id)
	if err != nil {
		panic(errors.NewNotFoundError(constanta.UserRoleNotFound))
	}

	userRole = domain.UserRoleModel{
		Id:        sql.NullInt64{Int64: request.Id, Valid: true},
		Role:      sql.NullString{String: request.Role, Valid: true},
		CreatedAt: userRole.CreatedAt,
	}

	userRole = service.UserRoleRepository.Update(ctx, tx, userRole)

	return helper.ToUserRoleResponse(userRole)
}

func (service *UserRoleServiceImpl) Delete(ctx ctx.Context, userRoleId int64, roles ...string) {
	err := helper.CheckRoles(ctx, roles...)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	userRole, err := service.UserRoleRepository.FindById(ctx, service.DB, userRoleId)
	if err != nil {
		panic(errors.NewNotFoundError(constanta.UserRoleNotFound))
	}

	service.UserRoleRepository.Delete(ctx, tx, userRole)
}

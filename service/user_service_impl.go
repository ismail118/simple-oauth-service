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
	"strconv"
	"time"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
	DB             *sql.DB
	Validator      *validator.Validate
}

func NewUserService(userRepository repository.UserRepository, DB *sql.DB, validator *validator.Validate) UserService {
	return &UserServiceImpl{
		UserRepository: userRepository,
		DB:             DB,
		Validator:      validator,
	}
}

func (service *UserServiceImpl) FindAll(ctx ctx.Context, roles ...string) []response.UserResponse {
	err := helper.CheckRoles(ctx, roles...)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	users := service.UserRepository.FindAll(ctx, tx)

	return helper.ToUserResponses(users)
}

func (service *UserServiceImpl) FindById(ctx ctx.Context, userId int64, roles ...string) response.UserResponse {
	err := helper.CheckRoles(ctx, roles...)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user, err := service.UserRepository.FindById(ctx, tx, userId)
	if err != nil {
		panic(errors.NewNotFoundError(constanta.UserNotFound))
	}

	return helper.ToUserResponse(user)
}

func (service *UserServiceImpl) Create(ctx ctx.Context, request request.UserCreatedRequest, roles ...string) response.UserResponse {
	err := helper.CheckRoles(ctx, roles...)
	helper.PanicIfError(err)

	err = service.Validator.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user := domain.UserModel{
		Email:         request.Email,
		Password:      helper.HashAndSalt(request.Password),
		FirstName:     request.FirstName,
		LastName:      request.LastName,
		UserRoleId:    request.UserRoleId,
		CompanyId:     request.CompanyId,
		PrincipalId:   request.PrincipalId,
		DistributorId: request.DistributorId,
		BuyerId:       request.BuyerId,
		TokenVersion:  0,
		IsVerified:    false,
		IsDelete:      false,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
		CreatedBy:     strconv.FormatInt(ctx.User.Id, 10),
		UpdatedBy:     strconv.FormatInt(ctx.User.Id, 10),
	}

	user = service.UserRepository.Save(ctx, tx, user)

	return helper.ToUserResponse(user)
}

func (service *UserServiceImpl) Update(ctx ctx.Context, request request.UserUpdateRequest, roles ...string) response.UserResponse {
	err := helper.CheckRoles(ctx, roles...)
	helper.PanicIfError(err)

	err = service.Validator.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user, err := service.UserRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(errors.NewNotFoundError(constanta.UserNotFound))
	}

	user = domain.UserModel{
		Id:            user.Id,
		Email:         request.Email,
		Password:      user.Password,
		FirstName:     request.FirstName,
		LastName:      request.LastName,
		UserRoleId:    request.UserRoleId,
		CompanyId:     request.CompanyId,
		PrincipalId:   request.PrincipalId,
		DistributorId: request.DistributorId,
		BuyerId:       request.BuyerId,
		TokenVersion:  user.TokenVersion,
		IsVerified:    request.IsVerified,
		IsDelete:      request.IsDelete,
		UpdatedAt:     time.Now(),
		CreatedAt:     user.CreatedAt,
		UpdatedBy:     strconv.FormatInt(ctx.User.Id, 10),
		CreatedBy:     user.CreatedBy,
	}

	user = service.UserRepository.Update(ctx, tx, user)

	return helper.ToUserResponse(user)
}

func (service *UserServiceImpl) Delete(ctx ctx.Context, userId int64, roles ...string) {
	err := helper.CheckRoles(ctx, roles...)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user, err := service.UserRepository.FindById(ctx, tx, userId)
	if err != nil {
		panic(errors.NewNotFoundError(constanta.UserNotFound))
	}

	service.UserRepository.Delete(ctx, tx, user)
}

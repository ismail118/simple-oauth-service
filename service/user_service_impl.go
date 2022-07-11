package service

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/go-playground/validator"
	"github.com/go-redis/redis/v8"
	"gopkg.in/gomail.v2"
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

type UserServiceImpl struct {
	UserRepository     repository.UserRepository
	UserRoleRepository repository.UserRoleRepository
	DB                 *sql.DB
	RDB                *redis.Client
	Validator          *validator.Validate
	Dialer             *gomail.Dialer
}

func NewUserService(userRepository repository.UserRepository, userRoleRepository repository.UserRoleRepository, db *sql.DB, rdb *redis.Client, validator *validator.Validate, dialer *gomail.Dialer) UserService {
	return &UserServiceImpl{
		UserRepository:     userRepository,
		UserRoleRepository: userRoleRepository,
		DB:                 db,
		RDB:                rdb,
		Validator:          validator,
		Dialer:             dialer,
	}
}

func (service *UserServiceImpl) FindAll(ctx ctx.Context, roles ...string) []response.UserResponse {
	err := helper.CheckRoles(ctx, roles...)
	helper.PanicIfError(err)

	users := service.UserRepository.FindAll(ctx, service.DB)

	return helper.ToUserResponses(users)
}

func (service *UserServiceImpl) FindById(ctx ctx.Context, userId int64, roles ...string) response.UserResponse {
	err := helper.CheckRoles(ctx, roles...)
	helper.PanicIfError(err)

	user, err := service.UserRepository.FindById(ctx, service.DB, userId)
	if err != nil {
		panic(errors.NewNotFoundError(constanta.UserNotFound))
	}

	return helper.ToUserResponse(user)
}

func (service *UserServiceImpl) Create(ctx context.Context, request request.UserCreatedRequest) response.MessageResponse {
	var emailVerification string

	err := service.Validator.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	_, err = service.UserRepository.FindByEmail(ctx, service.DB, request.Email)
	if err == nil {
		helper.PanicIfError(errors.NewValidationErrors(constanta.SomeoneAlreadyUseThisEmail))
	}

	userRole, err := service.UserRoleRepository.FindById(ctx, service.DB, request.UserRoleId)
	if err != nil {
		helper.PanicIfError(errors.NewValidationErrors(constanta.UserRoleNotFound))
	}

	if userRole.Role.String == constanta.RoleAdmin {
		emailVerification = service.Dialer.Username
	} else {
		emailVerification = request.Email
	}

	user := domain.UserModel{
		Email:         sql.NullString{String: request.Email},
		Password:      sql.NullString{String: helper.HashAndSalt(request.Password)},
		FirstName:     sql.NullString{String: request.FirstName},
		LastName:      sql.NullString{String: request.LastName},
		UserRoleId:    sql.NullInt64{Int64: request.UserRoleId},
		CompanyId:     sql.NullInt64{Int64: request.CompanyId},
		PrincipalId:   sql.NullInt64{Int64: request.PrincipalId},
		DistributorId: sql.NullInt64{Int64: request.DistributorId},
		BuyerId:       sql.NullInt64{Int64: request.BuyerId},
		TokenVersion:  sql.NullInt64{Int64: 0},
		IsVerified:    sql.NullBool{Bool: false},
		IsDelete:      sql.NullBool{Bool: false},
		CreatedAt:     sql.NullTime{Time: time.Now()},
		UpdatedAt:     sql.NullTime{Time: time.Now()},
		CreatedBy:     sql.NullString{String: request.Email},
		UpdatedBy:     sql.NullString{String: request.Email},
	}

	user = service.UserRepository.Save(ctx, tx, user)

	otp := helper.RandRandomStringNumber(6)

	go helper.SendEmail(service.Dialer, helper.Message{
		From:        service.Dialer.Username,
		To:          []string{emailVerification},
		Subject:     constanta.SubjectUserCreate,
		CC:          "",
		BodyMessage: fmt.Sprintf("Dear user \nyour OTP for registration is %s.\nUse this password to validate your account.", otp),
		FilesAttach: nil,
	})

	err = helper.Set(ctx, service.RDB, otp, request.Email)
	helper.PanicIfError(err)

	return response.MessageResponse{
		Message: fmt.Sprintf("Please check we already send your OTP on email %s", emailVerification),
	}
}

func (service *UserServiceImpl) Update(ctx ctx.Context, request request.UserUpdateRequest, roles ...string) response.UserResponse {
	err := helper.CheckRoles(ctx, roles...)
	helper.PanicIfError(err)

	err = service.Validator.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user, err := service.UserRepository.FindById(ctx, service.DB, request.Id)
	if err != nil {
		panic(errors.NewNotFoundError(constanta.UserNotFound))
	}

	user = domain.UserModel{
		Id:            user.Id,
		Email:         sql.NullString{String: request.Email},
		Password:      user.Password,
		FirstName:     sql.NullString{String: request.FirstName},
		LastName:      sql.NullString{String: request.LastName},
		UserRoleId:    sql.NullInt64{Int64: request.UserRoleId},
		CompanyId:     sql.NullInt64{Int64: request.CompanyId},
		PrincipalId:   sql.NullInt64{Int64: request.PrincipalId},
		DistributorId: sql.NullInt64{Int64: request.DistributorId},
		BuyerId:       sql.NullInt64{Int64: request.BuyerId},
		TokenVersion:  user.TokenVersion,
		IsVerified:    user.IsVerified,
		IsDelete:      sql.NullBool{Bool: request.IsDelete},
		UpdatedAt:     sql.NullTime{Time: time.Now()},
		CreatedAt:     user.CreatedAt,
		UpdatedBy:     sql.NullString{String: ctx.User.Email},
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

	user, err := service.UserRepository.FindById(ctx, service.DB, userId)
	if err != nil {
		panic(errors.NewNotFoundError(constanta.UserNotFound))
	}

	service.UserRepository.Delete(ctx, tx, user)
}

func (service *UserServiceImpl) Validate(ctx context.Context, request request.ValidateRequest) response.MessageResponse {
	err := service.Validator.Struct(request)
	helper.PanicIfError(err)

	email, err := helper.Get(ctx, service.RDB, request.OTP)
	helper.PanicIfError(err)

	if email == "" {
		panic(errors.NewValidationErrors(constanta.InvalidOtp))
	}

	helper.Delete(ctx, service.RDB, request.OTP)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user, err := service.UserRepository.FindByEmail(ctx, service.DB, email)
	helper.PanicIfError(err)

	user = domain.UserModel{
		Id:            user.Id,
		Email:         user.Email,
		Password:      user.Password,
		FirstName:     user.FirstName,
		LastName:      user.LastName,
		UserRoleId:    user.UserRoleId,
		CompanyId:     user.CompanyId,
		PrincipalId:   user.PrincipalId,
		DistributorId: user.DistributorId,
		BuyerId:       user.BuyerId,
		TokenVersion:  user.TokenVersion,
		IsVerified:    sql.NullBool{Bool: true},
		IsDelete:      user.IsDelete,
		UpdatedAt:     sql.NullTime{Time: time.Now()},
		CreatedAt:     user.CreatedAt,
		UpdatedBy:     user.Email,
		CreatedBy:     user.CreatedBy,
	}

	service.UserRepository.Update(ctx, tx, user)

	return response.MessageResponse{
		Message: fmt.Sprintf("Congratulation your account have success verified"),
	}
}

func (service *UserServiceImpl) ChangePassword(ctx ctx.Context, request request.ChangePasswordRequest, roles ...string) response.MessageResponse {
	err := helper.CheckRoles(ctx, roles...)
	helper.PanicIfError(err)

	err = service.Validator.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	user, err := service.UserRepository.FindById(ctx, service.DB, request.Id)
	if err != nil {
		panic(errors.NewNotFoundError(constanta.UserNotFound))
	}

	err = helper.CompareHasPassword(user.Password.String, request.OldPassword)
	if err != nil {
		panic(errors.NewValidationErrors(constanta.WrongPassword))
	}

	user = domain.UserModel{
		Id:            user.Id,
		Email:         user.Email,
		Password:      sql.NullString{String: helper.HashAndSalt(request.NewPassword)},
		FirstName:     user.FirstName,
		LastName:      user.LastName,
		UserRoleId:    user.UserRoleId,
		CompanyId:     user.CompanyId,
		PrincipalId:   user.PrincipalId,
		DistributorId: user.DistributorId,
		BuyerId:       user.BuyerId,
		TokenVersion:  user.TokenVersion,
		IsVerified:    user.IsVerified,
		IsDelete:      user.IsDelete,
		UpdatedAt:     sql.NullTime{Time: time.Now()},
		CreatedAt:     user.CreatedAt,
		UpdatedBy:     user.Email,
		CreatedBy:     user.CreatedBy,
	}

	service.UserRepository.Update(ctx, tx, user)

	return response.MessageResponse{
		Message: fmt.Sprintf("Change password success"),
	}
}

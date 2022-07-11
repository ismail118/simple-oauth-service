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

type DataScopeServiceImpl struct {
	DataScopeRepository repository.DataScopeRepository
	DB                  *sql.DB
	Validator           *validator.Validate
}

func NewDataScopeService(dataScopeRepository repository.DataScopeRepository, DB *sql.DB, validator *validator.Validate) DataScopeService {
	return &DataScopeServiceImpl{
		DataScopeRepository: dataScopeRepository,
		DB:                  DB,
		Validator:           validator,
	}
}

func (service *DataScopeServiceImpl) FindAll(ctx ctx.Context, roles ...string) []response.DataScopeResponse {
	err := helper.CheckRoles(ctx, roles...)
	helper.PanicIfError(err)

	dataScopes := service.DataScopeRepository.FindAll(ctx, service.DB)

	return helper.ToDataScopeResponses(dataScopes)
}

func (service *DataScopeServiceImpl) FindById(ctx ctx.Context, dataScopeId int64, roles ...string) response.DataScopeResponse {
	err := helper.CheckRoles(ctx, roles...)
	helper.PanicIfError(err)

	dataScope, err := service.DataScopeRepository.FindById(ctx, service.DB, dataScopeId)
	if err != nil {
		panic(errors.NewNotFoundError(constanta.DataScopeNotFound))
	}

	return helper.ToDataScopeResponse(dataScope)
}

func (service *DataScopeServiceImpl) Create(ctx ctx.Context, request request.DataScopeCreateRequest, roles ...string) response.DataScopeResponse {
	err := helper.CheckRoles(ctx, roles...)
	helper.PanicIfError(err)

	err = service.Validator.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	dataScope := domain.DataScopeModel{
		PrincipalId:   sql.NullInt64{Int64: request.PrincipalId},
		DistributorId: sql.NullInt64{Int64: request.DistributorId},
		BuyerId:       sql.NullInt64{Int64: request.BuyerId},
		IsDelete:      sql.NullBool{Bool: false},
		CreatedAt:     sql.NullTime{Time: time.Now()},
		UpdatedAt:     sql.NullTime{Time: time.Now()},
		CreatedBy:     sql.NullString{String: ctx.User.Email},
		UpdatedBy:     sql.NullString{String: ctx.User.Email},
	}

	dataScope = service.DataScopeRepository.Save(ctx, tx, dataScope)

	return helper.ToDataScopeResponse(dataScope)
}

func (service *DataScopeServiceImpl) Update(ctx ctx.Context, request request.DataScopeUpdateRequest, roles ...string) response.DataScopeResponse {
	err := helper.CheckRoles(ctx, roles...)
	helper.PanicIfError(err)

	err = service.Validator.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	dataScope, err := service.DataScopeRepository.FindById(ctx, service.DB, request.Id)
	if err != nil {
		panic(errors.NewNotFoundError(constanta.DataScopeNotFound))
	}

	dataScope = domain.DataScopeModel{
		Id:            sql.NullInt64{Int64: request.Id},
		UserId:        sql.NullInt64{Int64: request.UserId},
		PrincipalId:   sql.NullInt64{Int64: request.PrincipalId},
		DistributorId: sql.NullInt64{Int64: request.DistributorId},
		BuyerId:       sql.NullInt64{Int64: request.BuyerId},
		IsDelete:      sql.NullBool{Bool: request.IsDelete},
		CreatedAt:     dataScope.CreatedAt,
		UpdatedAt:     sql.NullTime{Time: time.Now()},
		CreatedBy:     dataScope.CreatedBy,
		UpdatedBy:     sql.NullString{String: ctx.User.Email},
	}

	dataScope = service.DataScopeRepository.Update(ctx, tx, dataScope)

	return helper.ToDataScopeResponse(dataScope)
}

func (service *DataScopeServiceImpl) Delete(ctx ctx.Context, dataScopeId int64, roles ...string) {
	err := helper.CheckRoles(ctx, roles...)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	dataScope, err := service.DataScopeRepository.FindById(ctx, service.DB, dataScopeId)
	if err != nil {
		panic(errors.NewNotFoundError(constanta.DataScopeNotFound))
	}

	service.DataScopeRepository.Delete(ctx, tx, dataScope)
}

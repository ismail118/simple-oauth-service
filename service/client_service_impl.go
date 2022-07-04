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

type ClientServiceImpl struct {
	ClientRepository repository.ClientRepository
	DB               *sql.DB
	Validator        *validator.Validate
}

func NewClientService(clientRepository repository.ClientRepository, DB *sql.DB, validator *validator.Validate) ClientService {
	return &ClientServiceImpl{
		ClientRepository: clientRepository,
		DB:               DB,
		Validator:        validator,
	}
}

func (service *ClientServiceImpl) FindAll(ctx ctx.Context, roles ...string) []response.ClientResponse {
	err := helper.CheckRoles(ctx, roles...)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	clients := service.ClientRepository.FindAll(ctx, tx)

	return helper.ToClientResponses(clients)
}

func (service *ClientServiceImpl) FindById(ctx ctx.Context, clientId int64, roles ...string) response.ClientResponse {
	err := helper.CheckRoles(ctx, roles...)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	client, err := service.ClientRepository.FindById(ctx, tx, clientId)
	if err != nil {
		panic(errors.NewNotFoundError(constanta.ClientNotFound))
	}

	return helper.ToClientResponse(client)
}

func (service *ClientServiceImpl) Create(ctx ctx.Context, request request.ClientCreateRequest, roles ...string) response.ClientResponse {
	err := helper.CheckRoles(ctx, roles...)
	helper.PanicIfError(err)

	err = service.Validator.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	client := domain.ClientModel{
		ApplicationName: request.ApplicationName,
		UserId:          ctx.User.Id,
		ClientSecret:    helper.RandStringBytes(15),
		IsDelete:        false,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
		CreatedBy:       strconv.FormatInt(ctx.User.Id, 10),
		UpdatedBy:       strconv.FormatInt(ctx.User.Id, 10),
	}

	client = service.ClientRepository.Save(ctx, tx, client)

	return helper.ToClientResponse(client)
}

func (service *ClientServiceImpl) Update(ctx ctx.Context, request request.ClientUpdateRequest, roles ...string) response.ClientResponse {
	err := helper.CheckRoles(ctx, roles...)
	helper.PanicIfError(err)

	err = service.Validator.Struct(request)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	client, err := service.ClientRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(errors.NewNotFoundError(constanta.ClientNotFound))
	}

	client = domain.ClientModel{
		Id:              client.Id,
		UserId:          client.UserId,
		ApplicationName: request.ApplicationName,
		ClientSecret:    client.ClientSecret,
		IsDelete:        request.IsDelete,
		CreatedAt:       client.CreatedAt,
		UpdatedAt:       time.Now(),
		CreatedBy:       client.CreatedBy,
		UpdatedBy:       strconv.FormatInt(ctx.User.Id, 10),
	}

	client = service.ClientRepository.Update(ctx, tx, client)

	return helper.ToClientResponse(client)
}

func (service *ClientServiceImpl) Delete(ctx ctx.Context, clientId int64, roles ...string) {
	err := helper.CheckRoles(ctx, roles...)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	client, err := service.ClientRepository.FindById(ctx, tx, clientId)
	if err != nil {
		panic(errors.NewNotFoundError(constanta.ClientNotFound))
	}

	service.ClientRepository.Delete(ctx, tx, client)
}

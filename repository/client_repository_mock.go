package repository

import (
	"context"
	"database/sql"
	"errors"
	"simple-oauth-service/helper"
	"simple-oauth-service/model/domain"
	"time"
)

type ClientRepositoryMock struct {
	Clients map[int64]domain.ClientModel
}

func NewClientRepositoryMock() ClientRepository {
	clientRepositoryMock := &ClientRepositoryMock{Clients: map[int64]domain.ClientModel{
		1: domain.ClientModel{
			Id:              1,
			ApplicationName: "test1",
			ClientSecret:    helper.HashAndSalt("secret123456789"),
			IsDelete:        false,
			CreatedAt:       time.Now(),
			UpdatedAt:       time.Now(),
			CreatedBy:       "test1",
			UpdatedBy:       "test1",
		},
		2: domain.ClientModel{
			Id:              2,
			ApplicationName: "test2",
			ClientSecret:    helper.HashAndSalt("secret123456789"),
			IsDelete:        false,
			CreatedAt:       time.Now(),
			UpdatedAt:       time.Now(),
			CreatedBy:       "test2",
			UpdatedBy:       "test2",
		},
		3: domain.ClientModel{
			Id:              3,
			ApplicationName: "test3",
			ClientSecret:    helper.HashAndSalt("secret123456789"),
			IsDelete:        false,
			CreatedAt:       time.Now(),
			UpdatedAt:       time.Now(),
			CreatedBy:       "test3",
			UpdatedBy:       "test3",
		},
		4: domain.ClientModel{
			Id:              4,
			ApplicationName: "test4",
			ClientSecret:    helper.HashAndSalt("secret123456789"),
			IsDelete:        false,
			CreatedAt:       time.Now(),
			UpdatedAt:       time.Now(),
			CreatedBy:       "test4",
			UpdatedBy:       "test4",
		},
	}}

	return clientRepositoryMock
}
func (repository *ClientRepositoryMock) FindAll(ctx context.Context, tx *sql.Tx) []domain.ClientModel {
	var clients []domain.ClientModel

	for _, data := range repository.Clients {
		clients = append(clients, data)
	}

	return clients
}
func (repository *ClientRepositoryMock) FindById(ctx context.Context, tx *sql.Tx, clientId int64) (domain.ClientModel, error) {
	data, ok := repository.Clients[clientId]
	if ok {
		return data, nil
	} else {
		return data, errors.New("data not found")
	}
}

func (repository *ClientRepositoryMock) Save(ctx context.Context, tx *sql.Tx, client domain.ClientModel) domain.ClientModel {
	newId := int64(len(repository.Clients) + 1)
	client.Id = newId
	repository.Clients[newId] = client

	return client
}
func (repository *ClientRepositoryMock) Update(ctx context.Context, tx *sql.Tx, client domain.ClientModel) domain.ClientModel {
	repository.Clients[client.Id] = client

	return client
}
func (repository *ClientRepositoryMock) Delete(ctx context.Context, tx *sql.Tx, client domain.ClientModel) {
	delete(repository.Clients, client.Id)
}

package repository

import (
	"context"
	"database/sql"
	"errors"
	"simple-oauth-service/helper"
	"simple-oauth-service/model/domain"
	"time"
)

type Oauth2RepositoryMock struct {
	Clients map[int64]domain.ClientModel
	Users   map[int64]domain.UserModel
}

func NewOauth2RepositoryMock() Oauth2Repository {
	clients := map[int64]domain.ClientModel{
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
	}
	users := map[int64]domain.UserModel{
		1: domain.UserModel{
			Id:            1,
			Email:         "test@gmail.com",
			Password:      helper.HashAndSalt("test"),
			FirstName:     "test",
			LastName:      "test",
			UserRoleId:    0,
			CompanyId:     0,
			PrincipalId:   0,
			DistributorId: 0,
			BuyerId:       0,
			IsVerified:    false,
			IsDelete:      false,
			CreatedAt:     time.Now(),
			UpdatedAt:     time.Now(),
			CreatedBy:     "test",
			UpdatedBy:     "test",
		},
	}
	return &Oauth2RepositoryMock{
		Clients: clients,
		Users:   users,
	}
}

func (repository *Oauth2RepositoryMock) FindUserByEmail(ctx context.Context, tx *sql.Tx, email string) (domain.UserModel, error) {
	var user domain.UserModel
	for _, each := range repository.Users {
		if each.Email == email {
			return each, nil
		}
	}
	return user, errors.New("data not found")
}

func (repository *Oauth2RepositoryMock) FindUserById(ctx context.Context, tx *sql.Tx, userId int64) (domain.UserModel, error) {
	panic("do something")
}

func (repository *Oauth2RepositoryMock) FindClientById(ctx context.Context, tx *sql.Tx, clientId int64) (domain.ClientModel, error) {
	data, ok := repository.Clients[clientId]
	if ok {
		return data, nil
	} else {
		return data, errors.New("data not found")
	}
}

func (repository *Oauth2RepositoryMock) FindDataContextByUserId(ctx context.Context, tx *sql.Tx, userId int64) (domain.DataContextModel, error) {
	panic("do something")
}

func (repository *Oauth2RepositoryMock) UpdateUserTokenVersion(ctx context.Context, tx *sql.Tx, userId, tokenVersion int64) {
	panic("do something")
}

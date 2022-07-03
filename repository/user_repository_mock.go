package repository

import (
	"context"
	"database/sql"
	"errors"
	"simple-oauth-service/helper"
	"simple-oauth-service/model/domain"
	"time"
)

type UserRepositoryMock struct {
	Users map[int64]domain.UserModel
}

func NewUserRepositoryMock() UserRepository {
	userRepositoryMock := &UserRepositoryMock{Users: map[int64]domain.UserModel{
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
	}}

	return userRepositoryMock
}
func (repository *UserRepositoryMock) FindAll(ctx context.Context, tx *sql.Tx) []domain.UserModel {
	var users []domain.UserModel

	for _, data := range repository.Users {
		users = append(users, data)
	}

	return users
}
func (repository *UserRepositoryMock) FindById(ctx context.Context, tx *sql.Tx, userId int64) (domain.UserModel, error) {
	data, ok := repository.Users[userId]
	if ok {
		return data, nil
	} else {
		return data, errors.New("data not found")
	}
}

func (repository *UserRepositoryMock) FindByEmail(ctx context.Context, tx *sql.Tx, email string) (domain.UserModel, error) {
	var user domain.UserModel
	for _, each := range repository.Users {
		if each.Email == email {
			return each, nil
		}
	}
	return user, errors.New("data not found")
}

func (repository *UserRepositoryMock) Save(ctx context.Context, tx *sql.Tx, user domain.UserModel) domain.UserModel {
	newId := int64(len(repository.Users) + 1)
	user.Id = newId
	repository.Users[newId] = user

	return user
}
func (repository *UserRepositoryMock) Update(ctx context.Context, tx *sql.Tx, user domain.UserModel) domain.UserModel {
	repository.Users[user.Id] = user

	return user
}
func (repository *UserRepositoryMock) Delete(ctx context.Context, tx *sql.Tx, user domain.UserModel) {
	delete(repository.Users, user.Id)
}

package test

import (
	"context"
	"simple-oauth-service/helper"
	"simple-oauth-service/model/domain"
	"simple-oauth-service/repository"
	"testing"
	"time"
)

func TestSaveUserRepository(t *testing.T) {
	db := NewDBTest()
	userRepository := repository.NewUserRepository()
	tx, err := db.Begin()
	helper.PanicIfError(err)

	user := domain.UserModel{
		Email:         "test",
		Password:      "test",
		FirstName:     "tsets",
		LastName:      "sets",
		UserRoleId:    1,
		CompanyId:     0,
		PrincipalId:   0,
		DistributorId: 0,
		BuyerId:       0,
		TokenVersion:  0,
		IsVerified:    false,
		IsDelete:      false,
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
		CreatedBy:     "TEST",
		UpdatedBy:     "SETS",
	}

	userRepository.Save(context.Background(), tx, user)
	err = tx.Commit()
	helper.PanicIfError(err)
}

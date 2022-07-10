package repository

import (
	"context"
	"database/sql"
	"simple-oauth-service/model/domain"
)

type UserRepository interface {
	FindAll(ctx context.Context, db *sql.DB) []domain.UserModel
	FindById(ctx context.Context, db *sql.DB, userId int64) (domain.UserModel, error)
	FindByEmail(ctx context.Context, db *sql.DB, email string) (domain.UserModel, error)
	Save(ctx context.Context, tx *sql.Tx, user domain.UserModel) domain.UserModel
	Update(ctx context.Context, tx *sql.Tx, user domain.UserModel) domain.UserModel
	Delete(ctx context.Context, tx *sql.Tx, user domain.UserModel)
}

package repository

import (
	"context"
	"database/sql"
	"simple-oauth-service/model/domain"
)

type Oauth2Repository interface {
	FindUserByEmail(ctx context.Context, db *sql.DB, email string) (domain.UserModel, error)
	FindUserById(ctx context.Context, db *sql.DB, userId int64) (domain.UserModel, error)
	FindClientById(ctx context.Context, db *sql.DB, clientId int64) (domain.ClientModel, error)
	FindDataContextByUserId(ctx context.Context, db *sql.DB, userId int64) (domain.DataContextModel, error)
	UpdateUserTokenVersion(ctx context.Context, tx *sql.Tx, userId, tokenVersion int64)
}

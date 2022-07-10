package repository

import (
	"context"
	"database/sql"
	"simple-oauth-service/model/domain"
)

type ClientRepository interface {
	FindAll(ctx context.Context, db *sql.DB) []domain.ClientModel
	FindById(ctx context.Context, db *sql.DB, clientId int64) (domain.ClientModel, error)
	Save(ctx context.Context, tx *sql.Tx, client domain.ClientModel) domain.ClientModel
	Update(ctx context.Context, tx *sql.Tx, client domain.ClientModel) domain.ClientModel
	Delete(ctx context.Context, tx *sql.Tx, client domain.ClientModel)
}

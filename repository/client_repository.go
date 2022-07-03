package repository

import (
	"context"
	"database/sql"
	"simple-oauth-service/model/domain"
)

type ClientRepository interface {
	FindAll(ctx context.Context, tx *sql.Tx) []domain.ClientModel
	FindById(ctx context.Context, tx *sql.Tx, clientId int64) (domain.ClientModel, error)
	Save(ctx context.Context, tx *sql.Tx, client domain.ClientModel) domain.ClientModel
	Update(ctx context.Context, tx *sql.Tx, client domain.ClientModel) domain.ClientModel
	Delete(ctx context.Context, tx *sql.Tx, client domain.ClientModel)
}

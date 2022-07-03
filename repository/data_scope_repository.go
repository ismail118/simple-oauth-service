package repository

import (
	"context"
	"database/sql"
	"simple-oauth-service/model/domain"
)

type DataScopeRepository interface {
	FindAll(ctx context.Context, tx *sql.Tx) []domain.DataScopeModel
	FindById(ctx context.Context, tx *sql.Tx, dataScopeId int64) (domain.DataScopeModel, error)
	Save(ctx context.Context, tx *sql.Tx, dataScope domain.DataScopeModel) domain.DataScopeModel
	Update(ctx context.Context, tx *sql.Tx, dataScope domain.DataScopeModel) domain.DataScopeModel
	Delete(ctx context.Context, tx *sql.Tx, dataScope domain.DataScopeModel)
}

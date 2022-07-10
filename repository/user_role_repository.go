package repository

import (
	"context"
	"database/sql"
	"simple-oauth-service/model/domain"
)

type UserRoleRepository interface {
	FindAll(ctx context.Context, db *sql.DB) []domain.UserRoleModel
	FindById(ctx context.Context, db *sql.DB, userRoleId int64) (domain.UserRoleModel, error)
	Save(ctx context.Context, tx *sql.Tx, userRole domain.UserRoleModel) domain.UserRoleModel
	Update(ctx context.Context, tx *sql.Tx, userRole domain.UserRoleModel) domain.UserRoleModel
	Delete(ctx context.Context, tx *sql.Tx, userRole domain.UserRoleModel)
}

package repository

import (
	"context"
	"database/sql"
	"errors"
	"simple-oauth-service/helper"
	"simple-oauth-service/model/domain"
)

type DataScopeRepositoryImpl struct {
}

func NewDataScopeRepository() DataScopeRepository {
	return &DataScopeRepositoryImpl{}
}

func (repository *DataScopeRepositoryImpl) FindAll(ctx context.Context, db *sql.DB) []domain.DataScopeModel {
	querySql := "SELECT id, user_id, principal_id, distributor_id, buyer_id, is_delete, created_at, updated_at, created_by, updated_by FROM data_scope"

	conn, err := db.Conn(ctx)
	rows, err := conn.QueryContext(ctx, querySql)
	helper.PanicIfError(err)
	defer rows.Close()

	var dataScopes []domain.DataScopeModel

	for rows.Next() {
		var dataScope domain.DataScopeModel
		rows.Scan(
			&dataScope.Id,
			&dataScope.UserId,
			&dataScope.PrincipalId,
			&dataScope.DistributorId,
			&dataScope.BuyerId,
			&dataScope.IsDelete,
			&dataScope.CreatedAt,
			&dataScope.UpdatedAt,
			&dataScope.CreatedBy,
			&dataScope.UpdatedBy,
		)

		dataScopes = append(dataScopes, dataScope)
	}

	return dataScopes
}
func (repository *DataScopeRepositoryImpl) FindById(ctx context.Context, db *sql.DB, dataScopeId int64) (domain.DataScopeModel, error) {
	querySql := "SELECT id, user_id, principal_id, distributor_id, buyer_id, is_delete, created_at, updated_at, created_by, updated_by FROM data_scope WHERE id = ?"

	conn, err := db.Conn(ctx)
	rows, err := conn.QueryContext(ctx, querySql, dataScopeId)
	helper.PanicIfError(err)
	defer rows.Close()

	var dataScope domain.DataScopeModel

	if rows.Next() {
		rows.Scan(
			&dataScope.Id,
			&dataScope.UserId,
			&dataScope.PrincipalId,
			&dataScope.DistributorId,
			&dataScope.BuyerId,
			&dataScope.IsDelete,
			&dataScope.CreatedAt,
			&dataScope.UpdatedAt,
			&dataScope.CreatedBy,
			&dataScope.UpdatedBy,
		)
	} else {
		return dataScope, errors.New("data not found")
	}

	return dataScope, nil
}
func (repository *DataScopeRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, dataScope domain.DataScopeModel) domain.DataScopeModel {
	querySql := "INSERT INTO data_scope(user_id, principal_id, distributor_id, buyer_id, is_delete, created_at, updated_at, created_by, updated_by) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)"

	result, err := tx.ExecContext(ctx, querySql,
		dataScope.UserId,
		dataScope.PrincipalId,
		dataScope.DistributorId,
		dataScope.BuyerId,
		dataScope.IsDelete,
		dataScope.CreatedAt,
		dataScope.UpdatedAt,
		dataScope.CreatedBy,
		dataScope.UpdatedBy,
	)

	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	dataScope.Id = id
	return dataScope
}
func (repository *DataScopeRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, dataScope domain.DataScopeModel) domain.DataScopeModel {
	querySql := "UPDATE data_scope SET user_id = ?, principal_id = ?, distributor_id = ?, buyer_id = ?, is_delete = ?, created_at = ?, updated_at = ?, created_by = ?, updated_by = ? WHERE id = ?"

	_, err := tx.ExecContext(ctx, querySql,
		dataScope.UserId,
		dataScope.PrincipalId,
		dataScope.DistributorId,
		dataScope.BuyerId,
		dataScope.IsDelete,
		dataScope.CreatedAt,
		dataScope.UpdatedAt,
		dataScope.CreatedBy,
		dataScope.UpdatedBy,
		dataScope.Id,
	)
	helper.PanicIfError(err)

	return dataScope
}

func (repository *DataScopeRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, dataScope domain.DataScopeModel) {
	querySql := "DELETE FROM data_scope WHERE id = ?"

	_, err := tx.ExecContext(ctx, querySql, dataScope.Id)
	helper.PanicIfError(err)
}

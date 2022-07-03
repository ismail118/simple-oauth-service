package repository

import (
	"context"
	"database/sql"
	"errors"
	"simple-oauth-service/helper"
	"simple-oauth-service/model/domain"
)

type ClientRepositoryImpl struct {
}

func NewClientRepository() ClientRepository {
	return &ClientRepositoryImpl{}
}

func (repository *ClientRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, clientId int64) (domain.ClientModel, error) {
	querySql := "SELECT id, user_id, application_name, client_secret, is_delete, created_at, updated_at, created_by, updated_by FROM client WHERE id = ?"
	rows, err := tx.QueryContext(ctx, querySql, clientId)
	helper.PanicIfError(err)
	defer rows.Close()

	var client domain.ClientModel

	if rows.Next() {
		err = rows.Scan(
			&client.Id,
			&client.UserId,
			&client.ApplicationName,
			&client.ClientSecret,
			&client.IsDelete,
			&client.CreatedAt,
			&client.UpdatedAt,
			&client.CreatedBy,
			&client.UpdatedBy,
		)
		if err != nil {
			panic(err)
		}
		return client, nil
	} else {
		return client, errors.New("data not found")
	}
}

func (repository *ClientRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.ClientModel {
	querySql := "SELECT id, user_id, application_name, client_secret, is_delete, created_at, updated_at, created_by, updated_by FROM client"
	rows, err := tx.QueryContext(ctx, querySql)
	helper.PanicIfError(err)
	defer rows.Close()

	var clients []domain.ClientModel

	for rows.Next() {
		var client domain.ClientModel
		err = rows.Scan(
			&client.Id,
			&client.UserId,
			&client.ApplicationName,
			&client.ClientSecret,
			&client.IsDelete,
			&client.CreatedAt,
			&client.UpdatedAt,
			&client.CreatedBy,
			&client.UpdatedBy,
		)
		if err != nil {
			panic(err)
		}

		clients = append(clients, client)
	}

	return clients
}

func (repository *ClientRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, client domain.ClientModel) domain.ClientModel {
	querySql := "INSERT INTO client(user_id, application_name, client_secret, is_delete, created_at, updated_at, created_by, updated_by) VALUES (?, ?, ?, ?, ?, ?, ?, ?)"
	result, err := tx.ExecContext(ctx, querySql,
		client.UserId,
		client.ApplicationName,
		client.ClientSecret,
		client.IsDelete,
		client.CreatedAt,
		client.UpdatedAt,
		client.CreatedBy,
		client.UpdatedBy,
	)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	client.Id = id
	return client
}

func (repository *ClientRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, client domain.ClientModel) domain.ClientModel {
	querySql := "UPDATE client SET application_name = ?, client_secret = ?, is_delete = ?, updated_at = ?, updated_by = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, querySql,
		client.ApplicationName,
		client.ClientSecret,
		client.IsDelete,
		client.UpdatedAt,
		client.UpdatedBy,
		client.Id,
	)
	helper.PanicIfError(err)
	return client
}

func (repository *ClientRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, client domain.ClientModel) {
	querySql := "DELETE FROM client WHERE id = ?"
	_, err := tx.ExecContext(ctx, querySql, client.Id)
	helper.PanicIfError(err)
}

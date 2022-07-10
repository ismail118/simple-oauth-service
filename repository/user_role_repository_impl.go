package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"simple-oauth-service/helper"
	"simple-oauth-service/model/domain"
)

type UserRoleRepositoryImpl struct {
}

func NewUserRoleRepository() UserRoleRepository {
	return &UserRoleRepositoryImpl{}
}

func (repository *UserRoleRepositoryImpl) FindAll(ctx context.Context, db *sql.DB) []domain.UserRoleModel {
	querySql := "SELECT id, role, created_at FROM user_role"

	conn, err := db.Conn(ctx)
	rows, err := conn.QueryContext(ctx, querySql)
	defer rows.Close()
	helper.PanicIfError(err)

	var userRoles []domain.UserRoleModel

	for rows.Next() {
		var userRole domain.UserRoleModel
		rows.Scan(
			&userRole.Id,
			&userRole.Role,
			&userRole.CreatedAt,
		)
		userRoles = append(userRoles, userRole)
	}

	return userRoles
}

func (repository *UserRoleRepositoryImpl) FindById(ctx context.Context, db *sql.DB, userRoleId int64) (domain.UserRoleModel, error) {
	querySql := "SELECT id, role, created_at FROM user_role WHERE id = ?"

	conn, err := db.Conn(ctx)
	row, err := conn.QueryContext(ctx, querySql, userRoleId)
	helper.PanicIfError(err)

	var userRole domain.UserRoleModel

	if row.Next() {
		row.Scan(
			&userRole.Id,
			&userRole.Role,
			&userRole.CreatedAt,
		)
	} else {
		return userRole, errors.New(fmt.Sprint("data not found"))
	}

	return userRole, nil
}

func (repository *UserRoleRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, userRole domain.UserRoleModel) domain.UserRoleModel {
	querySql := "INSERT INTO user_role(role, created_at) VALUES (?, ?)"
	result, err := tx.ExecContext(ctx, querySql,
		userRole.Role,
		userRole.CreatedAt,
	)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	userRole.Id = id
	return userRole
}

func (repository *UserRoleRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, userRole domain.UserRoleModel) domain.UserRoleModel {
	querySql := "UPDATE user_role SET role = ?, created_at = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, querySql,
		userRole.Role,
		userRole.CreatedAt,
		userRole.Id,
	)
	helper.PanicIfError(err)

	return userRole
}

func (repository *UserRoleRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, userRole domain.UserRoleModel) {
	querySql := "DELETE FROM user_role WHERE id = ?"
	_, err := tx.ExecContext(ctx, querySql, userRole.Id)
	helper.PanicIfError(err)
}

package repository

import (
	"context"
	"database/sql"
	"errors"
	"simple-oauth-service/helper"
	"simple-oauth-service/model/domain"
)

type UserRepositoryImpl struct {
}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

func (repository *UserRepositoryImpl) FindAll(ctx context.Context, tx *sql.DB) []domain.UserModel {
	querySql := "SELECT id, email, password, first_name, last_name, user_role_id, company_id, principal_id, distributor_id, buyer_id, is_verified, is_delete, created_at, updated_at, created_by, updated_by FROM user"

	rows, err := tx.QueryContext(ctx, querySql)
	helper.PanicIfError(err)
	defer rows.Close()

	var users []domain.UserModel

	for rows.Next() {
		var user domain.UserModel
		rows.Scan(
			&user.Id,
			&user.Email,
			&user.Password,
			&user.FirstName,
			&user.LastName,
			&user.UserRoleId,
			&user.CompanyId,
			&user.PrincipalId,
			&user.DistributorId,
			&user.BuyerId,
			&user.IsVerified,
			&user.IsDelete,
			&user.CreatedAt,
			&user.UpdatedAt,
			&user.CreatedBy,
			&user.UpdatedBy,
		)

		users = append(users, user)
	}

	return users
}

func (repository *UserRepositoryImpl) FindById(ctx context.Context, db *sql.DB, userId int64) (domain.UserModel, error) {
	querySql := "SELECT id, email, password, first_name, last_name, user_role_id, company_id, principal_id, distributor_id, buyer_id, is_verified, is_delete, created_at, updated_at, created_by, updated_by FROM user WHERE id = ?"

	conn, err := db.Conn(ctx)
	helper.PanicIfError(err)
	rows, err := conn.QueryContext(ctx, querySql, userId)
	helper.PanicIfError(err)
	defer rows.Close()

	var user domain.UserModel

	if rows.Next() {
		rows.Scan(
			&user.Id,
			&user.Email,
			&user.Password,
			&user.FirstName,
			&user.LastName,
			&user.UserRoleId,
			&user.CompanyId,
			&user.PrincipalId,
			&user.DistributorId,
			&user.BuyerId,
			&user.IsVerified,
			&user.IsDelete,
			&user.CreatedAt,
			&user.UpdatedAt,
			&user.CreatedBy,
			&user.UpdatedBy,
		)

	} else {
		return user, errors.New("data not found")
	}

	return user, nil
}

func (repository *UserRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, user domain.UserModel) domain.UserModel {
	querySql := "INSERT INTO user(email, password, first_name, last_name, user_role_id, company_id, principal_id, distributor_id, buyer_id, token_version, is_verified, is_delete, created_at, updated_at, created_by, updated_by) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"

	result, err := tx.ExecContext(ctx, querySql,
		user.Email,
		user.Password,
		user.FirstName,
		user.LastName,
		user.UserRoleId,
		user.CompanyId,
		user.PrincipalId,
		user.DistributorId,
		user.BuyerId,
		user.TokenVersion,
		user.IsVerified,
		user.IsDelete,
		user.CreatedAt,
		user.UpdatedAt,
		user.CreatedBy,
		user.UpdatedBy,
	)
	helper.PanicIfError(err)
	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	user.Id.Int64 = id
	return user
}

func (repository *UserRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, user domain.UserModel) domain.UserModel {
	querySql := "UPDATE user SET email = ?, password = ?, first_name = ?, last_name = ?, user_role_id = ?, company_id = ?, principal_id = ?, distributor_id = ?, buyer_id = ?, token_version = ?, is_verified = ?, is_delete = ?, created_at = ?, updated_at = ?, created_by = ?, updated_by = ? WHERE id = ?"

	_, err := tx.ExecContext(ctx, querySql,
		user.Email,
		user.Password,
		user.FirstName,
		user.LastName,
		user.UserRoleId,
		user.CompanyId,
		user.PrincipalId,
		user.DistributorId,
		user.BuyerId,
		user.TokenVersion,
		user.IsVerified,
		user.IsDelete,
		user.CreatedAt,
		user.UpdatedAt,
		user.CreatedBy,
		user.UpdatedBy,
		user.Id,
	)
	helper.PanicIfError(err)

	return user
}

func (repository *UserRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, user domain.UserModel) {
	querySql := "DELETE FROM user WHERE id = ?"

	_, err := tx.ExecContext(ctx, querySql, user.Id)
	helper.PanicIfError(err)
}

func (repository *UserRepositoryImpl) FindByEmail(ctx context.Context, db *sql.DB, email string) (domain.UserModel, error) {
	querySql := "SELECT id, email, password, first_name, last_name, user_role_id, company_id, principal_id, distributor_id, buyer_id, token_version, is_verified, is_delete, created_at, updated_at, created_by, updated_by FROM user WHERE email = ?"

	conn, err := db.Conn(ctx)
	helper.PanicIfError(err)
	rows, err := conn.QueryContext(ctx, querySql, email)
	helper.PanicIfError(err)
	defer rows.Close()

	var user domain.UserModel

	if rows.Next() {
		rows.Scan(
			&user.Id,
			&user.Email,
			&user.Password,
			&user.FirstName,
			&user.LastName,
			&user.UserRoleId,
			&user.CompanyId,
			&user.PrincipalId,
			&user.DistributorId,
			&user.BuyerId,
			&user.TokenVersion,
			&user.IsVerified,
			&user.IsDelete,
			&user.CreatedAt,
			&user.UpdatedAt,
			&user.CreatedBy,
			&user.UpdatedBy,
		)

	} else {
		return user, errors.New("data not found")
	}

	return user, nil
}

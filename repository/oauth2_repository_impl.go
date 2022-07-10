package repository

import (
	"context"
	"database/sql"
	"errors"
	"simple-oauth-service/helper"
	"simple-oauth-service/model/domain"
)

type Oauth2RepositoryImpl struct {
}

func NewOauth2Repository() Oauth2Repository {
	return &Oauth2RepositoryImpl{}
}

func (repository *Oauth2RepositoryImpl) FindUserByEmail(ctx context.Context, db *sql.DB, email string) (domain.UserModel, error) {
	querySql := "SELECT id, email, password, first_name, last_name, user_role_id, company_id, principal_id, distributor_id, buyer_id, token_version, is_verified, is_delete, created_at, updated_at, created_by, updated_by FROM user WHERE email = ?"

	conn, err := db.Conn(ctx)
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

func (repository *Oauth2RepositoryImpl) FindUserById(ctx context.Context, db *sql.DB, userId int64) (domain.UserModel, error) {
	querySql := "SELECT id, email, password, first_name, last_name, user_role_id, company_id, principal_id, distributor_id, buyer_id, token_version, is_verified, is_delete, created_at, updated_at, created_by, updated_by FROM user WHERE id = ?"

	conn, err := db.Conn(ctx)
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

func (repository *Oauth2RepositoryImpl) FindClientById(ctx context.Context, db *sql.DB, clientId int64) (domain.ClientModel, error) {
	querySql := "SELECT id, user_id, application_name, client_secret, is_delete, created_at, updated_at, created_by, updated_by FROM client WHERE id = ?"

	conn, err := db.Conn(ctx)
	rows, err := conn.QueryContext(ctx, querySql, clientId)
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

func (repository *Oauth2RepositoryImpl) FindDataContextByUserId(ctx context.Context, db *sql.DB, userId int64) (domain.DataContextModel, error) {
	querySql := "SELECT a.id, a.email, a.password, a.first_name, a.last_name, a.user_role_id, a.company_id, a.principal_id, a.distributor_id, a.buyer_id, a.token_version, a.is_verified, a.is_delete, a.created_at, a.updated_at, a.created_by, a.updated_by " +
		" b.id, b.role, b.created_at " +
		" c.id, c.user_id, c.principal_id, c.distributor_id, c.buyer_id, c.is_delete, c.created_at, c.updated_at, c.created_by, c.updated_by " +
		" FROM user AS a " +
		" JOIN user_role AS b ON a.user_role_id = b.id " +
		" JOIN data_scope AS c ON a.id = c.user_id " +
		" WHERE a.id = ? "

	conn, err := db.Conn(ctx)
	rows, err := conn.QueryContext(ctx, querySql, userId)
	helper.PanicIfError(err)

	var user domain.UserModel
	var userRole domain.UserRoleModel
	var dataScopes []domain.DataScopeModel

	if rows.Next() {
		for rows.Next() {
			var dataScope domain.DataScopeModel
			err := rows.Scan(
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
				&userRole.Id,
				&userRole.Role,
				&userRole.CreatedAt,
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
			helper.PanicIfError(err)

			dataScopes = append(dataScopes, dataScope)
		}

		return domain.DataContextModel{
			User:       user,
			UserRole:   userRole,
			DataScopes: dataScopes,
		}, nil
	} else {
		return domain.DataContextModel{}, errors.New("data not found")
	}
}

func (repository *Oauth2RepositoryImpl) UpdateUserTokenVersion(ctx context.Context, tx *sql.Tx, userId, tokenVersion int64) {
	querySql := "UPDATE user SET token_version = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, querySql, tokenVersion, userId)
	helper.PanicIfError(err)
}

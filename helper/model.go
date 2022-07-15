package helper

import (
	"simple-oauth-service/model/domain"
	"simple-oauth-service/model/response"
)

func ToUserResponse(user domain.UserModel) response.UserResponse {
	return response.UserResponse{
		Id:            user.Id.Int64,
		Email:         user.Email.String,
		FirstName:     user.FirstName.String,
		LastName:      user.LastName.String,
		UserRoleId:    user.UserRoleId.Int64,
		CompanyId:     user.CompanyId.Int64,
		PrincipalId:   user.PrincipalId.Int64,
		DistributorId: user.DistributorId.Int64,
		BuyerId:       user.BuyerId.Int64,
		TokenVersion:  user.TokenVersion.Int64,
		IsVerified:    user.IsVerified.Bool,
		IsDelete:      user.IsDelete.Bool,
		CreatedAt:     user.CreatedAt.Time,
		UpdatedAt:     user.UpdatedAt.Time,
		CreatedBy:     user.CreatedBy.String,
		UpdatedBy:     user.UpdatedBy.String,
	}
}

func ToUserResponses(users []domain.UserModel) []response.UserResponse {
	var userResponses []response.UserResponse
	for _, user := range users {
		userResponses = append(userResponses, ToUserResponse(user))
	}

	return userResponses
}

func ToClientResponse(client domain.ClientModel) response.ClientResponse {
	return response.ClientResponse{
		Id:              client.Id.Int64,
		UserId:          client.UserId.Int64,
		ApplicationName: client.ApplicationName.String,
		ClientSecret:    client.ClientSecret.String,
		IsDelete:        client.IsDelete.Bool,
		CreatedAt:       client.CreatedAt.Time,
		UpdatedAt:       client.UpdatedAt.Time,
		CreatedBy:       client.CreatedBy.String,
		UpdatedBy:       client.UpdatedBy.String,
	}
}

func ToClientResponses(clients []domain.ClientModel) []response.ClientResponse {
	var clientResponses []response.ClientResponse
	for _, client := range clients {
		clientResponses = append(clientResponses, ToClientResponse(client))
	}

	return clientResponses
}

func ToDataScopeResponse(dataScope domain.DataScopeModel) response.DataScopeResponse {
	return response.DataScopeResponse{
		Id:            dataScope.Id.Int64,
		UserId:        dataScope.UserId.Int64,
		PrincipalId:   dataScope.PrincipalId.Int64,
		DistributorId: dataScope.DistributorId.Int64,
		BuyerId:       dataScope.BuyerId.Int64,
		IsDelete:      dataScope.IsDelete.Bool,
		CreatedAt:     dataScope.CreatedAt.Time,
		UpdatedAt:     dataScope.UpdatedAt.Time,
		CreatedBy:     dataScope.CreatedBy.String,
		UpdatedBy:     dataScope.UpdatedBy.String,
	}
}

func ToDataScopeResponses(dataScopes []domain.DataScopeModel) []response.DataScopeResponse {
	var dataScopeResponses []response.DataScopeResponse
	for _, dataScope := range dataScopes {
		dataScopeResponses = append(dataScopeResponses, ToDataScopeResponse(dataScope))
	}

	return dataScopeResponses
}

func ToUserRoleResponse(userRole domain.UserRoleModel) response.UserRoleResponse {
	return response.UserRoleResponse{
		Id:        userRole.Id.Int64,
		Role:      userRole.Role.String,
		CreatedAt: userRole.CreatedAt.Time,
	}
}

func ToUserRoleResponses(userRoles []domain.UserRoleModel) []response.UserRoleResponse {
	var userRoleResponses []response.UserRoleResponse
	for _, userRole := range userRoles {
		userRoleResponses = append(userRoleResponses, ToUserRoleResponse(userRole))
	}

	return userRoleResponses
}

func IncreaseUserTokenVersion(user *domain.UserModel) {
	if user.TokenVersion.Int64 >= 1000 {
		user.TokenVersion.Int64 = 0
		user.TokenVersion.Valid = true
	} else {
		user.TokenVersion.Int64++
		user.TokenVersion.Valid = true
	}
}

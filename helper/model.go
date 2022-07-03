package helper

import (
	"simple-oauth-service/model/domain"
	"simple-oauth-service/model/response"
)

func ToUserResponse(user domain.UserModel) response.UserResponse {
	return response.UserResponse{
		Id:            user.Id,
		Email:         user.Email,
		FirstName:     user.FirstName,
		LastName:      user.LastName,
		UserRoleId:    user.UserRoleId,
		CompanyId:     user.CompanyId,
		PrincipalId:   user.PrincipalId,
		DistributorId: user.DistributorId,
		BuyerId:       user.BuyerId,
		TokenVersion:  user.TokenVersion,
		IsVerified:    user.IsVerified,
		IsDelete:      user.IsDelete,
		CreatedAt:     user.CreatedAt,
		UpdatedAt:     user.UpdatedAt,
		CreatedBy:     user.CreatedBy,
		UpdatedBy:     user.UpdatedBy,
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
		Id:              client.Id,
		UserId:          client.UserId,
		ApplicationName: client.ApplicationName,
		ClientSecret:    client.ClientSecret,
		IsDelete:        client.IsDelete,
		CreatedAt:       client.CreatedAt,
		UpdatedAt:       client.UpdatedAt,
		CreatedBy:       client.CreatedBy,
		UpdatedBy:       client.UpdatedBy,
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
		Id:            dataScope.Id,
		UserId:        dataScope.UserId,
		PrincipalId:   dataScope.PrincipalId,
		DistributorId: dataScope.DistributorId,
		BuyerId:       dataScope.BuyerId,
		IsDelete:      dataScope.IsDelete,
		CreatedAt:     dataScope.CreatedAt,
		UpdatedAt:     dataScope.UpdatedAt,
		CreatedBy:     dataScope.CreatedBy,
		UpdatedBy:     dataScope.UpdatedBy,
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
		Id:        userRole.Id,
		Role:      userRole.Role,
		CreatedAt: userRole.CreatedAt,
	}
}

func ToUserRoleResponses(userRoles []domain.UserRoleModel) []response.UserRoleResponse {
	var userRoleResponses []response.UserRoleResponse
	for _, userRole := range userRoles {
		userRoleResponses = append(userRoleResponses, ToUserRoleResponse(userRole))
	}

	return userRoleResponses
}

package ctx

import (
	"context"
	"errors"
	"reflect"
	"simple-oauth-service/constanta"
	"simple-oauth-service/model/response"
)

type Context struct {
	context.Context
	User              response.UserResponse        `json:"user"`
	UserRole          response.UserRoleResponse    `json:"user_role"`
	DataScopes        []response.DataScopeResponse `json:"data_scopes"`
	ListUserId        []int64                      `json:"list_user_id"`
	ListPrincipalId   []int64                      `json:"list_principal_id"`
	ListDistributorId []int64                      `json:"list_distributor_id"`
	ListBuyerId       []int64                      `json:"list_buyer_id"`
}

func NewContext(user response.UserResponse, userRole response.UserRoleResponse, dataScopes []response.DataScopeResponse) (Context, error) {
	listUserId, err := GetListIdFromDataScopes(dataScopes, "UserId")
	if err != nil {
		return Context{}, err
	}
	listPrincipalId, err := GetListIdFromDataScopes(dataScopes, "PrincipalId")
	if err != nil {
		return Context{}, err
	}
	listDistributorId, err := GetListIdFromDataScopes(dataScopes, "DistributorId")
	if err != nil {
		return Context{}, err
	}
	listBuyerId, err := GetListIdFromDataScopes(dataScopes, "BuyerId")
	if err != nil {
		return Context{}, err
	}
	return Context{
		User:              user,
		UserRole:          userRole,
		DataScopes:        dataScopes,
		ListUserId:        listUserId,
		ListPrincipalId:   listPrincipalId,
		ListDistributorId: listDistributorId,
		ListBuyerId:       listBuyerId,
	}, nil
}

func ToCtxContext(context context.Context) (Context, error) {
	ctx, ok := context.(Context)
	if !ok {
		return Context{}, errors.New(constanta.CannotConvertToCtxContext)
	}

	return ctx, nil
}

func GetListIdFromDataScopes(dataScopes []response.DataScopeResponse, fieldLookUp string) ([]int64, error) {
	var result []int64

	for _, each := range dataScopes {
		rt := reflect.TypeOf(each)
		rv := reflect.ValueOf(each)
		field, ok := rt.FieldByName(fieldLookUp)
		if !ok {
			return nil, errors.New("no such field exist with name " + fieldLookUp)
		}
		if field.Type.Kind() == reflect.Int64 {
			result = append(result, rv.FieldByName(fieldLookUp).Interface().(int64))
		} else if field.Type.Kind() == reflect.Int {
			result = append(result, int64(rv.FieldByName(fieldLookUp).Interface().(int)))
		} else {
			return nil, errors.New("type field lookup must int or int64")
		}

	}

	return result, nil
}

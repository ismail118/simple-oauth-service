package domain

type DataContextModel struct {
	User       UserModel
	UserRole   UserRoleModel
	DataScopes []DataScopeModel
}

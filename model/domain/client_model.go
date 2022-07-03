package domain

import "time"

type ClientModel struct {
	Id              int64
	UserId          int64
	ApplicationName string
	ClientSecret    string
	IsDelete        bool
	CreatedAt       time.Time
	UpdatedAt       time.Time
	CreatedBy       string
	UpdatedBy       string
}

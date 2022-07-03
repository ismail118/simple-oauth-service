package domain

import "time"

type UserRoleModel struct {
	Id        int64
	Role      string
	CreatedAt time.Time
}

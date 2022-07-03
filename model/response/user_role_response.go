package response

import "time"

type UserRoleResponse struct {
	Id        int64     `json:"id"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
}

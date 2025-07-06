package eventdata

import "time"

type User struct {
	Id                string    `json:"id"`
	Username          string    `json:"username"`
	Name              string    `json:"name"`
	Lastname          string    `json:"lastname"`
	Email             string    `json:"email"`
	Phone             string    `json:"phone"`
	Active            bool      `json:"active"`
	IsAdmin           bool      `json:"isAdmin"`
	RoleId            string    `json:"roleId"`
	PasswordChangedAt time.Time `json:"passwordChangedAt"`
	CreatedAt         time.Time `json:"createdAt"`
}

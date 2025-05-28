package model

type User struct {
	UserID    int    `json:"userId,omitempty" db:"user_id"`
	UserLogin string `json:"userLogin" db:"user_login"`
	Tasks     `json:"tasks,omitempty"`
}

type Users []User

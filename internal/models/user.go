package models

import "time"

type User struct {
	ID        int       `json:"id" db:"id"`
	FullName  string    `json:"full_name" db:"full_name"`
	Username  string    `json:"username" db:"username"`
	Password  string    `json:"password" db:"password"`
	Role      Role      `json:"role" db:"role"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

type Role string

const (
	RoleUser  = "USER"
	RoleAdmin = "ADMIN"
)

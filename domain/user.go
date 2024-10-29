package domain

import (
	"time"

	"gopkg.in/guregu/null.v4"
)

type UserData struct {
	ID        int64     `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	Email     string    `json:"email" db:"email"`
	Password  string    `json:"-" db:"password"`
	Role      string    `json:"role" db:"role"`
	CreatedAt time.Time `json:"-" db:"created_at"`
	UpdatedAt null.Time `json:"-" db:"updated_at"`
	DeletedAt null.Time `json:"-" db:"deleted_at"`
}

type RegisterUserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type RegisterUserResponse struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"empty=false"`
	Password string `json:"password" validate:"empty=false"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

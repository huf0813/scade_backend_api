package domain

import (
	"context"
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Name     string `gorm:"not_null" json:"name"`
	Address  string `gorm:"not_null" json:"address"`
	Email    string `gorm:"not_null;unique" json:"email"`
	Phone    string `gorm:"not_null;unique" json:"phone"`
	Password string `gorm:"not_null" json:"password"`
}

type UserSignInRequest struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UserSignUpRequest struct {
	Name     string `json:"name" validate:"required"`
	Address  string `json:"address" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Phone    string `json:"phone" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UserRepository interface {
	GetUserByEmail(ctx context.Context, email string) (User, error)
	SignUp(ctx context.Context, name, address, email, phone, password string) error
}

type UserUseCase interface {
	SignIn(ctx context.Context, email, password string, expiration int, timeType time.Duration) (string, error)
	SignUp(ctx context.Context, name, address, email, phone, password string) error
	Profile(ctx context.Context, email string) (User, error)
}

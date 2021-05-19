package myql

import (
	"context"
	"errors"
	"github.com/huf0813/scade_backend_api/domain"
	"gorm.io/gorm"
)

type UserRepoMysql struct {
	DB *gorm.DB
}

func NewUserRepoMysql(db *gorm.DB) domain.UserRepository {
	return &UserRepoMysql{DB: db}
}

func (u *UserRepoMysql) GetUserByEmail(ctx context.Context, email string) (domain.User, error) {
	var user domain.User
	result := u.DB.
		WithContext(ctx).
		Where("email = ?", email).
		Where("email = ?", email).
		First(&user)
	if err := result.Error; err != nil {
		return domain.User{}, err
	}
	return user, nil
}

func (u *UserRepoMysql) SignUp(ctx context.Context, name, address, email, phone, password string) error {
	user := domain.User{
		Name:     name,
		Address:  address,
		Email:    email,
		Phone:    phone,
		Password: password,
	}
	result := u.DB.WithContext(ctx).Create(&user)
	if err := result.Error; err != nil {
		return err
	}
	if rows := result.RowsAffected; rows <= 0 {
		return errors.New("failed to insert data, empty feedback")
	}
	return nil
}

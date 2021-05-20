package usecase

import (
	"context"
	"fmt"
	"github.com/huf0813/scade_backend_api/domain"
	"github.com/huf0813/scade_backend_api/utils/auth"
	"github.com/huf0813/scade_backend_api/utils/security"
	"time"
)

type UserUseCase struct {
	userRepoMysql domain.UserRepository
	timeOut       time.Duration
}

func NewUserUseCase(u domain.UserRepository, timeOut time.Duration) domain.UserUseCase {
	return &UserUseCase{
		userRepoMysql: u,
		timeOut:       timeOut,
	}
}

func (u *UserUseCase) SignIn(ctx context.Context,
	email, password string,
	expiration int,
	timeType time.Duration) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, u.timeOut)
	defer cancel()

	result, err := u.userRepoMysql.GetUserByEmail(ctx, email)
	if err != nil {
		return "", err
	}

	if err := security.NewValidatingValue(result.Password, password); err != nil {
		return "", err
	}

	duration := time.Duration(expiration) * timeType
	token, err := auth.NewJWT(duration, email, "user")
	if err != nil {
		return "", err
	}

	token = fmt.Sprintf("%s %s", "Bearer", token)
	return token, nil
}

func (u *UserUseCase) SignUp(ctx context.Context, name, address, email, phone, password string) error {
	ctx, cancel := context.WithTimeout(ctx, u.timeOut)
	defer cancel()

	password, err := security.NewHashingValue(password)
	if err != nil {
		return err
	}

	if err := u.userRepoMysql.SignUp(ctx, name, address, email, phone, password); err != nil {
		return err
	}

	return nil
}

func (u *UserUseCase) Profile(ctx context.Context, email string) (domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, u.timeOut)
	defer cancel()

	result, err := u.userRepoMysql.GetUserByEmail(ctx, email)
	if err != nil {
		return domain.User{}, err
	}

	return result, nil
}

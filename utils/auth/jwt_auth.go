package auth

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4/middleware"
	"os"
	"time"
)

type envTokenSecret struct {
	secret string
}

func readTokenEnv() (token envTokenSecret, err error) {
	if err = godotenv.Load(); err != nil {
		return envTokenSecret{}, err
	}
	token.secret = os.Getenv("TOKEN_SECRET")
	return token, nil
}

func NewAuthMiddleware() (middleware.JWTConfig, error) {
	env, err := readTokenEnv()
	if err != nil {
		return middleware.JWTConfig{}, err
	}
	secret := []byte(env.secret)
	isAuthorized := middleware.JWTConfig{
		SigningKey: secret,
	}
	return isAuthorized, nil
}

func NewJWT(expirationTime time.Duration, email, role string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = email
	claims["role"] = role
	exp := time.Now().Add(expirationTime).Unix()
	claims["exp"] = exp

	env, err := readTokenEnv()
	if err != nil {
		return "", err
	}
	secret := []byte(env.secret)
	t, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}
	return t, nil
}

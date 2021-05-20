package auth

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/huf0813/scade_backend_api/domain"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4/middleware"
	"os"
	"strings"
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
	exp := time.Now().Add(expirationTime).Unix()
	claims := &domain.CustomToken{
		Email: email,
		Role:  role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: exp,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

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

func TokenExtraction(hashedToken string) (*domain.CustomToken, error) {
	env, err := readTokenEnv()
	if err != nil {
		return nil, err
	}

	hashedToken = strings.Replace(hashedToken, "Bearer ", "", 1)
	token, err := jwt.Parse(hashedToken, func(token *jwt.Token) (interface{}, error) {
		result := []byte(env.secret)
		return result, nil
	})
	if err != nil {
		return nil, err
	}

	tokenClaims := token.Claims
	claims := tokenClaims.(jwt.MapClaims)
	customToken := &domain.CustomToken{
		Email: claims["email"].(string),
		Role:  claims["role"].(string),
	}
	return customToken, nil
}

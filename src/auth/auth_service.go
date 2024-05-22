package auth

import (
	"go-echo-boilerplate/common"
	config "go-echo-boilerplate/setup"
	UserModels "go-echo-boilerplate/src/models/users/model"
	"os"
	"sync"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type authService struct{}

var singleton AuthService
var once sync.Once

func GetAuthService() AuthService {
	once.Do(func() {
		singleton = &authService{}
	})
	return singleton
}

//func SetAuthService(service AuthService) AuthService {
//	original := singleton
//	singleton = service
//	return original
//}

type AuthService interface {
	GetAccessToken(user *UserModels.User) (string, error)
}

func (s *authService) GetAccessToken(user *UserModels.User) (string, error) {
	claims := &common.JwtCustomClaims{
		Id:   user.ID,
		Role: user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * config.TokenExpiresIn)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
}

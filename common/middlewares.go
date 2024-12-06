package common

import (
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

type JwtCustomClaims struct {
	Name string   `json:"name"`
	Id   uint     `json:"id"`
	Role UserRole `json:"role"`
	jwt.RegisteredClaims
}

func JwtMiddleWare() echo.MiddlewareFunc {
	key := viper.GetString("Jwt.SecretKey")
	return echojwt.WithConfig(echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(JwtCustomClaims)
		},
		SigningKey: []byte(key),
	})
}

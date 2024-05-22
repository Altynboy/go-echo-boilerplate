package common

import (
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"

	// "https://github.com/labstack/echo-jwt"
	// uuid "github.com/satori/go.uuid"
	"os"
)

type JwtCustomClaims struct {
	Name string    `json:"name"`
	Id   uint 	   `json:"id"`
	Role UserRole  `json:"role"`
	jwt.RegisteredClaims
}

// func(c echo.Context) NewClaim() jwt.Claims {
//     return new(JwtCustomClaims)
// }

func JwtMiddleWare() echo.MiddlewareFunc {
	key := os.Getenv("JWT_SECRET_KEY")

	return echojwt.WithConfig(echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims{
			return new(JwtCustomClaims)
		},
		SigningKey: []byte(key),
	})
}
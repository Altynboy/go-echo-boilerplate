package main

import (
	"go-echo-boilerplate/common"
	cp "go-echo-boilerplate/database/config"
	"go-echo-boilerplate/routes"
	"log"
	"os"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func setup() {
	env := os.Getenv("ENVIRONMENT")
	log.Println(env)
	
	// setup for dev
	if env == "" {
		err := cp.Config().Init("./config/config.json")
		if err != nil {
			log.Fatal("Error loading config file")	
		}
		os.Setenv("DB_HOST", cp.ConfigData.DB.HOST)
		os.Setenv("DB_PORT", cp.ConfigData.DB.PORT)
		os.Setenv("DB_DATABASE", cp.ConfigData.DB.DATABASE)
		os.Setenv("DB_USERNAME", cp.ConfigData.DB.USERNAME)
		os.Setenv("DB_PASSWORD", cp.ConfigData.DB.PASSWORD)
		os.Setenv("JWT_SECRET_KEY", cp.ConfigData.JWT.SECRET_KEY)
	}
	
	//Define API wrapper
	api := echo.New()
	api.Validator = &common.CustomValidator{Validator: validator.New()}
	api.Use(middleware.Logger())
	api.Use(middleware.Recover())
	// CORS middleware for API endpoint.
	api.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))
	routes.DefineApiRoute(api)

	server := echo.New()
	server.Any("/*", func(c echo.Context) (err error) {
		req := c.Request()
		res := c.Response()
		if req.URL.Path[:4] == "/api" {
			api.ServeHTTP(res, req)
		}

		return
	})
	server.Logger.Fatal(server.Start(":1200"))
}

func main() {
	setup()
}

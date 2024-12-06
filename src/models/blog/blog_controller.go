package blog

import (
	"go-echo-boilerplate/common"
	"go-echo-boilerplate/database"
	BlogModel "go-echo-boilerplate/src/models/blog/model"
	"go-echo-boilerplate/src/wrapper"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type (
	BlogController struct{}

	RequestCreateBlog struct {
		Title   string `json:"Title" validate:"required"`
		Content string `json:"Content" validate:"required"`
	}
)

func (controller BlogController) Routes() []common.Route {
	return []common.Route{
		{
			Method:     echo.GET,
			Path:       "/blogs",
			Handler:    controller.GetBlogs,
			Middleware: []echo.MiddlewareFunc{common.JwtMiddleWare()},
		},
		{
			Method:     echo.POST,
			Path:       "/blog",
			Handler:    controller.CreateBlog,
			Middleware: []echo.MiddlewareFunc{common.JwtMiddleWare()},
		},
	}
}

func (controller BlogController) GetBlogs(ctx echo.Context) error {
	db := database.Instance()
	var blogs []BlogModel.Blog

	token, ok := ctx.Get("user").(*jwt.Token)
	if !ok {
		return wrapper.Response(ctx, http.StatusUnauthorized, "Invalid token")
	}
	claims, ok := token.Claims.(*common.JwtCustomClaims)
	if !ok {
		return wrapper.Response(ctx, http.StatusUnauthorized, "Invalid token")
	}

	res := db.Find(&blogs, "user_id = ?", claims.Id)
	if res.Error != nil {
		return wrapper.Response(ctx, http.StatusInternalServerError, res.Error.Error())
	}
	// if res.RowsAffected == 0 {
	// 	return wrapper.Response(ctx, http.StatusNotFound, "No blogs found")
	// }

	return wrapper.Response(ctx, http.StatusOK, "Success", map[string]interface{}{"blogs": blogs})
}

func (controller BlogController) CreateBlog(ctx echo.Context) error {
	params := new(RequestCreateBlog)
	if err := ctx.Bind(params); err != nil {
		return wrapper.Response(ctx, http.StatusBadRequest, err.Error())
	}
	if err := ctx.Validate(params); err != nil {
		return wrapper.Response(ctx, http.StatusBadRequest, err.Error())
	}

	token, ok := ctx.Get("user").(*jwt.Token)
	if !ok {
		return wrapper.Response(ctx, http.StatusUnauthorized, "Invalid token")
	}
	claims, ok := token.Claims.(*common.JwtCustomClaims)
	if !ok {
		return wrapper.Response(ctx, http.StatusUnauthorized, "Invalid token")
	}

	db := database.Instance()
	res := db.Create(&BlogModel.Blog{
		UserId:  claims.Id,
		Title:   params.Title,
		Content: params.Content,
	})
	if res.Error != nil {
		return wrapper.Response(ctx, http.StatusInternalServerError, res.Error.Error())
	}
	if res.RowsAffected == 0 {
		return wrapper.Response(ctx, http.StatusInternalServerError, "Failed to create blog")
	}

	return wrapper.Response(ctx, http.StatusOK, "Success")
}

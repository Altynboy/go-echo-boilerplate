package auth

import (
	"net/http"

	"go-echo-boilerplate/common"
	"go-echo-boilerplate/common/utils"
	"go-echo-boilerplate/src/models/users"
	"go-echo-boilerplate/src/wrapper"

	"github.com/labstack/echo/v4"
)

type (
	AuthController struct {
	}

	RequestRegisterUser struct {
		Phone    string `json:"Phone"  validate:"required"`
		Password string `json:"Password" validate:"required"`
		// FirstName string `json:"firstName" validate:"omitempty"`
		// LastName string `json:"lastName" validate:"omitempty"`
		// MiddleName string `json:"middleName" validate:"omitempty"`
		// Email string `json:"email" validate:"omitempty"`
	}

	RequestLogin struct {
		Phone    string `json:"Phone"  validate:"required"`
		Password string `json:"Password" validate:"required"`
	}
	
	RequestBackupPwd struct {
		Phone    	string `json:"Phone"  validate:"required"`
		Password 	string `json:"Password" validate:"required"`
		NewPassword string `json:"NewPassword" validate:"required"`
	}

	RequestCheckPhone struct {
		Phone string `json:"Phone" validate:"required"`
	}
)


func (controller AuthController) Routes() []common.Route {
	return []common.Route{
		{
			Method:  echo.POST,
			Path:    "/auth/register",
			Handler: controller.Register,
		},
		{
			Method:  echo.POST,
			Path:    "/auth/login",
			Handler: controller.Login,
		},
		{
			Method:  echo.POST,
			Path:    "/auth/backup/pwd",
			Handler: controller.BackupPwd,
		},
		{
			Method:  echo.POST,
			Path:    "/auth/check/phone",
			Handler: controller.CheckPhone,
		},
	}
}

func (controller AuthController) Register(ctx echo.Context) error {
	params := new(RequestRegisterUser)
	if err := ctx.Bind(params); err != nil {
		return wrapper.Response(ctx, http.StatusBadRequest, err.Error())
	}
	if err := ctx.Validate(params); err != nil {
		return wrapper.Response(ctx, http.StatusBadRequest, err.Error())
	}
	
	if users.GetUsersService().IsPhoneExist(params.Phone) {
		return wrapper.Response(ctx, http.StatusBadRequest, "Phone already exist")
	}

	_, err := users.GetUsersService().AddUser( params.Phone, params.Password)
	if err != nil {
		return wrapper.Response(ctx, http.StatusBadRequest, err.Error())
	}
	
	// user, err := users.GetUsersService().FindUserByPhone(params.Phone)
	// if err != nil {
	// 	return wrapper.Response(ctx, http.StatusUnauthorized, err.Error())
	// }
	// if matched := utils.GetPasswordUtil().CheckPasswordHash(params.Password, user.Password); !matched {
	// 	return wrapper.Response(ctx, http.StatusUnauthorized, "Invalid phone or password")
	// }

	// Create token
	// token, _ := GetAuthService().GetAccessToken(user)

	// return wrapper.Response(ctx, http.StatusOK, "Success", map[string]string{
	// 	token: token,
	// })
	return wrapper.Response(ctx, http.StatusOK, "Success")
}

func (controller AuthController) Login(ctx echo.Context) error {
	params := new(RequestLogin)
	if err := ctx.Bind(params); err != nil {
		return wrapper.Response(ctx, http.StatusBadRequest, err.Error())
	}
	if err := ctx.Validate(params); err != nil {
		return wrapper.Response(ctx, http.StatusBadRequest, err.Error())
	}

	user, err := users.GetUsersService().FindUserByPhone(params.Phone)
	if err != nil {
		return wrapper.Response(ctx, http.StatusUnauthorized, err.Error())
	}
	if matched := utils.GetPasswordUtil().CheckPasswordHash(params.Password, user.Password); !matched {
		return wrapper.Response(ctx, http.StatusUnauthorized, "Invalid phone or password")
	}
	
	// Create token
	token, _ := GetAuthService().GetAccessToken(user)

	return wrapper.Response(ctx, http.StatusOK, "Success", map[string]string{
		"token": token,
	})
}


func (controller AuthController) BackupPwd(ctx echo.Context) error {
	params := new(RequestBackupPwd)
	if err := ctx.Bind(params); err != nil {
		return wrapper.Response(ctx, http.StatusBadRequest, err.Error())
	}
	if err := ctx.Validate(params); err != nil {
		return wrapper.Response(ctx, http.StatusBadRequest, err.Error())
	}
	
	 
	user, err := users.GetUsersService().FindUserByPhone(params.Phone)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "Invalid phone or password")
	}
	if matched := utils.GetPasswordUtil().CheckPasswordHash(params.Password, user.Password); !matched {
		return echo.NewHTTPError(http.StatusUnauthorized, "Invalid phone or password")
	}

	_, err = users.GetUsersService().ChangePwd(params.Phone, params.NewPassword )
	if err != nil {
		return wrapper.Response(ctx, http.StatusBadRequest, err.Error())
	}

	// Create token
	// token, _ := GetAuthService().GetAccessToken(user)
	// return wrapper.Response(ctx, http.StatusOK, "Success", map[string]string{
	// 	"token": token,
	// })
	return wrapper.Response(ctx, http.StatusOK, "Success")
}


func (controller AuthController) CheckPhone(ctx echo.Context) error {
	params := new(RequestCheckPhone)
	if err := ctx.Bind(params); err != nil {
		return wrapper.Response(ctx, http.StatusBadRequest, err.Error())
	}
	if err := ctx.Validate(params); err != nil {
		return wrapper.Response(ctx, http.StatusBadRequest, err.Error())
	}

	if !users.GetUsersService().IsPhoneExist(params.Phone) {
		return wrapper.Response(ctx, http.StatusBadRequest, "phone not found")
	} 
	
	return wrapper.Response(ctx, http.StatusOK, "Success")
}





package auth

import (
	"net/http"
	"todo-list-app/delivery/controllers/common"
	"todo-list-app/middlewares"
	"todo-list-app/repository/auth"

	"github.com/labstack/echo/v4"
)

type AuthController struct {
	repo auth.Auth
}

func New(repo auth.Auth) *AuthController {
	return &AuthController{
		repo: repo,
	}
}

func (ac *AuthController) Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		Userlogin := LoginReqFormat{}

		c.Bind(&Userlogin)
		err_validate := c.Validate(&Userlogin)

		if err_validate != nil {
			return c.JSON(http.StatusBadRequest, common.BadRequest(http.StatusBadRequest, "There is some problem from input", nil))
		}

		checkedUser, err_repo := ac.repo.Login(Userlogin.Email, Userlogin.Password)

		if err_repo != nil {
			var statusCode int = 500
			if err_repo.Error() == "email not found" {
				statusCode = http.StatusUnauthorized
			} else if err_repo.Error() == "incorrect password" {
				statusCode = http.StatusUnauthorized
			}
			return c.JSON(statusCode, common.InternalServerError(statusCode, err_repo.Error(), nil))
		}
		token, err := middlewares.GenerateToken(checkedUser)
		response := UserLoginResponse{
			User_uid: checkedUser.UserUid,
			Name:     checkedUser.Name,
			Email:    checkedUser.Email,
			Token:    token,
		}

		if err != nil {
			return c.JSON(http.StatusNotAcceptable, common.BadRequest(http.StatusNotAcceptable, "error in process token", nil))
		}

		return c.JSON(http.StatusOK, common.Success(http.StatusOK, "USERS - berhasil masuk, mendapatkan token baru", response))

	}
}

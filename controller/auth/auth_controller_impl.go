package auth_controller

import (
	"kredit_plus/models"
	auth_service "kredit_plus/service/auth"
	"net/http"

	"github.com/labstack/echo/v4"
)

type authControllerImpl struct {
	AuthService auth_service.AuthService
}

func NewAuthController(authService auth_service.AuthService) authControllerImpl {
	return authControllerImpl{
		AuthService: authService,
	}
}

func (a *authControllerImpl) Login(ctx echo.Context) error {
	model := new(models.Auths)
	if err := ctx.Bind(model); err != nil {
		return ctx.JSON(http.StatusBadRequest, models.Response{
			Status:  false,
			Message: "invalid body request",
		})
	}

	loginInfo, err := a.AuthService.Login(model)
	if err != nil {
		return ctx.JSON(http.StatusUnauthorized, models.Response{
			Status:  false,
			Message: err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, models.Response{
		Status:  true,
		Message: "success login",
		Data:    loginInfo,
	})
}

func (a *authControllerImpl) RefreshToken(ctx echo.Context) error {
	model := new(models.RefreshTokenRequest)
	if err := ctx.Bind(model); err != nil {
		return ctx.JSON(http.StatusBadRequest, models.Response{
			Status:  false,
			Message: "invalid body request",
		})
	}

	loginInfo, err := a.AuthService.RefreshToken(model)
	if err != nil {
		return ctx.JSON(http.StatusUnauthorized, models.Response{
			Status:  false,
			Message: err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, models.Response{
		Status:  true,
		Message: "success refresh token",
		Data:    loginInfo,
	})
}

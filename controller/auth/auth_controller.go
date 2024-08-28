package auth_controller

import "github.com/labstack/echo/v4"

type AuthController interface {
	Login(ctx echo.Context) error
	RefreshToken(ctx echo.Context) error
}

package user_controller

import (
	"github.com/labstack/echo/v4"
)

type UserController interface {
	Register(ctx echo.Context) error
	GetUserProfile(ctx echo.Context) error
}

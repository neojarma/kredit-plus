package usertenor_controller

import "github.com/labstack/echo/v4"

type UserTenorController interface {
	GetUserTenor(ctx echo.Context) error
}

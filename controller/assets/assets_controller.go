package assets_controller

import "github.com/labstack/echo/v4"

type AssetsController interface {
	GetAllAssets(ctx echo.Context) error
}

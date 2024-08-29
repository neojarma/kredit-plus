package assets_controller

import (
	"kredit_plus/models"
	assets_service "kredit_plus/service/assets"
	"net/http"

	"github.com/labstack/echo/v4"
)

type AssetsControllerImpl struct {
	AssetsService assets_service.AssetsService
}

func NewAssetsController(service assets_service.AssetsService) AssetsController {
	return &AssetsControllerImpl{
		AssetsService: service,
	}
}

func (c *AssetsControllerImpl) GetAllAssets(ctx echo.Context) error {
	assets, err := c.AssetsService.GetAllAsset()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, models.Response{
			Status:  false,
			Message: err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, models.Response{
		Status:  true,
		Message: "success",
		Data:    assets,
	})
}

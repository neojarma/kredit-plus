package peminjaman_controller

import (
	"kredit_plus/models"
	peminjaman_service "kredit_plus/service/peminjaman"
	"net/http"

	"github.com/labstack/echo/v4"
)

type peminjamanControllerImpl struct {
	PeminjamanService peminjaman_service.PeminjamanService
}

func NewPeminjaman(peminjamanService peminjaman_service.PeminjamanService) PeminjamanController {
	return &peminjamanControllerImpl{
		PeminjamanService: peminjamanService,
	}
}

func (c *peminjamanControllerImpl) Kredit(ctx echo.Context) error {
	model := new(models.PeminjamanRequest)
	if err := ctx.Bind(model); err != nil {
		return ctx.JSON(http.StatusBadRequest, models.Response{
			Status:  false,
			Message: "invalid body request",
		})
	}

	if model.Tenor != 1 && model.Tenor != 2 && model.Tenor != 3 && model.Tenor != 6 {
		return ctx.JSON(http.StatusBadRequest, models.Response{
			Status:  false,
			Message: "invalid tenor",
		})
	}

	userId := ctx.Get("user_id").(string)
	model.ID_User = userId
	if err := c.PeminjamanService.Kredit(model); err != nil {
		return ctx.JSON(http.StatusInternalServerError, models.Response{
			Status:  false,
			Message: err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, models.Response{
		Status:  true,
		Message: "success",
	})
}

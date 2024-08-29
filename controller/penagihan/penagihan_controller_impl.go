package penagihan_controller

import (
	"kredit_plus/models"
	penagihan_service "kredit_plus/service/penagihan"
	"net/http"

	"github.com/labstack/echo/v4"
)

type penagihanControllerImpl struct {
	PenagihanService penagihan_service.PenagihanService
}

func NewPenagihanController(penagihanService penagihan_service.PenagihanService) PenagihanController {
	return &penagihanControllerImpl{
		PenagihanService: penagihanService,
	}
}

func (c *penagihanControllerImpl) BayarTagihan(ctx echo.Context) error {

	payload := new(models.PenagihanRequest)
	if err := ctx.Bind(payload); err != nil {
		return ctx.JSON(http.StatusBadRequest, models.Response{
			Status:  false,
			Message: "invalid body request",
		})
	}

	payload.IDUser = ctx.Get("user_id").(string)
	if err := c.PenagihanService.BayarTagihan(payload); err != nil {
		return ctx.JSON(http.StatusInternalServerError, models.Response{
			Status:  false,
			Message: err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, models.Response{
		Status:  true,
		Message: "success melakukan pembayaran",
	})
}

func (c *penagihanControllerImpl) GetAllPenagihan(ctx echo.Context) error {
	peminjamanID := ctx.Param("id")

	penagihanInfo, err := c.PenagihanService.GetDataPenagihan(peminjamanID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, models.Response{
			Status:  false,
			Message: err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, models.Response{
		Status:  true,
		Message: "success melakukan pembayaran",
		Data:    penagihanInfo,
	})
}

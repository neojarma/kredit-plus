package usertenor_controller

import (
	"kredit_plus/models"
	usertenor_service "kredit_plus/service/user_tenor"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserTenorControllerImpl struct {
	UserTenorService usertenor_service.UserTenorService
}

func NewUserTenorController(service usertenor_service.UserTenorService) UserTenorController {
	return &UserTenorControllerImpl{
		UserTenorService: service,
	}
}

func (c *UserTenorControllerImpl) GetUserTenor(ctx echo.Context) error {

	idUser := ctx.Get("user_id").(string)
	res, err := c.UserTenorService.GetUserTenor(idUser)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, models.Response{
			Status:  false,
			Message: err.Error(),
		})
	}

	return ctx.JSON(http.StatusOK, models.Response{
		Status:  true,
		Message: "success",
		Data:    res,
	})
}

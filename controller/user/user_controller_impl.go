package user_controller

import (
	"kredit_plus/models"
	user_service "kredit_plus/service/user"
	"net/http"

	"github.com/labstack/echo/v4"
)

type userControllerImpl struct {
	UserService user_service.UserService
}

func NewUserController(userService user_service.UserService) UserController {
	return &userControllerImpl{
		UserService: userService,
	}
}

func (c *userControllerImpl) GetUserProfile(ctx echo.Context) error {
	userId := ctx.Get("user_id").(string)
	userInfo, err := c.UserService.GetUser(userId)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, models.Response{
			Status:  false,
			Message: "user not found",
		})
	}

	return ctx.JSON(http.StatusOK, models.Response{
		Status:  true,
		Message: "success",
		Data:    userInfo,
	})
}

func (c *userControllerImpl) Register(ctx echo.Context) error {
	payload := new(models.UserRequestRegister)
	if err := ctx.Bind(payload); err != nil {
		return ctx.JSON(http.StatusBadRequest, models.Response{
			Status:  false,
			Message: "invalid body request",
		})
	}

	ktp, err := ctx.FormFile("foto_ktp")
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, models.Response{
			Status:  false,
			Message: "invalid body request",
		})
	}

	selfie, err := ctx.FormFile("foto_selfie")
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, models.Response{
			Status:  false,
			Message: "invalid body request",
		})
	}

	payload.FotoKTP = ktp
	payload.FotoSelfie = selfie

	if err := c.UserService.Register(payload); err != nil {
		return ctx.JSON(http.StatusInternalServerError, models.Response{
			Status:  false,
			Message: err.Error(),
		})
	}

	return ctx.JSON(http.StatusCreated, models.Response{
		Status:  true,
		Message: "Success Regist New User",
	})

}

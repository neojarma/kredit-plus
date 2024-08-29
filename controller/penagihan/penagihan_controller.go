package penagihan_controller

import "github.com/labstack/echo/v4"

type PenagihanController interface {
	BayarTagihan(ctx echo.Context) error
	GetAllPenagihan(ctx echo.Context) error
}

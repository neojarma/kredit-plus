package peminjaman_controller

import (
	"github.com/labstack/echo/v4"
)

type PeminjamanController interface {
	Kredit(ctx echo.Context) error
	GetAllPeminjaman(ctx echo.Context) error
}

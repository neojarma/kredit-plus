package router

import (
	"kredit_plus/auth"
	auth_controller "kredit_plus/controller/auth"
	peminjaman_controller "kredit_plus/controller/peminjaman"
	penagihan_controller "kredit_plus/controller/penagihan"
	user_controller "kredit_plus/controller/user"
	assets_repository "kredit_plus/repository/assets"
	auth_repository "kredit_plus/repository/auth"
	limit_repository "kredit_plus/repository/limit"
	peminjaman_repository "kredit_plus/repository/peminjaman"
	penagihan_repository "kredit_plus/repository/penagihan"
	user_repository "kredit_plus/repository/user"
	user_tenor_repository "kredit_plus/repository/user_tenor"
	auth_service "kredit_plus/service/auth"
	peminjaman_service "kredit_plus/service/peminjaman"
	penagihan_service "kredit_plus/service/penagihan"
	user_service "kredit_plus/service/user"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Setups struct {
	Echo *echo.Echo
	DB   *gorm.DB
}

func Router(setup Setups) {

	userRepo := user_repository.NewUserRepository(setup.DB)
	authRepo := auth_repository.NewAuthRepository(setup.DB)
	userLimitRepo := limit_repository.NewTenorLimitRepository(setup.DB)
	userTenorRepo := user_tenor_repository.NewUserTenorRepository(setup.DB)
	userService := user_service.NewUserService(authRepo, userRepo, userTenorRepo, userLimitRepo)
	userController := user_controller.NewUserController(userService)

	authService := auth_service.NewAuthService(authRepo)
	authController := auth_controller.NewAuthController(authService)

	assetRepo := assets_repository.NewAssetRepository(setup.DB)
	peminjamanRepository := peminjaman_repository.NewPeminjaman(setup.DB)
	penagihanRepo := penagihan_repository.NewPenagihanRepository(setup.DB)
	peminjamanService := peminjaman_service.NewPeminjamanService(peminjamanRepository, assetRepo, userTenorRepo, penagihanRepo)
	peminjamanController := peminjaman_controller.NewPeminjaman(peminjamanService)
	penagihanService := penagihan_service.NewPenagihanService(penagihanRepo, peminjamanRepository, userTenorRepo)
	penagihanController := penagihan_controller.NewPenagihanController(penagihanService)

	setup.Echo.POST("/register", userController.Register)
	setup.Echo.POST("/login", authController.Login)
	setup.Echo.POST("/refresh-token", authController.RefreshToken)

	setup.Echo.POST("/kredit", peminjamanController.Kredit, auth.VerifyToken)
	setup.Echo.POST("/bayar-kredit", penagihanController.BayarTagihan, auth.VerifyToken)

	setup.Echo.GET("/profile", userController.GetUserProfile, auth.VerifyToken)
	setup.Echo.Static("/assets", "assets")
	setup.Echo.GET("/peminjaman", peminjamanController.GetAllPeminjaman, auth.VerifyToken)
	setup.Echo.GET("/penagihan/:id", penagihanController.GetAllPenagihan, auth.VerifyToken)

	// get all asset
	// get limit user
}

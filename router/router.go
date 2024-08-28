package router

import (
	auth_controller "kredit_plus/controller/auth"
	user_controller "kredit_plus/controller/user"
	auth_repository "kredit_plus/repository/auth"
	limit_repository "kredit_plus/repository/limit"
	user_repository "kredit_plus/repository/user"
	"kredit_plus/repository/user_tenor"
	auth_service "kredit_plus/service/auth"
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
	userTenorRepo := user_tenor.NewUserTenorRepository(setup.DB)
	userService := user_service.NewUserService(authRepo, userRepo, userTenorRepo, userLimitRepo)
	userController := user_controller.NewUserController(userService)

	authService := auth_service.NewAuthService(authRepo)
	authController := auth_controller.NewAuthController(authService)

	setup.Echo.POST("/register", userController.Register)
	setup.Echo.POST("/login", authController.Login)
	setup.Echo.POST("/refresh-token", authController.RefreshToken)

}

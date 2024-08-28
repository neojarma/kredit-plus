package auth_service

import "kredit_plus/models"

type AuthService interface {
	Login(payload *models.Auths) (*models.LoginResponse, error)
	RefreshToken(payload *models.RefreshTokenRequest) (*models.LoginResponse, error)
}

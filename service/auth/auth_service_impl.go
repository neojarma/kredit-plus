package auth_service

import (
	"errors"
	"kredit_plus/auth"
	"kredit_plus/helper"
	"kredit_plus/models"
	auth_repository "kredit_plus/repository/auth"
)

type authServiceImpl struct {
	AuthRepo auth_repository.AuthRepository
}

func NewAuthService(authRepo auth_repository.AuthRepository) AuthService {
	return &authServiceImpl{
		AuthRepo: authRepo,
	}
}

func (a *authServiceImpl) Login(payload *models.Auths) (*models.LoginResponse, error) {

	// check to db
	loginInfo, err := a.AuthRepo.Login(payload)
	if err != nil {
		return nil, err
	}

	// check pass hash
	if ok := helper.CheckPasswordHash(payload.Password, loginInfo.Password); !ok {
		return nil, errors.New("wrong username or password")
	}

	// generate access Token
	accessToken, err := auth.GenerateAccessToken(loginInfo.IDUser)
	if err != nil {
		return nil, err
	}

	// generate Refresh Token
	refreshToken, err := auth.GenerateRefreshToken(loginInfo.IDUser)
	if err != nil {
		return nil, err
	}

	return &models.LoginResponse{
		ID:            loginInfo.IDUser,
		AccessToken:   accessToken,
		RefreshTokens: refreshToken,
	}, nil
}

func (a *authServiceImpl) RefreshToken(payload *models.RefreshTokenRequest) (*models.LoginResponse, error) {
	return auth.CheckRefreshToken(payload.RefreshToken)
}

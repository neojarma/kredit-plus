package usertenor_service

import "kredit_plus/models"

type UserTenorService interface {
	GetUserTenor(userID string) (*models.UserTenor, error)
}

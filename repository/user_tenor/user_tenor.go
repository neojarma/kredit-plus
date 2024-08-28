package user_tenor_repository

import "kredit_plus/models"

type UserTenor interface {
	InsertUserTenorProfile(payload *models.UserTenor) error
	GetUserTenorProfile(payload *models.UserTenor) (*models.UserTenor, error)
	UpdateUserTenorProfile(payload *models.UserTenor) error
}

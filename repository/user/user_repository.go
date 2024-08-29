package user_repository

import "kredit_plus/models"

type UserRepository interface {
	InsertUserProfile(payload *models.User) error
	GetUserProfile(payload *models.User) (*models.User, error)
}

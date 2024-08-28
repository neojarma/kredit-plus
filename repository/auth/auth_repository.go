package auth_repository

import "kredit_plus/models"

type AuthRepository interface {
	InsertToAuth(payload *models.Auths) error
	Login(payload *models.Auths) (*models.Auths, error)
}

package user_service

import "kredit_plus/models"

type UserService interface {
	Register(payload *models.UserRequestRegister) error
	GetUser(idUser string) (*models.User, error)
}

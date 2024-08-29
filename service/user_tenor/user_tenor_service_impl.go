package usertenor_service

import (
	"kredit_plus/models"
	user_tenor_repository "kredit_plus/repository/user_tenor"
)

type UserTenorServiceImpl struct {
	UserTenorRepo user_tenor_repository.UserTenor
}

func NewUserTenorService(repo user_tenor_repository.UserTenor) UserTenorService {
	return &UserTenorServiceImpl{
		UserTenorRepo: repo,
	}
}

func (s *UserTenorServiceImpl) GetUserTenor(userID string) (*models.UserTenor, error) {
	return s.UserTenorRepo.GetUserTenorProfile(&models.UserTenor{
		ID_User: userID,
	})
}

package user_tenor

import "kredit_plus/models"

type UserTenor interface {
	InsertUserTenorProfile(payload *models.UserTenor) error
}

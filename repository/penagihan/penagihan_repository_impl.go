package penagihan_repository

import (
	"kredit_plus/models"

	"gorm.io/gorm"
)

type penagihanRepositoryImpl struct {
	DB *gorm.DB
}

func NewPenagihanRepository(db *gorm.DB) PenagihanRepository {
	return &penagihanRepositoryImpl{
		DB: db,
	}
}

func (r *penagihanRepositoryImpl) InsertPenagihan(payload []*models.Penagihan) error {
	return r.DB.Create(payload).Error
}

package peminjaman_repository

import (
	"kredit_plus/models"

	"gorm.io/gorm"
)

type peminjamanRepositoryImpl struct {
	DB *gorm.DB
}

func NewPeminjaman(db *gorm.DB) PeminjamanRepository {
	return &peminjamanRepositoryImpl{
		DB: db,
	}
}

func (c *peminjamanRepositoryImpl) InsertDataPeminjaman(payload *models.Peminjaman) error {
	return c.DB.Table("peminjamans").Create(payload).Error
}

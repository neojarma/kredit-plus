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

func (c *peminjamanRepositoryImpl) UpdateDataPeminjamanSisaPembayaran(payload *models.Peminjaman) error {
	return c.DB.Table("peminjamans").Where("id_peminjaman = ?", payload.ID).Updates(map[string]interface{}{"status": payload.Status, "updated_at": payload.UpdatedAt, "sisa_pembayaran": payload.SisaPembayaran}).Error
}

func (c *peminjamanRepositoryImpl) UpdateDataPeminjaman(payload *models.Peminjaman) error {
	return c.DB.Table("peminjamans").Where("id_peminjaman = ?", payload.ID).Updates(payload).Error
}

func (c *peminjamanRepositoryImpl) GetDataPeminjaman(payload *models.Peminjaman) (*models.Peminjaman, error) {
	model := new(models.Peminjaman)
	return model, c.DB.Table("peminjamans").Find(model).Error
}

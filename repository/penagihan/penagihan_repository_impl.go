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

func (r *penagihanRepositoryImpl) UpdatePenagihan(payload *models.Penagihan) error {
	return r.DB.Model(payload).Where("id_penagihan = ?", payload.ID).Updates(payload).Error
}

func (r *penagihanRepositoryImpl) CariDataPenagihan(payload *models.Penagihan) (*models.Penagihan, error) {
	model := new(models.Penagihan)
	return model, r.DB.Where("id_penagihan = ?", payload.ID).First(model).Error
}

func (r *penagihanRepositoryImpl) GetListPenagihanByPeminjamanID(payload *models.Penagihan) ([]*models.Penagihan, error) {
	var model []*models.Penagihan
	return model, r.DB.Where("id_peminjaman = ?", payload.ID_Peminjaman).Find(&model).Error
}

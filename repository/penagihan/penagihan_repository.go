package penagihan_repository

import "kredit_plus/models"

type PenagihanRepository interface {
	InsertPenagihan(payload []*models.Penagihan) error
	UpdatePenagihan(payload *models.Penagihan) error
	CariDataPenagihan(payload *models.Penagihan) (*models.Penagihan, error)
}

package penagihan_service

import "kredit_plus/models"

type PenagihanService interface {
	BayarTagihan(payload *models.PenagihanRequest) error
	GetDataPenagihan(idPeminjaman string) ([]*models.Penagihan, error)
}

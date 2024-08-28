package penagihan_repository

import "kredit_plus/models"

type PenagihanRepository interface {
	InsertPenagihan(payload []*models.Penagihan) error
}

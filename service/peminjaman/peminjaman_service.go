package peminjaman_service

import "kredit_plus/models"

type PeminjamanService interface {
	Kredit(payload *models.PeminjamanRequest) error
}

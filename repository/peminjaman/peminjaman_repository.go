package peminjaman_repository

import "kredit_plus/models"

type PeminjamanRepository interface {
	InsertDataPeminjaman(payload *models.Peminjaman) error
}

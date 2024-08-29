package peminjaman_repository

import "kredit_plus/models"

type PeminjamanRepository interface {
	InsertDataPeminjaman(payload *models.Peminjaman) error
	UpdateDataPeminjaman(payload *models.Peminjaman) error
	UpdateDataPeminjamanSisaPembayaran(payload *models.Peminjaman) error
	GetDataPeminjaman(payload *models.Peminjaman) (*models.Peminjaman, error)
}

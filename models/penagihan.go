package models

import "time"

type Penagihan struct {
	ID                string    `gorm:"column:id_penagihan;primaryKey" json:"id"`
	ID_Peminjaman     string    `json:"id_peminjaman"`
	ID_User           string    `json:"id_user"`
	ID_Asset          string    `json:"id_asset"`
	NominalPembayaran float64   `json:"nominal_pembayaran"`
	Status            string    `json:"string"`
	TanggalPembayaran *string   `json:"tanggal_pembayaran"`
	TanggalJatuhTempo string    `json:"tanggal_jatuh_tempo"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         *string   `json:"updated_at"`
}

type PenagihanRequest struct {
	ID_Penagihan      string  `gorm:"column:id_penagihan;primaryKey" json:"id_penagihan"`
	NominalPembayaran float64 `json:"nominal_pembayaran"`
	IDUser            string  `json:"id_user"`
}

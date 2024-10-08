package models

type Peminjaman struct {
	ID                     string  `gorm:"column:id_peminjaman;primaryKey" json:"id"`
	ID_User                string  `json:"id_user"`
	ID_Assets              string  `json:"id_asset"`
	ID_Limit_Tenor         string  `json:"id_limit_tenor"`
	TotalPembayaran        float64 `json:"total_pembayaran"`
	TotalBunga             float64 `gorm:"column:total_bunga" json:"total_bunga"`
	Status                 string  `json:"status"`
	TanggalPeminjaman      string  `json:"tanggal_peminjaman"`
	TanggalAkhirPembayaran string  `json:"tanggal_akhir_pembayaran"`
	UpdatedAt              *string `json:"updated_at"`
	SisaPembayaran         float64 `gorm:"column:sisa_pembayaran" json:"sisa_pembayaran"`
	LamaTenor              int     `json:"lama_tenor"`
	OTR                    float64 `gorm:"column:otr" json:"otr"`
}

type PeminjamanRequest struct {
	ID_User   string `json:"id_user"`
	ID_Assets string `json:"id_asset"`
	Tenor     int    `json:"tenor"`
}

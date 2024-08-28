package models

type Assets struct {
	ID        string  `gorm:"column:id_assets;primaryKey" json:"id_assets"`
	NamaAsset string  `json:"nama_asset"`
	TypeAsset string  `json:"type_asset"`
	OTRAmount float64 `json:"otr_amount"`
	Bunga     float64 `json:"bunga"`
	AdminFee  float64 `json:"admin_fee"`
}

package models

type UserTenor struct {
	ID             string  `gorm:"column:id_user_tenor;primaryKey" json:"id"`
	ID_User        string  `json:"id_user"`
	ID_Limit_Tenor string  `json:"id_limit_tenor"`
	Tenor_1_Bulan  float64 `gorm:"column:1_bulan;" json:"tenor_1_bulan"`
	Tenor_2_Bulan  float64 `gorm:"column:2_bulan;" json:"tenor_2_bulan"`
	Tenor_3_Bulan  float64 `gorm:"column:3_bulan;" json:"tenor_3_bulan"`
	Tenor_6_Bulan  float64 `gorm:"column:6_bulan;" json:"tenor_6_bulan"`
}

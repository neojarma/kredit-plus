package models

type LimitTenor struct {
	ID            string  `gorm:"column:id_limit_tenor;primaryKey" json:"id"`
	Tenor_1_Bulan float64 `gorm:"column:1_bulan;" json:"tenor_1_bulan"`
	Tenor_2_Bulan float64 `gorm:"column:2_bulan;" json:"tenor_2_bulan"`
	Tenor_3_Bulan float64 `gorm:"column:3_bulan;" json:"tenor_3_bulan"`
	Tenor_6_Bulan float64 `gorm:"column:6_bulan;" json:"tenor_6_bulan"`
	SalaryMin     float64 `json:"salary_min"`
	SalaryMax     float64 `json:"salary_max"`
}

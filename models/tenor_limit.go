package models

type LimitTenor struct {
	ID            string  `gorm:"column:id_limit_tenor;primaryKey" json:"id"`
	Tenor_1_Bulan string  `json:"tenor_1_bulan"`
	Tenor_2_Bulan string  `json:"tenor_2_bulan"`
	Tenor_3_Bulan string  `json:"tenor_3_bulan"`
	Tenor_6_Bulan string  `json:"tenor_6_bulan"`
	SalaryMin     float64 `json:"salary_min"`
	SalaryMax     float64 `json:"salary_max"`
}

package models

type UserTenor struct {
	ID             string `gorm:"column:id_user_tenor;primaryKey" json:"id"`
	ID_User        string `json:"id_user"`
	ID_Limit_Tenor string `json:"id_limit_tenor"`
}

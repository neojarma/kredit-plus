package models

import "mime/multipart"

type UserRequestRegister struct {
	ID             string                `gorm:"primaryKey" json:"id"`
	Email          string                `form:"email" json:"email" validate:"required"`
	Password       string                `form:"password" json:"password" validate:"required"`
	NIK            string                `form:"nik" json:"nik" validate:"required"`
	FullName       string                `form:"full_name" json:"full_name" validate:"required"`
	LegalName      string                `form:"legal_name" json:"legal_name" validate:"required"`
	TempatLahir    string                `form:"tempat_lahir" json:"tempat_lahir" validate:"required"`
	TanggalLahir   string                `form:"tanggal_lahir" json:"tanggal_lahir" validate:"required"`
	Gaji           float64               `form:"gaji" json:"gaji" validate:"required"`
	FotoKTP        *multipart.FileHeader `form:"foto_ktp" json:"foto_ktp" validate:"required"`
	FotoSelfie     *multipart.FileHeader `form:"foto_selfie" json:"foto_selfie" validate:"required"`
	FotoKTPPath    string                `json:"foto_ktp_path"`
	FotoSelfiePath string                `json:"foto_selfie_path"`
}

type User struct {
	ID             string  `gorm:"column:id_user;primaryKey" json:"id"`
	NIK            string  `form:"nik" json:"nik" validate:"required"`
	FullName       string  `form:"full_name" json:"full_name" validate:"required"`
	LegalName      string  `form:"legal_name" json:"legal_name" validate:"required"`
	TempatLahir    string  `form:"tempat_lahir" json:"tempat_lahir" validate:"required"`
	TanggalLahir   string  `form:"tanggal_lahir" json:"tanggal_lahir" validate:"required"`
	Gaji           float64 `form:"gaji" json:"gaji" validate:"required"`
	FotoKTPPath    string  `json:"foto_ktp_path"`
	FotoSelfiePath string  `json:"foto_selfie_path"`
}

type Auths struct {
	ID       string `gorm:"column:id_auth;primaryKey" json:"id"`
	IDUser   string `json:"id_user"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	ID            string `json:"user_id"`
	AccessToken   string `json:"access_token"`
	RefreshTokens string `json:"refresh_token"`
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token"`
}

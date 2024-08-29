package user_service

import (
	"fmt"
	"io"
	"kredit_plus/helper"
	"kredit_plus/models"
	auth_repository "kredit_plus/repository/auth"
	limit_repository "kredit_plus/repository/limit"
	user_repository "kredit_plus/repository/user"
	user_tenor_repository "kredit_plus/repository/user_tenor"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
)

type userServiceImpl struct {
	AuthRepo      auth_repository.AuthRepository
	UserRepo      user_repository.UserRepository
	UserLimitRepo user_tenor_repository.UserTenor
	LimitRepo     limit_repository.TenorLimitRepository
}

func NewUserService(authRepo auth_repository.AuthRepository, userRepo user_repository.UserRepository, userLimitRepo user_tenor_repository.UserTenor, limitRepo limit_repository.TenorLimitRepository) UserService {
	return &userServiceImpl{
		AuthRepo:      authRepo,
		UserRepo:      userRepo,
		LimitRepo:     limitRepo,
		UserLimitRepo: userLimitRepo,
	}
}

func (s *userServiceImpl) Register(payload *models.UserRequestRegister) error {
	newUserID := helper.GenerateRandomString(15)
	KTPPath, err := saveUploadedFile(payload.FotoKTP, fmt.Sprintf("assets/%s/KTP/", newUserID))
	if err != nil {
		return err
	}

	SelfiePath, err := saveUploadedFile(payload.FotoSelfie, fmt.Sprintf("assets/%s/Selfie/", newUserID))
	if err != nil {
		return err
	}

	// insert to user table
	err = s.UserRepo.InsertUserProfile(
		&models.User{
			ID:             newUserID,
			NIK:            payload.NIK,
			FullName:       payload.FullName,
			LegalName:      payload.LegalName,
			TempatLahir:    payload.TempatLahir,
			TanggalLahir:   payload.TanggalLahir,
			Gaji:           payload.Gaji,
			FotoKTPPath:    KTPPath,
			FotoSelfiePath: SelfiePath,
		},
	)
	if err != nil {
		return err
	}

	// insert to auth table
	hashedPassword, err := helper.HashPassword(payload.Password)
	if err != nil {
		return err
	}

	err = s.AuthRepo.InsertToAuth(&models.Auths{
		ID:       helper.GenerateRandomString(15),
		IDUser:   newUserID,
		Email:    payload.Email,
		Password: hashedPassword,
	})
	if err != nil {
		return err
	}

	// check profile limit base on input salary
	tenorProfile, err := s.LimitRepo.GetTenorLimit(payload.Gaji)
	if err != nil {
		return err
	}

	// insert to user limit table
	err = s.UserLimitRepo.InsertUserTenorProfile(&models.UserTenor{
		ID:             helper.GenerateRandomString(15),
		ID_User:        newUserID,
		ID_Limit_Tenor: tenorProfile.ID,
		Tenor_1_Bulan:  tenorProfile.Tenor_1_Bulan,
		Tenor_2_Bulan:  tenorProfile.Tenor_2_Bulan,
		Tenor_3_Bulan:  tenorProfile.Tenor_3_Bulan,
		Tenor_6_Bulan:  tenorProfile.Tenor_6_Bulan,
	})
	if err != nil {
		return err
	}

	return nil
}

func (s *userServiceImpl) GetUser(idUser string) (*models.User, error) {
	userInfo, err := s.UserRepo.GetUserProfile(&models.User{
		ID: idUser,
	})
	if err != nil {
		return nil, err
	}

	userInfo.FotoKTPPath = strings.ReplaceAll(userInfo.FotoKTPPath, "\\", "/")
	userInfo.FotoSelfiePath = strings.ReplaceAll(userInfo.FotoSelfiePath, "\\", "/")

	return userInfo, nil
}

func saveUploadedFile(file *multipart.FileHeader, destDir string) (string, error) {
	// Ensure the destination directory exists
	if err := os.MkdirAll(destDir, os.ModePerm); err != nil {
		return "", err
	}

	// Open the uploaded file
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	// Create destination path
	dstPath := filepath.Join(destDir, file.Filename)
	dst, err := os.Create(dstPath)
	if err != nil {
		return "", err
	}
	defer dst.Close()

	// Copy the file content to the destination
	if _, err := io.Copy(dst, src); err != nil {
		return "", err
	}

	return dstPath, nil
}

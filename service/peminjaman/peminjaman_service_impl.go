package peminjaman_service

import (
	"errors"
	const_value "kredit_plus/const"
	"kredit_plus/helper"
	"kredit_plus/models"
	assets_repository "kredit_plus/repository/assets"
	peminjaman_repository "kredit_plus/repository/peminjaman"
	penagihan_repository "kredit_plus/repository/penagihan"
	user_tenor_repository "kredit_plus/repository/user_tenor"
	"time"
)

type peminjamanServiceImpl struct {
	PeminjamanRepo peminjaman_repository.PeminjamanRepository
	AssetRepo      assets_repository.AssetsRepository
	UserTenorRepo  user_tenor_repository.UserTenor
	PenagihanRepo  penagihan_repository.PenagihanRepository
}

func NewPeminjamanService(repo peminjaman_repository.PeminjamanRepository, assetRepo assets_repository.AssetsRepository, userTenorRepo user_tenor_repository.UserTenor, penagihanRepo penagihan_repository.PenagihanRepository) PeminjamanService {
	return &peminjamanServiceImpl{
		PeminjamanRepo: repo,
		PenagihanRepo:  penagihanRepo,
		UserTenorRepo:  userTenorRepo,
		AssetRepo:      assetRepo,
	}
}

func (s *peminjamanServiceImpl) GetAllPeminjaman(userID string) ([]*models.Peminjaman, error) {
	return s.PeminjamanRepo.GetAllPeminjamanUser(&models.Peminjaman{
		ID_User: userID,
	})
}

func (s *peminjamanServiceImpl) Kredit(payload *models.PeminjamanRequest) error {

	// get total price assets
	assetsInfo, err := s.AssetRepo.GetAsset(&models.Assets{
		ID: payload.ID_Assets,
	})
	if err != nil {
		return err
	}

	// read limit from user tenor
	// if insufficient balance, return
	tenorUserInfo, err := s.UserTenorRepo.GetUserTenorProfile(&models.UserTenor{
		ID_User: payload.ID_User,
	})
	if err != nil {
		return err
	}

	limitUser := getLimitPerTenor(payload.Tenor, tenorUserInfo)
	if limitUser < assetsInfo.OTRAmount {
		return errors.New("insufficient balance")
	}

	// calculate
	// totalBunga = (OTR + admin fee) * bunga * tenor
	// totalBayar = OTR + totalBunga
	// bayarPerBulan = totalBayar / tenor
	totalBunga := (assetsInfo.OTRAmount + assetsInfo.AdminFee) * float64(assetsInfo.Bunga/float64(100)) * float64(payload.Tenor)
	totalBayar := assetsInfo.OTRAmount + totalBunga
	bayarPerBulan := totalBayar / float64(payload.Tenor)

	// insert peminjaman
	newIDPeminjaman := helper.GenerateRandomString(15)
	now := time.Now()
	err = s.PeminjamanRepo.InsertDataPeminjaman(&models.Peminjaman{
		ID:                     newIDPeminjaman,
		ID_User:                payload.ID_User,
		ID_Assets:              assetsInfo.ID,
		ID_Limit_Tenor:         tenorUserInfo.ID,
		TotalPembayaran:        totalBayar,
		TotalBunga:             totalBunga,
		Status:                 const_value.STATUS_PEMINJAMAN_ACTIVE,
		TanggalPeminjaman:      now.Format(const_value.DATE_FORMAT_WITH_TIME),
		TanggalAkhirPembayaran: now.AddDate(0, payload.Tenor, 0).Format(const_value.DATE_FORMAT),
		SisaPembayaran:         totalBayar,
		UpdatedAt:              nil,
		LamaTenor:              payload.Tenor,
		OTR:                    assetsInfo.OTRAmount,
	})
	if err != nil {
		return err
	}

	// insert penagihan
	dataPenagihan := make([]*models.Penagihan, 0)
	tempDate := time.Now()
	for i := 0; i < payload.Tenor; i++ {
		newDataPenagihan := models.Penagihan{
			ID:                helper.GenerateRandomString(15),
			ID_Peminjaman:     newIDPeminjaman,
			ID_User:           payload.ID_User,
			ID_Asset:          assetsInfo.ID,
			NominalPembayaran: bayarPerBulan,
			Status:            const_value.STATUS_PENAGIHAN_UNPAID,
			TanggalJatuhTempo: tempDate.AddDate(0, 1, 0).Format(const_value.DATE_FORMAT),
			CreatedAt:         time.Now(),
			TanggalPembayaran: nil,
			UpdatedAt:         nil,
		}
		dataPenagihan = append(dataPenagihan, &newDataPenagihan)
		tempDate = tempDate.AddDate(0, 1, 0)
	}
	if err := s.PenagihanRepo.InsertPenagihan(dataPenagihan); err != nil {
		return err
	}

	// decrease user tenor
	if err := s.UserTenorRepo.UpdateUserTenorProfile(updateValueUserTenor(payload.Tenor, tenorUserInfo, assetsInfo.OTRAmount)); err != nil {
		return nil
	}

	return nil
}

func updateValueUserTenor(tenor int, tenorUser *models.UserTenor, otr float64) *models.UserTenor {
	switch tenor {
	case 1:
		return &models.UserTenor{
			ID_User:       tenorUser.ID_User,
			Tenor_1_Bulan: tenorUser.Tenor_1_Bulan - otr,
		}
	case 2:
		return &models.UserTenor{
			ID_User:       tenorUser.ID_User,
			Tenor_2_Bulan: tenorUser.Tenor_2_Bulan - otr,
		}
	case 3:
		return &models.UserTenor{
			ID_User:       tenorUser.ID_User,
			Tenor_3_Bulan: tenorUser.Tenor_3_Bulan - otr,
		}
	case 6:
		return &models.UserTenor{
			ID_User:       tenorUser.ID_User,
			Tenor_6_Bulan: tenorUser.Tenor_6_Bulan - otr,
		}
	}

	return nil
}

func getLimitPerTenor(month int, tenorProfile *models.UserTenor) float64 {
	switch month {
	case 1:
		return tenorProfile.Tenor_1_Bulan
	case 2:
		return tenorProfile.Tenor_2_Bulan
	case 3:
		return tenorProfile.Tenor_3_Bulan
	case 6:
		return tenorProfile.Tenor_6_Bulan
	}

	return 0
}

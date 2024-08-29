package penagihan_service

import (
	"errors"
	const_value "kredit_plus/const"
	"kredit_plus/models"
	peminjaman_repository "kredit_plus/repository/peminjaman"
	penagihan_repository "kredit_plus/repository/penagihan"
	user_tenor_repository "kredit_plus/repository/user_tenor"
	"time"
)

type penagihanServiceImpl struct {
	PenagihanRepository  penagihan_repository.PenagihanRepository
	PeminjamanRepository peminjaman_repository.PeminjamanRepository
	UserTenorRepository  user_tenor_repository.UserTenor
}

func NewPenagihanService(penagihanRepo penagihan_repository.PenagihanRepository, peminjamanRepo peminjaman_repository.PeminjamanRepository, userTenor user_tenor_repository.UserTenor) PenagihanService {
	return &penagihanServiceImpl{
		PenagihanRepository:  penagihanRepo,
		PeminjamanRepository: peminjamanRepo,
		UserTenorRepository:  userTenor,
	}
}

func (s *penagihanServiceImpl) GetDataPenagihan(idPeminjaman string) ([]*models.Penagihan, error) {
	return s.PenagihanRepository.GetListPenagihanByPeminjamanID(&models.Penagihan{
		ID_Peminjaman: idPeminjaman,
	})
}

func (s *penagihanServiceImpl) BayarTagihan(payload *models.PenagihanRequest) error {

	// get data penagihan
	penagihanInfo, err := s.PenagihanRepository.CariDataPenagihan(&models.Penagihan{
		ID: payload.ID_Penagihan,
	})
	if err != nil {
		return err
	}

	// check nominal pembayaran
	if penagihanInfo.NominalPembayaran != payload.NominalPembayaran {
		return errors.New("nominal pembayaran tidak sesuai")
	}

	// get peminjaman
	// cek selisih setelah bayar
	peminjamanInfo, err := s.PeminjamanRepository.GetDataPeminjaman(&models.Peminjaman{
		ID: penagihanInfo.ID_Peminjaman,
	})
	if err != nil {
		return err
	}

	isLunas := peminjamanInfo.SisaPembayaran-payload.NominalPembayaran == 0
	newDate := time.Now().Format(const_value.DATE_FORMAT_WITH_TIME)

	// update peminjaman
	if isLunas {

		err := s.PeminjamanRepository.UpdateDataPeminjamanSisaPembayaran(&models.Peminjaman{
			ID:             penagihanInfo.ID_Peminjaman,
			Status:         const_value.STATUS_PEMINJAMAN_LUNAS,
			UpdatedAt:      &newDate,
			SisaPembayaran: peminjamanInfo.SisaPembayaran - payload.NominalPembayaran,
		})
		if err != nil {
			return err
		}

		// update user tenor limit
		userTenorInfo, err := s.UserTenorRepository.GetUserTenorProfile(&models.UserTenor{
			ID_User: payload.IDUser,
		})
		if err != nil {
			return err
		}

		if err := s.UserTenorRepository.UpdateUserTenorProfile(updateValueUserTenor(peminjamanInfo.LamaTenor, userTenorInfo, peminjamanInfo.OTR)); err != nil {
			return err
		}
	} else {
		err := s.PeminjamanRepository.UpdateDataPeminjaman(&models.Peminjaman{
			ID:             penagihanInfo.ID_Peminjaman,
			UpdatedAt:      &newDate,
			SisaPembayaran: peminjamanInfo.SisaPembayaran - payload.NominalPembayaran,
		})
		if err != nil {
			return err
		}
	}

	// update penagihan
	if err := s.PenagihanRepository.UpdatePenagihan(&models.Penagihan{
		ID:                penagihanInfo.ID,
		Status:            const_value.STATUS_PENAGIHAN_PAID,
		UpdatedAt:         &newDate,
		TanggalPembayaran: &newDate,
	}); err != nil {
		return err
	}

	return nil
}

func updateValueUserTenor(tenor int, tenorUser *models.UserTenor, otr float64) *models.UserTenor {
	switch tenor {
	case 1:
		return &models.UserTenor{
			ID_User:       tenorUser.ID_User,
			Tenor_1_Bulan: tenorUser.Tenor_1_Bulan + otr,
		}
	case 2:
		return &models.UserTenor{
			ID_User:       tenorUser.ID_User,
			Tenor_2_Bulan: tenorUser.Tenor_2_Bulan + otr,
		}
	case 3:
		return &models.UserTenor{
			ID_User:       tenorUser.ID_User,
			Tenor_3_Bulan: tenorUser.Tenor_3_Bulan + otr,
		}
	case 6:
		return &models.UserTenor{
			ID_User:       tenorUser.ID_User,
			Tenor_6_Bulan: tenorUser.Tenor_6_Bulan + otr,
		}
	}

	return nil
}

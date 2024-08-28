package limit_repository

import "kredit_plus/models"

type TenorLimitRepository interface {
	GetTenorLimit(salary float64) (*models.LimitTenor, error)
}

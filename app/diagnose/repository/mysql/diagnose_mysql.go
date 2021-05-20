package mysql

import (
	"context"
	"github.com/huf0813/scade_backend_api/domain"
	"gorm.io/gorm"
)

type DiagnoseRepoMysql struct {
	DB *gorm.DB
}

func NewDiagnoseRepoMysql(db *gorm.DB) domain.DiagnoseRepository {
	return &DiagnoseRepoMysql{DB: db}
}

func (d *DiagnoseRepoMysql) GetDiagnoses(ctx context.Context, email string) ([]domain.Diagnose, error) {
	var diagnoses []domain.Diagnose
	if err := d.DB.
		WithContext(ctx).
		Where("users.email = ?", email).
		Find(&diagnoses).Error; err != nil {
		return nil, err
	}
	return diagnoses, nil
}

func (d *DiagnoseRepoMysql) GetDiagnoseByID(ctx context.Context, email string, diagnoseID int) (domain.Diagnose, error) {
	var diagnose domain.Diagnose
	if err := d.DB.
		WithContext(ctx).
		Where("users.email = ?", email).
		Where("diagnoses.id = ?", diagnoseID).
		Find(&diagnose).Error; err != nil {
		return domain.Diagnose{}, err
	}
	return diagnose, nil
}

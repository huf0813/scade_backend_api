package mysql

import (
	"context"
	"errors"
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
		Joins("JOIN users ON diagnoses.user_id = users.id").
		Where("users.email = ?", email).
		Find(&diagnoses).Error; err != nil {
		return nil, err
	}
	return diagnoses, nil
}

func (d *DiagnoseRepoMysql) GetDiagnoseByID(ctx context.Context, email string, diagnoseID uint) (domain.Diagnose, error) {
	var diagnose domain.Diagnose
	if err := d.DB.
		WithContext(ctx).
		Joins("JOIN users ON diagnoses.user_id = users.id").
		Where("users.email = ?", email).
		Where("diagnoses.id = ?", diagnoseID).
		Find(&diagnose).Error; err != nil {
		return domain.Diagnose{}, err
	}
	return diagnose, nil
}

func (d *DiagnoseRepoMysql) CreateDiagnose(ctx context.Context, diagnose *domain.Diagnose) (uint, error) {
	create := domain.Diagnose{
		CancerName:  diagnose.CancerName,
		CancerImage: diagnose.CancerImage,
		Position:    diagnose.Position,
		Price:       diagnose.Price,
		UserID:      diagnose.UserID,
	}

	result := d.DB.
		WithContext(ctx).
		Create(&create)
	lastID := create.ID
	if err := result.Error; err != nil {
		return 0, err
	}
	rows := result.RowsAffected
	if rows <= 0 {
		return 0, errors.New("failed to insert data, empty feedback")
	}
	return lastID, nil
}

package mysql

import (
	"context"
	"github.com/huf0813/scade_backend_api/domain"
	"gorm.io/gorm"
)

type HospitalRepoMysql struct {
	DB *gorm.DB
}

func NewHospitalRepoMysql(db *gorm.DB) domain.HospitalRepository {
	return &HospitalRepoMysql{DB: db}
}

func (h *HospitalRepoMysql) GetHospitals(ctx context.Context) ([]domain.Hospital, error) {
	var hospitals []domain.Hospital
	if err := h.DB.
		WithContext(ctx).
		Find(&hospitals).
		Error; err != nil {
		return nil, err
	}
	return hospitals, nil
}

func (h *HospitalRepoMysql) GetHospitalsByCity(ctx context.Context, city string) ([]domain.Hospital, error) {
	var hospitals []domain.Hospital
	if err := h.DB.
		WithContext(ctx).
		Where("region LIKE ?", city+"%").
		Find(&hospitals).
		Error; err != nil {
		return nil, err
	}
	return hospitals, nil
}

func (h *HospitalRepoMysql) GetHospitalByID(ctx context.Context, id int) (domain.Hospital, error) {
	var hospital domain.Hospital
	if err := h.DB.
		WithContext(ctx).
		First(&hospital, id).
		Error; err != nil {
		return domain.Hospital{}, err
	}
	return hospital, nil
}

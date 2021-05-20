package usecase

import (
	"context"
	"github.com/huf0813/scade_backend_api/domain"
	"time"
)

type HospitalUseCase struct {
	hospitalRepoMysql domain.HospitalRepository
	timeOut           time.Duration
}

func NewHospitalUseCase(h domain.HospitalRepository, timeOut time.Duration) domain.HospitalUseCase {
	return &HospitalUseCase{
		hospitalRepoMysql: h,
		timeOut:           timeOut,
	}
}

func (h *HospitalUseCase) GetHospitals(ctx context.Context) ([]domain.Hospital, error) {
	result, err := h.hospitalRepoMysql.GetHospitals(ctx)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (h *HospitalUseCase) GetHospitalsByCity(ctx context.Context, city string) ([]domain.Hospital, error) {
	ctx, cancel := context.WithTimeout(ctx, h.timeOut)
	defer cancel()

	result, err := h.hospitalRepoMysql.GetHospitalsByCity(ctx, city)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (h *HospitalUseCase) GetHospitalByID(ctx context.Context, id int) (domain.Hospital, error) {
	ctx, cancel := context.WithTimeout(ctx, h.timeOut)
	defer cancel()

	result, err := h.hospitalRepoMysql.GetHospitalByID(ctx, id)
	if err != nil {
		return domain.Hospital{}, err
	}
	return result, nil
}

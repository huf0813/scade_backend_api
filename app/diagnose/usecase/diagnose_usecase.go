package usecase

import (
	"context"
	"github.com/huf0813/scade_backend_api/domain"
	"time"
)

type DiagnoseUseCase struct {
	diagnoseRepoMysql domain.DiagnoseRepository
	timeOut           time.Duration
}

func NewDiagnoseUseCase(d domain.DiagnoseRepository, timeOut time.Duration) domain.DiagnoseUseCase {
	return &DiagnoseUseCase{
		diagnoseRepoMysql: d,
		timeOut:           timeOut,
	}
}

func (d *DiagnoseUseCase) GetDiagnoses(ctx context.Context, email string) ([]domain.Diagnose, error) {
	ctx, cancel := context.WithTimeout(ctx, d.timeOut)
	defer cancel()

	result, err := d.diagnoseRepoMysql.GetDiagnoses(ctx, email)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (d *DiagnoseUseCase) GetDiagnoseByID(ctx context.Context, email string, diagnoseID int) (domain.Diagnose, error) {
	ctx, cancel := context.WithTimeout(ctx, d.timeOut)
	defer cancel()

	result, err := d.diagnoseRepoMysql.GetDiagnoseByID(ctx, email, diagnoseID)
	if err != nil {
		return domain.Diagnose{}, err
	}

	return result, nil
}

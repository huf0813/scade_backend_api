package usecase

import (
	"context"
	"fmt"
	"github.com/huf0813/scade_backend_api/domain"
	"github.com/huf0813/scade_backend_api/utils/file_upload"
	"mime/multipart"
	"os"
	"time"
)

type DiagnoseUseCase struct {
	diagnoseRepoMysql domain.DiagnoseRepository
	userRepository    domain.UserRepository
	timeOut           time.Duration
}

func NewDiagnoseUseCase(d domain.DiagnoseRepository, u domain.UserRepository, timeOut time.Duration) domain.DiagnoseUseCase {
	return &DiagnoseUseCase{
		diagnoseRepoMysql: d,
		userRepository:    u,
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

func (d *DiagnoseUseCase) GetDiagnoseByID(ctx context.Context, email string, diagnoseID int) (
	domain.Diagnose,
	error) {
	ctx, cancel := context.WithTimeout(ctx, d.timeOut)
	defer cancel()

	result, err := d.diagnoseRepoMysql.GetDiagnoseByID(ctx, email, diagnoseID)
	if err != nil {
		return domain.Diagnose{}, err
	}

	return result, nil
}

func (d *DiagnoseUseCase) CreateDiagnose(ctx context.Context,
	diagnose *domain.DiagnoseRequest,
	fileHeader *multipart.FileHeader) error {
	ctx, cancel := context.WithTimeout(ctx, d.timeOut)
	defer cancel()

	path := fmt.Sprintf("%s/%s", "assets", "skin_image")
	result, err := file_upload.NewFileUpload(path, fileHeader)
	if err != nil {
		return err
	}

	user, err := d.userRepository.GetUserByEmail(ctx, diagnose.UserEmail)
	if err != nil {
		if err := os.Remove(result); err != nil {
			return err
		}
		return err
	}

	create := domain.Diagnose{
		CancerName:  diagnose.CancerName,
		CancerImage: fileHeader.Filename,
		Position:    diagnose.Position,
		Price:       diagnose.Price,
		UserID:      user.ID,
	}
	if err := d.diagnoseRepoMysql.CreateDiagnose(ctx, &create); err != nil {
		return err
	}

	return nil
}

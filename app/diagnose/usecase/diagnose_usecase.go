package usecase

import (
	"context"
	"fmt"
	"github.com/huf0813/scade_backend_api/domain"
	"github.com/huf0813/scade_backend_api/utils/file_upload"
	"mime/multipart"
	"time"
)

type DiagnoseUseCase struct {
	diagnoseRepoMysql      domain.DiagnoseRepository
	userRepository         domain.UserRepository
	subscriptionRepository domain.SubscriptionRepository
	timeOut                time.Duration
}

func NewDiagnoseUseCase(d domain.DiagnoseRepository,
	u domain.UserRepository,
	s domain.SubscriptionRepository,
	timeOut time.Duration) domain.DiagnoseUseCase {
	return &DiagnoseUseCase{
		diagnoseRepoMysql:      d,
		userRepository:         u,
		subscriptionRepository: s,
		timeOut:                timeOut,
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

func (d *DiagnoseUseCase) GetDiagnoseByID(ctx context.Context, email string, diagnoseID uint) (
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
	fileHeader *multipart.FileHeader) (uint, error) {
	ctx, cancel := context.WithTimeout(ctx, d.timeOut)
	defer cancel()

	path := fmt.Sprintf("%s/%s", "assets", "skin_image")
	filename, err := file_upload.NewFileUpload(path, fileHeader)
	if err != nil {
		return 0, err
	}

	user, err := d.userRepository.GetUserByEmail(ctx, diagnose.UserEmail)
	if err != nil {
		return 0, err
	}

	price := 0
	isActiveSubscription, err := d.subscriptionRepository.CheckSubscription(ctx, user.ID)
	if err != nil {
		return 0, err
	}
	if !isActiveSubscription {
		price = diagnose.Price
	}

	create := domain.Diagnose{
		CancerName:  diagnose.CancerName,
		CancerImage: filename,
		Position:    diagnose.Position,
		Price:       price,
		UserID:      user.ID,
	}
	lastID, err := d.diagnoseRepoMysql.CreateDiagnose(ctx, &create)
	if err != nil {
		return 0, err
	}

	return lastID, nil
}

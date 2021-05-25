package usecase

import (
	"context"
	"github.com/huf0813/scade_backend_api/domain"
	"time"
)

type InvoiceUseCase struct {
	invoiceRepoMysql  domain.InvoiceRepository
	userRepoMysql     domain.UserRepository
	diagnoseRepoMysql domain.DiagnoseRepository
	timeOut           time.Duration
}

func NewInvoiceUseCase(i domain.InvoiceRepository,
	u domain.UserRepository,
	d domain.DiagnoseRepository,
	timeOut time.Duration) domain.InvoiceUseCase {
	return &InvoiceUseCase{
		invoiceRepoMysql:  i,
		userRepoMysql:     u,
		diagnoseRepoMysql: d,
		timeOut:           timeOut,
	}
}

func (i *InvoiceUseCase) GetInvoices(ctx context.Context,
	email string) ([]domain.Invoice,
	error) {
	ctx, cancel := context.WithTimeout(ctx, i.timeOut)
	defer cancel()

	user, err := i.userRepoMysql.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	result, err := i.invoiceRepoMysql.GetInvoices(ctx, user.ID)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (i *InvoiceUseCase) GetInvoiceByID(ctx context.Context,
	invoiceID int,
	email string) (domain.Invoice,
	error) {
	ctx, cancel := context.WithTimeout(ctx, i.timeOut)
	defer cancel()

	user, err := i.userRepoMysql.GetUserByEmail(ctx, email)
	if err != nil {
		return domain.Invoice{}, err
	}

	result, err := i.invoiceRepoMysql.GetInvoiceByID(ctx, invoiceID, user.ID)
	if err != nil {
		return domain.Invoice{}, err
	}
	return result, nil
}

func (i *InvoiceUseCase) CreateInvoice(ctx context.Context,
	create *domain.InvoiceRequest,
	email string) error {
	ctx, cancel := context.WithTimeout(ctx, i.timeOut)
	defer cancel()

	result, err := i.diagnoseRepoMysql.GetDiagnoseByID(ctx, email, create.DiagnoseID)
	if err != nil {
		return err
	}
	create.DiagnoseID = result.ID

	if err := i.invoiceRepoMysql.CreateInvoice(ctx, create); err != nil {
		return err
	}

	return nil
}

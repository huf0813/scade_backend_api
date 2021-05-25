package mysql

import (
	"context"
	"github.com/huf0813/scade_backend_api/domain"
	"gorm.io/gorm"
)

type InvoiceRepoMysql struct {
	DB *gorm.DB
}

func NewInvoiceRepoMysql(db *gorm.DB) domain.InvoiceRepository {
	return &InvoiceRepoMysql{DB: db}
}

func (i *InvoiceRepoMysql) GetInvoices(ctx context.Context,
	userID uint) ([]domain.Invoice,
	error) {
	var invoices []domain.Invoice

	if err := i.DB.
		WithContext(ctx).
		Joins("JOIN diagnoses ON diagnoses.id = invoices.diagnose_id").
		Where("diagnoses.user_id = ?", userID).
		Find(&invoices).Error; err != nil {
		return nil, err
	}

	return invoices, nil
}

func (i *InvoiceRepoMysql) GetInvoiceByID(ctx context.Context,
	invoiceID int,
	userID uint) (domain.Invoice,
	error) {
	var invoice domain.Invoice

	if err := i.DB.
		WithContext(ctx).
		Joins("JOIN diagnoses ON diagnoses.id = invoices.diagnose_id").
		Where("diagnoses.user_id = ?", userID).
		First(&invoice, invoiceID).Error; err != nil {
		return domain.Invoice{}, err
	}

	return invoice, nil
}

func (i *InvoiceRepoMysql) CreateInvoice(ctx context.Context,
	req *domain.InvoiceRequest) error {
	create := domain.Invoice{
		HospitalID: req.HospitalID,
		DiagnoseID: req.DiagnoseID,
	}

	if err := i.DB.
		WithContext(ctx).
		Create(&create).Error; err != nil {
		return err
	}

	return nil
}

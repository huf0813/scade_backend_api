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
	userID uint) ([]domain.InvoiceResponse,
	error) {
	var invoices []domain.InvoiceResponse

	if err := i.DB.
		WithContext(ctx).
		Model(&domain.Invoice{}).
		Select("invoices.id as invoice_id, hospitals.name as hospital_name, hospitals.address as hospital_address, hospitals.phone as hospital_phone, hospitals.city as hospital_city, hospitals.province as hospital_province, diagnoses.cancer_name, diagnoses.cancer_image, diagnoses.position as cancer_position").
		Joins("JOIN diagnoses ON diagnoses.id = invoices.diagnose_id").
		Joins("JOIN hospitals ON hospitals.id = invoices.hospital_id").
		Where("diagnoses.user_id = ?", userID).
		Scan(&invoices).Error; err != nil {
		return nil, err
	}

	return invoices, nil
}

func (i *InvoiceRepoMysql) GetInvoiceByID(ctx context.Context,
	invoiceID int,
	userID uint) (domain.InvoiceResponse,
	error) {
	var invoice domain.InvoiceResponse

	if err := i.DB.
		WithContext(ctx).
		Model(&domain.Invoice{}).
		Select("invoices.id as invoice_id, hospitals.name as hospital_name, hospitals.address as hospital_address, hospitals.phone as hospital_phone, hospitals.city as hospital_city, hospitals.province as hospital_province, diagnoses.cancer_name, diagnoses.cancer_image, diagnoses.position as cancer_position").
		Joins("JOIN diagnoses ON diagnoses.id = invoices.diagnose_id").
		Joins("JOIN hospitals ON hospitals.id = invoices.hospital_id").
		Where("diagnoses.user_id = ?", userID).
		Where("invoices.id = ?", invoiceID).
		Scan(&invoice).Error; err != nil {
		return domain.InvoiceResponse{}, err
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

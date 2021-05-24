package domain

import (
	"context"
	"gorm.io/gorm"
)

type Invoice struct {
	gorm.Model
	HospitalID uint `gorm:"not_null" json:"hospital_id"`
	DiagnoseID uint `gorm:"not_null" json:"diagnose_id"`
}

type InvoiceRequest struct {
	HospitalID uint `json:"hospital_id"`
	DiagnoseID uint `json:"diagnose_id"`
}

type InvoiceRepository interface {
	GetInvoices(ctx context.Context, userID uint) ([]Invoice, error)
	GetInvoiceByID(ctx context.Context, invoiceID int, userID uint) (Invoice, error)
	CreateInvoice(ctx context.Context, req *InvoiceRequest) error
}

type InvoiceUseCase interface {
	GetInvoices(ctx context.Context, email string) ([]Invoice, error)
	GetInvoiceByID(ctx context.Context, invoiceID int, email string) (Invoice, error)
	CreateInvoice(ctx context.Context, create *InvoiceRequest, email string) error
}

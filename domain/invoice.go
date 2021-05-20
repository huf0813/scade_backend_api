package domain

import (
	"context"
	"gorm.io/gorm"
)

type Invoice struct {
	gorm.Model
	UserID     uint `gorm:"not_null"`
	HospitalID uint `gorm:"not_null"`
	DiagnoseID uint `gorm:"not_null"`
}

type InvoiceRepository interface {
	GetInvoices(ctx context.Context, userID int) ([]Invoice, error)
	GetInvoiceByID(ctx context.Context, userID, invoiceID int) (Invoice, error)
}

type InvoiceUseCase interface {
	GetInvoices(ctx context.Context, userID int) ([]Invoice, error)
	GetInvoiceByID(ctx context.Context, userID, invoiceID int) (Invoice, error)
}
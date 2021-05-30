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

type InvoiceUpdateHospitalRequest struct {
	HospitalID uint `json:"hospital_id" validate:"required"`
}

type InvoiceResponse struct {
	InvoiceID        uint   `json:"invoice_id"`
	HospitalName     string `json:"hospital_name"`
	HospitalAddress  string `json:"hospital_address"`
	HospitalPhone    string `json:"hospital_phone"`
	HospitalCity     string `json:"hospital_city"`
	HospitalProvince string `json:"hospital_province"`
	CancerName       string `json:"cancer_name"`
	CancerImage      string `json:"cancer_image"`
	CancerPosition   string `json:"cancer_position"`
	InvoiceCreatedAt string `json:"invoice_created_at"`
	InvoiceUpdatedAt string `json:"invoice_updated_at"`
}

type InvoiceRepository interface {
	GetInvoices(ctx context.Context, userID uint) ([]InvoiceResponse, error)
	GetInvoiceByID(ctx context.Context, invoiceID int, userID uint) (InvoiceResponse, error)
	CreateInvoice(ctx context.Context, req *InvoiceRequest) error
	UpdateInvoice(ctx context.Context, req *InvoiceRequest, invoiceID int) error
}

type InvoiceUseCase interface {
	GetInvoices(ctx context.Context, email string) ([]InvoiceResponse, error)
	GetInvoiceByID(ctx context.Context, invoiceID int, email string) (InvoiceResponse, error)
	CreateInvoice(ctx context.Context, create *InvoiceRequest, email string) error
	UpdateInvoice(ctx context.Context, update *InvoiceUpdateHospitalRequest, email string, invoiceID int) error
}

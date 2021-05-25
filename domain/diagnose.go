package domain

import (
	"context"
	"gorm.io/gorm"
	"mime/multipart"
)

type Diagnose struct {
	gorm.Model
	CancerName  string    `gorm:"not_null" json:"cancer_name"`
	CancerImage string    `gorm:"not_null" json:"cancer_image"`
	Position    string    `gorm:"not_null" json:"position"`
	Price       int       `gorm:"not_null" json:"price"`
	UserID      uint      `json:"user_id"`
	Invoices    []Invoice `gorm:"foreignKey:DiagnoseID" json:"invoices"`
}

type DiagnoseRequest struct {
	CancerName  string `json:"cancer_name"`
	CancerImage string `json:"cancer_image"`
	Position    string `json:"position"`
	Price       int    `json:"price"`
	UserEmail   string `json:"user_email"`
}

type DiagnoseRepository interface {
	GetDiagnoses(ctx context.Context, email string) ([]Diagnose, error)
	GetDiagnoseByID(ctx context.Context, email string, diagnoseID uint) (Diagnose, error)
	CreateDiagnose(ctx context.Context, diagnose *Diagnose) (uint, error)
}

type DiagnoseUseCase interface {
	GetDiagnoses(ctx context.Context, email string) ([]Diagnose, error)
	GetDiagnoseByID(ctx context.Context, email string, diagnoseID uint) (Diagnose, error)
	CreateDiagnose(ctx context.Context,
		diagnose *DiagnoseRequest,
		fileHeader *multipart.FileHeader) (uint, error)
}

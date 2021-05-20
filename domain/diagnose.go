package domain

import (
	"context"
	"gorm.io/gorm"
)

type Diagnose struct {
	gorm.Model
	CancerName string `gorm:"not_null"`
	Position   string `gorm:"not_null"`
	DiagnoseID uint   `json:"diagnose_id"`
}

type DiagnoseRepository interface {
	GetDiagnoses(ctx context.Context, email string) ([]Diagnose, error)
	GetDiagnoseByID(ctx context.Context, email string, diagnoseID int) (Diagnose, error)
}

type DiagnoseUseCase interface {
	GetDiagnoses(ctx context.Context, email string) ([]Diagnose, error)
	GetDiagnoseByID(ctx context.Context, email string, diagnoseID int) (Diagnose, error)
}

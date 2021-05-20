package domain

import (
	"context"
	"gorm.io/gorm"
)

type Diagnose struct {
	gorm.Model
	CancerName string `gorm:"not_null"`
	Position   string `gorm:"not_null"`
}

type DiagnoseRepository interface {
	GetDiagnoses(ctx context.Context, userID int) ([]Diagnose, error)
	GetDiagnoseByID(ctx context.Context, userID, diagnoseID int) (Diagnose, error)
}

type DiagnoseUseCase interface {
	GetDiagnoses(ctx context.Context, userID int) ([]Diagnose, error)
	GetDiagnoseByID(ctx context.Context, userID, diagnoseID int) (Diagnose, error)
}

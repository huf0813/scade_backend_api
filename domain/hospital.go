package domain

import (
	"context"
	"gorm.io/gorm"
)

type Hospital struct {
	gorm.Model
	Name     string `gorm:"not_null"`
	Address  string `gorm:"not_null"`
	Phone    string `gorm:"not_null"`
	City     string `gorm:"not_null"`
	Province string `gorm:"not_null"`
}

type HospitalRepository interface {
	GetHospitals(ctx context.Context) ([]Hospital, error)
	GetHospitalsByCity(ctx context.Context, city string) ([]Hospital, error)
	GetHospitalByID(ctx context.Context, id int) (Hospital, error)
}

type HospitalUseCase interface {
	GetHospitals(ctx context.Context) ([]Hospital, error)
	GetHospitalsByCity(ctx context.Context, city string) ([]Hospital, error)
	GetHospitalByID(ctx context.Context, id int) (Hospital, error)
}
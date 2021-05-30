package domain

import (
	"context"
	"gorm.io/gorm"
)

type Hospital struct {
	gorm.Model
	Name     string    `gorm:"not_null" json:"name"`
	Address  string    `gorm:"not_null" json:"address"`
	Phone    string    `gorm:"not_null" json:"phone"`
	Region   string    `gorm:"not_null" json:"region"`
	Province string    `gorm:"not_null" json:"province"`
	Invoices []Invoice `gorm:"foreignKey:HospitalID" json:"invoices"`
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

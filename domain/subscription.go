package domain

import (
	"context"
	"gorm.io/gorm"
	"time"
)

type Subscription struct {
	gorm.Model
	Price    int       `gorm:"not_null" json:"price"`
	FinishAt time.Time `gorm:"not_null" json:"finish_at"`
	UserID   uint      `gorm:"not_null" json:"user_id"`
}

type SubscriptionRepository interface {
	GetSubscription(ctx context.Context, email string) ([]Subscription, error)
	GetSubscriptionByID(ctx context.Context, email string, subscriptionID int) (Subscription, error)
	CreateSubscriptionByUser(ctx context.Context, email uint) error
	CheckSubscription(ctx context.Context, userID uint) (bool, error)
}

type SubscriptionUseCase interface {
	GetSubscriptionByUser(ctx context.Context, email string) ([]Subscription, error)
	GetSubscriptionByID(ctx context.Context, email string, subscriptionID int) (Subscription, error)
	CreateSubscriptionByUser(ctx context.Context, email string) error
	CheckActiveSubscription(ctx context.Context, email string) (bool, error)
}

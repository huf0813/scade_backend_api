package mysql

import (
	"context"
	"github.com/huf0813/scade_backend_api/domain"
	"gorm.io/gorm"
	"time"
)

type SubscriptionRepoMysql struct {
	DB *gorm.DB
}

func NewSubscriptionRepoMysql(db *gorm.DB) domain.SubscriptionRepository {
	return &SubscriptionRepoMysql{DB: db}
}

func (s SubscriptionRepoMysql) GetSubscription(ctx context.Context, email string) ([]domain.Subscription, error) {
	var subs []domain.Subscription
	if err := s.DB.
		WithContext(ctx).
		Joins("JOIN users ON subscriptions.user_id = users.id").
		Where("users.email = ?", email).
		Find(&subs).Error; err != nil {
		return nil, err
	}
	return subs, nil
}

func (s SubscriptionRepoMysql) GetSubscriptionByID(ctx context.Context,
	email string,
	subscriptionID int) (domain.Subscription, error) {
	var subs domain.Subscription
	if err := s.DB.
		WithContext(ctx).
		Joins("JOIN users ON subscriptions.user_id = users.id").
		Where("users.email = ?", email).
		First(&subs, subscriptionID).Error; err != nil {
		return domain.Subscription{}, err
	}
	return subs, nil
}

func (s SubscriptionRepoMysql) CreateSubscriptionByUser(ctx context.Context, userID uint) error {
	subs := domain.Subscription{
		Price: 10,
		// add subscription for 3 month
		FinishAt: time.Now().Add(2160 * time.Hour),
		UserID:   userID,
	}
	if err := s.DB.
		WithContext(ctx).
		Create(&subs).Error; err != nil {
		return err
	}
	return nil
}

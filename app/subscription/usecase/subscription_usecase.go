package usecase

import (
	"context"
	"github.com/huf0813/scade_backend_api/domain"
	"time"
)

type SubscriptionUseCase struct {
	subscriptionRepoMysql domain.SubscriptionRepository
	userRepoMysql         domain.UserRepository
	timeOut               time.Duration
}

func NewSubscriptionUseCase(s domain.SubscriptionRepository,
	u domain.UserRepository,
	timeOut time.Duration) domain.SubscriptionUseCase {
	return &SubscriptionUseCase{
		subscriptionRepoMysql: s,
		userRepoMysql:         u,
		timeOut:               timeOut,
	}
}

func (s *SubscriptionUseCase) GetSubscriptionByUser(ctx context.Context, email string) ([]domain.Subscription, error) {
	res, err := s.subscriptionRepoMysql.GetSubscription(ctx, email)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *SubscriptionUseCase) GetSubscriptionByID(ctx context.Context, email string, subscriptionID int) (domain.Subscription, error) {
	ctx, cancel := context.WithTimeout(ctx, s.timeOut)
	defer cancel()

	res, err := s.subscriptionRepoMysql.GetSubscriptionByID(ctx, email, subscriptionID)
	if err != nil {
		return domain.Subscription{}, err
	}

	return res, nil
}

func (s *SubscriptionUseCase) CreateSubscriptionByUser(ctx context.Context, email string) error {
	ctx, cancel := context.WithTimeout(ctx, s.timeOut)
	defer cancel()

	user, err := s.userRepoMysql.GetUserByEmail(ctx, email)
	if err != nil {
		return err
	}

	if err := s.subscriptionRepoMysql.CreateSubscriptionByUser(ctx, user.ID); err != nil {
		return err
	}

	return nil
}

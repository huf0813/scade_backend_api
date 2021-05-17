package domain

import "context"

type Migration struct{}

type MigrationRepository interface {
	Migrate(ctx context.Context) error
	Seed(ctx context.Context) error
}

type MigrationUseCase interface {
	Migrate(ctx context.Context) error
	Seed(ctx context.Context) error
}

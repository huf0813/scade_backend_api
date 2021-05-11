package mysql

import (
	"context"
	"github.com/huf0813/scade_backend_api/domain"
	"gorm.io/gorm"
)

type MigrationRepoMysql struct {
	DB *gorm.DB
}

func NewMigrationRepoMysql(conn *gorm.DB) domain.MigrationRepository {
	return &MigrationRepoMysql{DB: conn}
}

func (m *MigrationRepoMysql) Migrate(ctx context.Context) error {
	if err := m.DB.
		WithContext(ctx).
		Migrator().
		DropTable(&domain.Article{}); err != nil {
		return err
	}
	if err := m.DB.
		WithContext(ctx).
		Set("gorm:table_options", "ENGINE=InnoDB").
		Migrator().
		CreateTable(&domain.Article{}); err != nil {
		return err
	}
	return nil
}

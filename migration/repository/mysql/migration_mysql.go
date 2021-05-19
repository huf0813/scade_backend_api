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
	// layer two
	if err := m.DB.
		WithContext(ctx).
		Migrator().
		DropTable(&domain.Article{}); err != nil {
		return err
	}
	// layer one
	if err := m.DB.
		WithContext(ctx).
		Migrator().
		DropTable(&domain.User{}); err != nil {
		return err
	}
	if err := m.DB.
		WithContext(ctx).
		Migrator().
		DropTable(&domain.ArticleLanguage{}); err != nil {
		return err
	}

	// layer one
	if err := m.DB.
		WithContext(ctx).
		Set("gorm:table_options", "ENGINE=InnoDB").
		Migrator().
		CreateTable(&domain.User{}); err != nil {
		return err
	}
	if err := m.DB.
		WithContext(ctx).
		Set("gorm:table_options", "ENGINE=InnoDB").
		Migrator().
		CreateTable(&domain.ArticleLanguage{}); err != nil {
		return err
	}

	// layer two
	if err := m.DB.
		WithContext(ctx).
		Set("gorm:table_options", "ENGINE=InnoDB").
		Migrator().
		CreateTable(&domain.Article{}); err != nil {
		return err
	}
	return nil
}

func (m *MigrationRepoMysql) Seed(ctx context.Context) error {
	var lang []domain.ArticleLanguage
	lang = append(lang, domain.ArticleLanguage{
		Language: "english",
	})
	lang = append(lang, domain.ArticleLanguage{
		Language: "indonesia",
	})
	for _, v := range lang {
		if err := m.DB.WithContext(ctx).Create(&v).Error; err != nil {
			return err
		}
	}

	var articles []domain.Article
	articles = append(articles, domain.Article{
		Title:             "What is Cancer?",
		Body:              "Cancer is a disease",
		Thumbnail:         "image",
		ArticleLanguageID: 1,
	})
	articles = append(articles, domain.Article{
		Title:             "Apa itu Kanker?",
		Body:              "Kancker adalah penyakit",
		Thumbnail:         "image",
		ArticleLanguageID: 2,
	})
	for _, v := range articles {
		if err := m.DB.WithContext(ctx).Create(&v).Error; err != nil {
			return err
		}
	}

	return nil
}

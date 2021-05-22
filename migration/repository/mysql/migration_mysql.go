package mysql

import (
	"context"
	"github.com/huf0813/scade_backend_api/domain"
	"github.com/huf0813/scade_backend_api/utils/security"
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
	if err := m.DB.
		WithContext(ctx).
		Migrator().
		DropTable(&domain.Diagnose{}); err != nil {
		return err
	}
	if err := m.DB.
		WithContext(ctx).
		Migrator().
		DropTable(&domain.Subscription{}); err != nil {
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
	if err := m.DB.
		WithContext(ctx).
		Migrator().
		DropTable(&domain.Hospital{}); err != nil {
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
	if err := m.DB.
		WithContext(ctx).
		Set("gorm:table_options", "ENGINE=InnoDB").
		Migrator().
		CreateTable(&domain.Hospital{}); err != nil {
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
	if err := m.DB.
		WithContext(ctx).
		Set("gorm:table_options", "ENGINE=InnoDB").
		Migrator().
		CreateTable(&domain.Diagnose{}); err != nil {
		return err
	}
	if err := m.DB.
		WithContext(ctx).
		Set("gorm:table_options", "ENGINE=InnoDB").
		Migrator().
		CreateTable(&domain.Subscription{}); err != nil {
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
		if err := m.DB.WithContext(ctx).Create(&v).
			Error; err != nil {
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
		if err := m.DB.WithContext(ctx).Create(&v).
			Error; err != nil {
			return err
		}
	}

	pass, err := security.NewHashingValue("1234567890")
	if err != nil {
		return err
	}
	user := domain.User{
		Name:     "Harun Ulum Fajar",
		Address:  "Malang",
		Email:    "harun@gmail.com",
		Phone:    "081308130813",
		Password: pass,
	}
	if err := m.DB.WithContext(ctx).Create(&user).
		Error; err != nil {
		return err
	}

	hospital := domain.Hospital{
		Name:     "RS Bagus",
		Address:  "JL. Sigura",
		Phone:    "0813",
		City:     "Malang",
		Province: "Jawa Timur",
	}
	if err := m.DB.WithContext(ctx).Create(&hospital).
		Error; err != nil {
		return err
	}

	return nil
}

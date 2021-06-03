package mysql

import (
	"context"
	"github.com/huf0813/scade_backend_api/domain"
	"github.com/huf0813/scade_backend_api/upstream/covid_hospital"
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
	// layer three
	if err := m.DB.
		WithContext(ctx).
		Migrator().
		DropTable(&domain.Invoice{}); err != nil {
		return err
	}
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

	// layer three
	if err := m.DB.
		WithContext(ctx).
		Set("gorm:table_options", "ENGINE=InnoDB").
		Migrator().
		CreateTable(&domain.Invoice{}); err != nil {
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
		Thumbnail:         "default.jpg",
		ArticleLanguageID: 1,
	})
	articles = append(articles, domain.Article{
		Title:             "Let's get in closer with skin cancer",
		Body:              "skin cancer has similar criteria like this",
		Thumbnail:         "default.jpg",
		ArticleLanguageID: 1,
	})
	articles = append(articles, domain.Article{
		Title:             "Skin Cancer Growth in Indonesia",
		Body:              "Skin Cancer Growth just rising by 5 - 7 percent in a year",
		Thumbnail:         "default.jpg",
		ArticleLanguageID: 1,
	})
	articles = append(articles, domain.Article{
		Title:             "How to do first aid in skin cancer",
		Body:              "You can do the first aid in skin cancer",
		Thumbnail:         "default.jpg",
		ArticleLanguageID: 1,
	})
	articles = append(articles, domain.Article{
		Title:             "How to prevent skin cancer growth",
		Body:              "Prevent skin cancer growth can be done by doing these",
		Thumbnail:         "default.jpg",
		ArticleLanguageID: 1,
	})
	articles = append(articles, domain.Article{
		Title:             "Apa itu Kanker?",
		Body:              "Kancker adalah penyakit",
		Thumbnail:         "default.jpg",
		ArticleLanguageID: 2,
	})
	articles = append(articles, domain.Article{
		Title:             "Kenali Dini Kanker Kulit",
		Body:              "Kanker kulit memiliki ciri-ciri sebagai berikut",
		Thumbnail:         "default.jpg",
		ArticleLanguageID: 2,
	})
	articles = append(articles, domain.Article{
		Title:             "Perkembangan Kanker Kulit di Indonesia",
		Body:              "Kanker kulit di Indonesia memiliki rasio 5 - 7 persen setiap tahun",
		Thumbnail:         "default.jpg",
		ArticleLanguageID: 2,
	})
	articles = append(articles, domain.Article{
		Title:             "Pertolongan pertama pada Kanker Kulit",
		Body:              "Pertolongan pertama pada Kanker Kulit dilakukan sebagai berikut",
		Thumbnail:         "default.jpg",
		ArticleLanguageID: 2,
	})
	articles = append(articles, domain.Article{
		Title:             "Cara mencegah Kanker Kulit",
		Body:              "Mencegah Kanker Kulit dapat dilakukan sebagai berikut",
		Thumbnail:         "default.jpg",
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

	result, err := covid_hospital.Load()
	if err != nil {
		return err
	}
	if err := m.DB.WithContext(ctx).Create(result).
		Error; err != nil {
		return err
	}

	return nil
}

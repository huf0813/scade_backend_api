package domain

import (
	"context"
	"gorm.io/gorm"
)

type ArticleLanguage struct {
	gorm.Model
	Language string    `gorm:"not_null;unique" json:"language"`
	Articles []Article `gorm:"foreignKey:ArticleLanguageID" json:"articles"`
}

type ArticleLanguageRepository interface {
	GetArticleLanguages(ctx context.Context) ([]ArticleLanguage, error)
}

type ArticleLanguageUseCase interface {
	GetArticleLanguages(ctx context.Context) ([]ArticleLanguage, error)
}

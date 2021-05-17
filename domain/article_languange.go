package domain

import (
	"context"
	"gorm.io/gorm"
)

type ArticleLanguage struct {
	gorm.Model
	ID       uint      `gorm:"primaryKey;autoIncrement;not_null" json:"id"`
	Language string    `gorm:"not_null" json:"language"`
	Articles []Article `gorm:"foreignKey:ArticleID" json:"articles"`
}

type ArticleLanguageRepository interface {
	GetArticleLanguages(ctx context.Context) ([]ArticleLanguage, error)
}

type ArticleLanguageUseCase interface {
	GetArticleLanguages(ctx context.Context) ([]ArticleLanguage, error)
}

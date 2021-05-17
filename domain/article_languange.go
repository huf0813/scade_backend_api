package domain

import (
	"context"
	"gorm.io/gorm"
)

type ArticleLanguage struct {
	gorm.Model
	ID       uint      `gorm:"primaryKey;autoIncrement;not_null"`
	Language string    `gorm:"not_null"`
	Articles []Article `gorm:"foreignKey:ArticleID"`
}

type ArticleLanguageResponse struct {
	ID       uint   `json:"id"`
	Language string `json:"language"`
}

type ArticleLanguageRepository interface {
	GetArticleLanguages(ctx context.Context) ([]ArticleLanguageResponse, error)
}

type ArticleLanguageUseCase interface {
	GetArticleLanguages(ctx context.Context) ([]ArticleLanguageResponse, error)
}

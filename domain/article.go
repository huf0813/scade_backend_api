package domain

import (
	"context"
	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	ID    uint   `gorm:"primaryKey;autoIncrement;not_null"`
	Title string `gorm:"not_null"`
	Body  string `gorm:"not_null"`
}

type ArticleResponse struct {
	ID    uint   `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

type ArticleRepository interface {
	GetArticles(ctx context.Context) ([]ArticleResponse, error)
}

type ArticleUseCase interface {
	GetArticles(ctx context.Context) ([]ArticleResponse, error)
}

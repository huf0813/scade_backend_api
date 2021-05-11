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

type ArticleRepository interface {
	GetArticle(ctx context.Context) ([]Article, error)
}

type ArticleUseCase interface {
	GetArticle(ctx context.Context) ([]Article, error)
}

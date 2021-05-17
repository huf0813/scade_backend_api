package domain

import (
	"context"
	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	ID        uint   `gorm:"primaryKey;autoIncrement;not_null" json:"id"`
	Title     string `gorm:"not_null" json:"title"`
	Body      string `gorm:"not_null" json:"body"`
	Thumbnail string `gorm:"not_null" json:"thumbnail"`
	ArticleID uint   `json:"article_id"`
}

type ArticleRepository interface {
	GetArticles(ctx context.Context) ([]Article, error)
}

type ArticleUseCase interface {
	GetArticles(ctx context.Context) ([]Article, error)
}

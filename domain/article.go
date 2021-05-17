package domain

import (
	"context"
	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	//ID        uint   `gorm:"primaryKey;autoIncrement;not_null" json:"id"`
	Title             string `gorm:"not_null" json:"title"`
	Body              string `gorm:"not_null" json:"body"`
	Thumbnail         string `gorm:"not_null" json:"thumbnail"`
	ArticleLanguageID uint   `json:"article_language_id"`
}

type ArticleRepository interface {
	GetArticles(ctx context.Context) ([]Article, error)
	GetArticlesBasedOnLanguage(ctx context.Context, language string) ([]Article, error)
}

type ArticleUseCase interface {
	GetArticles(ctx context.Context) ([]Article, error)
	GetArticlesBasedOnLanguage(ctx context.Context, language string) ([]Article, error)
}

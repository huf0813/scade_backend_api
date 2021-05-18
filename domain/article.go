package domain

import (
	"context"
	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
	Title             string `gorm:"not_null" json:"title"`
	Body              string `gorm:"not_null" json:"body"`
	Thumbnail         string `gorm:"not_null" json:"thumbnail"`
	ArticleLanguageID uint   `json:"article_language_id"`
}

type ArticleRequest struct {
	Title             string `json:"title" validate:"required"`
	Body              string `json:"body" validate:"required"`
	Thumbnail         string `json:"thumbnail" validate:"required"`
	ArticleLanguageID uint   `json:"article_language_id" validate:"required"`
}

type ArticleRepository interface {
	GetArticles(ctx context.Context) ([]Article, error)
	GetArticlesBasedOnLanguage(ctx context.Context, language string) ([]Article, error)
	GetArticlesBasedOnLanguageByID(ctx context.Context, language string, articleID int) (Article, error)
	CreateArticle(ctx context.Context, title, body, thumbnail string, articleLanguageID uint) error
}

type ArticleUseCase interface {
	GetArticles(ctx context.Context) ([]Article, error)
	GetArticlesBasedOnLanguage(ctx context.Context, language string) ([]Article, error)
	GetArticlesBasedOnLanguageByID(ctx context.Context, language string, articleID int) (Article, error)
	CreateArticle(ctx context.Context, title, body, thumbnail string, articleLanguageID uint) error
}

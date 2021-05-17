package mysql

import (
	"context"
	"github.com/huf0813/scade_backend_api/domain"
	"gorm.io/gorm"
)

type ArticleRepoMysql struct {
	DB *gorm.DB
}

func NewArticleRepoMysql(db *gorm.DB) domain.ArticleRepository {
	return &ArticleRepoMysql{DB: db}
}

func (a *ArticleRepoMysql) GetArticles(ctx context.Context) ([]domain.Article, error) {
	var articles []domain.Article
	if err := a.DB.WithContext(ctx).Find(&articles).Error; err != nil {
		return nil, err
	}
	return articles, nil
}
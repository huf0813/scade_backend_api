package mysql

import (
	"context"
	"github.com/huf0813/scade_backend_api/domain"
	"gorm.io/gorm"
)

type ArticleLanguageRepoMysql struct {
	DB *gorm.DB
}

func NewArticleLanguageRepoMysql(db *gorm.DB) domain.ArticleLanguageRepository {
	return &ArticleLanguageRepoMysql{DB: db}
}

func (a *ArticleLanguageRepoMysql) GetArticleLanguages(ctx context.Context) ([]domain.ArticleLanguage, error) {
	var lang []domain.ArticleLanguage
	if err := a.DB.WithContext(ctx).Find(&lang).Error; err != nil {
		return nil, err
	}
	return lang, nil
}

package usecase

import (
	"context"
	"github.com/huf0813/scade_backend_api/domain"
	"time"
)

type ArticleLanguageUseCase struct {
	articleLanguageRepoMysql domain.ArticleLanguageRepository
	timeOut                  time.Duration
}

func NewArticleLanguageUseCase(a domain.ArticleLanguageRepository, t time.Duration) domain.ArticleLanguageUseCase {
	return &ArticleLanguageUseCase{
		articleLanguageRepoMysql: a,
		timeOut:                  t,
	}
}

func (a ArticleLanguageUseCase) GetArticleLanguages(ctx context.Context) ([]domain.ArticleLanguage, error) {
	ctx, cancel := context.WithTimeout(ctx, a.timeOut)
	defer cancel()

	res, err := a.articleLanguageRepoMysql.GetArticleLanguages(ctx)
	if err != nil {
		return nil, err
	}

	return res, nil
}

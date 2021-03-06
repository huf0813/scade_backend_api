package usecase

import (
	"context"
	"github.com/huf0813/scade_backend_api/domain"
	"time"
)

type ArticleUseCase struct {
	articleRepoMysql domain.ArticleRepository
	timeOut          time.Duration
}

func NewArticleUseCase(ar domain.ArticleRepository, timeOut time.Duration) domain.ArticleUseCase {
	return &ArticleUseCase{
		articleRepoMysql: ar,
		timeOut:          timeOut,
	}
}

func (a *ArticleUseCase) GetArticles(ctx context.Context) ([]domain.Article, error) {
	ctx, cancel := context.WithTimeout(ctx, a.timeOut)
	defer cancel()

	res, err := a.articleRepoMysql.GetArticles(ctx)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (a *ArticleUseCase) GetArticlesBasedOnLanguage(ctx context.Context, language string) ([]domain.Article, error) {
	ctx, cancel := context.WithTimeout(ctx, a.timeOut)
	defer cancel()

	res, err := a.articleRepoMysql.GetArticlesBasedOnLanguage(ctx, language)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (a *ArticleUseCase) GetArticlesBasedOnLanguageByID(ctx context.Context, language string, articleID int) (domain.Article, error) {
	ctx, cancel := context.WithTimeout(ctx, a.timeOut)
	defer cancel()

	res, err := a.articleRepoMysql.GetArticlesBasedOnLanguageByID(ctx, language, articleID)
	if err != nil {
		return domain.Article{}, err
	}

	return res, nil
}

func (a *ArticleUseCase) CreateArticle(ctx context.Context, title, body, thumbnail string, articleLanguageID uint) error {
	ctx, cancel := context.WithTimeout(ctx, a.timeOut)
	defer cancel()

	if err := a.articleRepoMysql.CreateArticle(ctx,
		title,
		body,
		thumbnail,
		articleLanguageID); err != nil {
		return err
	}

	return nil
}

package mysql

import (
	"context"
	"errors"
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

func (a *ArticleRepoMysql) GetArticlesBasedOnLanguage(ctx context.Context, language string) ([]domain.Article, error) {
	var articles []domain.Article
	if err := a.DB.WithContext(ctx).
		Joins("JOIN article_languages ON articles.article_language_id = article_languages.id").
		Where("article_languages.language = ?", language).
		Find(&articles).Error; err != nil {
		return nil, err
	}
	return articles, nil
}

func (a *ArticleRepoMysql) GetArticlesBasedOnLanguageByID(ctx context.Context, language string, articleID int) (domain.Article, error) {
	var article domain.Article
	if err := a.DB.WithContext(ctx).
		Joins("JOIN article_languages ON articles.article_language_id = article_languages.id").
		Where("article_languages.language = ?", language).
		First(&article, articleID).Error; err != nil {
		return domain.Article{}, err
	}
	return article, nil
}

func (a *ArticleRepoMysql) CreateArticle(ctx context.Context, title, body, thumbnail string, articleLanguageID uint) error {
	article := domain.Article{
		Title:             title,
		Body:              body,
		Thumbnail:         thumbnail,
		ArticleLanguageID: articleLanguageID,
	}
	result := a.DB.
		WithContext(ctx).
		Create(&article)
	if err := result.Error; err != nil {
		return err
	}
	if rows := result.RowsAffected; rows <= 0 {
		return errors.New("failed to insert data, empty feedback")
	}
	return nil
}

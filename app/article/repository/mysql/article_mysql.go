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

func (a *ArticleRepoMysql) GetArticles(ctx context.Context) ([]domain.ArticleResponse, error) {
	rows, err := a.DB.
		WithContext(ctx).Table("articles").
		Order("articles.created_at desc").
		Select("articles.id", "articles.title", "articles.body").
		Rows()
	if err != nil {
		return nil, err
	}
	var result []domain.ArticleResponse
	for rows.Next() {
		var row domain.ArticleResponse
		if err := rows.Scan(&row.ID, &row.Title, &row.Body); err != nil {
			return nil, err
		}
		result = append(result, row)
	}
	return result, nil
}

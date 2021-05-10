package domain

type Article struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

type ArticleRepository interface {
	GetArticle() ([]Article, error)
}

type ArticleUseCase interface {
	GetArticle() ([]Article, error)
}

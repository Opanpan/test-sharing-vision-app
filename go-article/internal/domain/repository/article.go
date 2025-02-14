package domain

import "github.com/Opanpan/go-article-service/internal/domain/request"

type Article struct {
	ID        int64  `json:"id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	Category  string `json:"category"`
	Status    string `json:"status"`
	CreatedAt string `json:"created_date"`
}

// ArticleRepository defines the contract for article-related database operations.
type ArticleRepository interface {
	Create(article *request.CreateArticleRequest) (int64, error)
	GetByID(id int64) (*Article, error)
	GetAllArticles(limit int, offset int, status string) ([]Article, error)
	Update(article *Article) error
	Delete(id int64) error
}

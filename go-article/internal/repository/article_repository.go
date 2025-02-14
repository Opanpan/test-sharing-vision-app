package repository

import (
	"database/sql"

	domain "github.com/Opanpan/go-article-service/internal/domain/repository"
	"github.com/Opanpan/go-article-service/internal/domain/request"
)

type ArticleRepositoryImpl struct {
	db *sql.DB
}

func NewArticleRepository(db *sql.DB) domain.ArticleRepository {
	return &ArticleRepositoryImpl{db: db}
}

func (r *ArticleRepositoryImpl) Create(article *request.CreateArticleRequest) (int64, error) {
	query := "INSERT INTO posts (title, content, category, status) VALUES (?, ?, ?, ?)"
	result, err := r.db.Exec(query, article.Title, article.Content, article.Category, article.Status)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

func (r *ArticleRepositoryImpl) GetByID(id int64) (*domain.Article, error) {
	query := "SELECT id, title, content, category, status, created_date FROM posts WHERE id = ?"
	row := r.db.QueryRow(query, id)

	var article domain.Article
	err := row.Scan(&article.ID, &article.Title, &article.Content, &article.Category, &article.Status, &article.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &article, nil
}

func (r *ArticleRepositoryImpl) GetAllArticles(limit int, offset int, status string) ([]domain.Article, error) {
	query := "SELECT id, title, content, category, status, created_date FROM posts WHERE status = ? LIMIT ? OFFSET ?"
	rows, err := r.db.Query(query, status, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var articles []domain.Article
	for rows.Next() {
		var article domain.Article
		err := rows.Scan(&article.ID, &article.Title, &article.Content, &article.Category, &article.Status, &article.CreatedAt)
		if err != nil {
			return nil, err
		}
		articles = append(articles, article)
	}

	return articles, nil
}

func (r *ArticleRepositoryImpl) Update(article *domain.Article) error {
	query := "UPDATE posts SET title = ?, content = ?, category = ?, status = ? WHERE id = ?"
	_, err := r.db.Exec(query, article.Title, article.Content, article.Category, article.Status, article.ID)
	if err != nil {
		return err
	}
	return nil
}

func (r *ArticleRepositoryImpl) Delete(id int64) error {
	query := "DELETE FROM posts WHERE id = ?"
	_, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}

package service

import (
	"fmt"

	domain "github.com/Opanpan/go-article-service/internal/domain/repository"
	"github.com/Opanpan/go-article-service/internal/domain/request"
	"github.com/Opanpan/go-article-service/internal/domain/response"
)

type ArticleService struct {
	repo domain.ArticleRepository
}

func NewArticleService(repo domain.ArticleRepository) *ArticleService {
	return &ArticleService{repo: repo}
}

func (s *ArticleService) CreateArticle(article *request.CreateArticleRequest) (int64, error) {
	return s.repo.Create(article)
}

func (s *ArticleService) GetArticleByID(id int) (*response.ArticleResponse, error) {
	// convert int to int64
	idArticle := int64(id)
	res, err := s.repo.GetByID(idArticle)

	if err != nil {
		fmt.Printf("%# v\n", err.Error())
		if err.Error() == "sql: no rows in result set" {
			return nil, nil
		}
		return nil, err
	}

	newRes := &response.ArticleResponse{
		Title:    res.Title,
		Content:  res.Content,
		Category: res.Category,
		Status:   res.Status,
	}
	return newRes, nil

}

func (s *ArticleService) GetAllArticles(limit int, offset int, status string) ([]response.ArticleResponse, error) {
	res, err := s.repo.GetAllArticles(limit, offset, status)

	if err != nil {
		return nil, err
	}

	var newRes []response.ArticleResponse
	for _, v := range res {
		newRes = append(newRes, response.ArticleResponse{
			ID:	      v.ID,
			Title:    v.Title,
			Content:  v.Content,
			Category: v.Category,
			Status:   v.Status,
		})
	}
	return newRes, nil
}

func (s *ArticleService) UpdateArticle(id int, article *request.UpdateArticleRequest) error {
	// convert int to int64
	idArticle := int64(id)
	updateArticle := &domain.Article{
		ID:       idArticle,
		Title:    article.Title,
		Content:  article.Content,
		Category: article.Category,
		Status:   article.Status,
	}
	return s.repo.Update(updateArticle)
}

func (s *ArticleService) DeleteArticle(id int) error {
	return s.repo.Delete(int64(id))
}

package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Opanpan/go-article-service/internal/domain/request"
	"github.com/Opanpan/go-article-service/internal/domain/response"
	"github.com/Opanpan/go-article-service/internal/helper"
	"github.com/Opanpan/go-article-service/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type ArticleController struct {
	service *service.ArticleService
}

func NewArticleController(service *service.ArticleService) *ArticleController {
	return &ArticleController{service: service}
}

// CreateArticle godoc
// @Summary Create Article
// @Description Create Article
// @Tags Article
// @Accept  json
// @Produce  json
// @Param user body request.CreateArticleRequest true "Article"
// @Success 201 {object} response.SuccessResponse
// @Failure 400 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /article [post]
func (c *ArticleController) CreateArticle(ctx *gin.Context) {
	var req request.CreateArticleRequest

	// Bind request body to struct
	if err := ctx.ShouldBindJSON(&req); err != nil {
		message, data := helper.GlobalCheckingErrorBindJson(err.Error())
		logrus.Println(fmt.Sprintf("Error: %s", message)) //NOSONAR
		helper.ReturnJSONError(ctx, http.StatusBadRequest, message, nil, data)
		return
	}

	res := helper.DoValidation(req)

	if len(res) > 0 {
		logrus.Println("Error: Validation error")
		helper.ReturnJSONError(ctx, http.StatusBadRequest, "Validation error", nil, res)
		return
	}

	_, err := c.service.CreateArticle(&req)

	if err != nil {
		logrus.Println(fmt.Sprintf("Error: %s", err.Error()))
		helper.ReturnJSONError(ctx, http.StatusInternalServerError, err.Error(), nil, nil)
		return
	}

	ctx.JSON(http.StatusCreated, response.SuccessResponse{
		Code:    http.StatusCreated,
		Message: http.StatusText(http.StatusCreated),
	})
}

// GetArticleById godoc
// @Summary Get Article By Id
// @Description Get Article By Id
// @Tags Article
// @Accept  json
// @Produce  json
// @Param id path int true "Article ID"
// @Success 200 {object} response.ArticleResponse
// @Failure 400 {object} response.ErrorResponse
// @Failure 404 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /article/{id} [get]
func (c *ArticleController) GetArticleById(ctx *gin.Context) {
	id := ctx.Param("id")
	// Convert string to int
	idInt, err := strconv.Atoi(id)
	if err != nil {
		logrus.Println(fmt.Sprintf("Error: %s", err.Error()))
		helper.ReturnJSONError(ctx, http.StatusBadRequest, err.Error(), nil, nil)
	}

	article, err := c.service.GetArticleByID(idInt)

	if err != nil {
		logrus.Println(fmt.Sprintf("Error: %s", err.Error()))
		helper.ReturnJSONError(ctx, http.StatusInternalServerError, err.Error(), nil, nil)
		return
	}

	if article == nil {
		logrus.Println(fmt.Sprintf("Error: Article not found"))
		helper.ReturnJSONError(ctx, http.StatusNotFound, "Article not found", nil, nil)
		return
	}

	ctx.JSON(http.StatusOK, article)
}

// GetAllArticles godoc
// @Summary Get All Articles
// @Description Get All Articles
// @Tags Article
// @Accept  json
// @Produce  json
// @Success 200 {object} response.ArticlesResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /articles/:limit/:offset [get]
func (c *ArticleController) GetAllArticles(ctx *gin.Context) {
	// get limit
	limit, err := strconv.Atoi(ctx.Param("limit"))
	if err != nil {
		logrus.Println(fmt.Sprintf("Error: %s", err.Error()))
		helper.ReturnJSONError(ctx, http.StatusBadRequest, err.Error(), nil, nil)
	}

	offset, err := strconv.Atoi(ctx.Param("offset"))
	if err != nil {
		logrus.Println(fmt.Sprintf("Error: %s", err.Error()))
		helper.ReturnJSONError(ctx, http.StatusBadRequest, err.Error(), nil, nil)
	}

	articles, err := c.service.GetAllArticles(limit, offset, ctx.Param("status"))

	if err != nil {
		logrus.Println(fmt.Sprintf("Error: %s", err.Error()))
		helper.ReturnJSONError(ctx, http.StatusInternalServerError, err.Error(), nil, nil)
		return
	}

	ctx.JSON(http.StatusOK,articles)
}

// UpdateArticle godoc
// @Summary Update Article
// @Description Update Article
// @Tags Article
// @Accept  json
// @Produce  json
// @Param id path int true "Article ID"
// @Param user body request.UpdateArticleRequest true "Article"
// @Success 200 {object} response.SuccessResponse
// @Failure 400 {object} response.ErrorResponse
// @Failure 404 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /article/{id} [put]
func (c *ArticleController) UpdateArticle(ctx *gin.Context) {
	id := ctx.Param("id")
	// Convert string to int
	idInt, err := strconv.Atoi(id)
	if err != nil {
		logrus.Println(fmt.Sprintf("Error: %s", err.Error()))
		helper.ReturnJSONError(ctx, http.StatusBadRequest, err.Error(), nil, nil)
	}

	var req request.UpdateArticleRequest

	// Bind request body to struct
	if err := ctx.ShouldBindJSON(&req); err != nil {
		message, data := helper.GlobalCheckingErrorBindJson(err.Error())
		logrus.Println(fmt.Sprintf("Error: %s", message))
		helper.ReturnJSONError(ctx, http.StatusBadRequest, message, nil, data)
		return
	}

	res := helper.DoValidation(req)

	if len(res) > 0 {
		logrus.Println("Error: Validation error")
		helper.ReturnJSONError(ctx, http.StatusBadRequest, "Validation error", nil, res)
		return
	}

	err = c.service.UpdateArticle(idInt, &req)

	if err != nil {
		logrus.Println(fmt.Sprintf("Error: %s", err.Error()))
		helper.ReturnJSONError(ctx, http.StatusInternalServerError, err.Error(), nil, nil)
		return
	}

	ctx.JSON(http.StatusOK, response.SuccessResponse{
		Code:    http.StatusOK,
		Message: http.StatusText(http.StatusOK),
	})
}

// DeleteArticle godoc
// @Summary Delete Article
// @Description Delete Article
// @Tags Article
// @Accept  json
// @Produce  json
// @Param id path int true "Article ID"
// @Success 200 {object} response.SuccessResponse
// @Failure 400 {object} response.ErrorResponse
// @Failure 404 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /article/{id} [delete]
func (c *ArticleController) DeleteArticle(ctx *gin.Context) {
	id := ctx.Param("id")
	// Convert string to int
	idInt, err := strconv.Atoi(id)
	if err != nil {
		logrus.Println(fmt.Sprintf("Error: %s", err.Error()))
		helper.ReturnJSONError(ctx, http.StatusBadRequest, err.Error(), nil, nil)
	}

	err = c.service.DeleteArticle(idInt)

	if err != nil {
		logrus.Println(fmt.Sprintf("Error: %s", err.Error()))
		helper.ReturnJSONError(ctx, http.StatusInternalServerError, err.Error(), nil, nil)
		return
	}

	ctx.JSON(http.StatusOK, response.SuccessResponse{
		Code:    http.StatusOK,
		Message: http.StatusText(http.StatusOK),
	})
}

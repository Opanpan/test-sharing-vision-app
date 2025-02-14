package router

import (
	"github.com/Opanpan/go-article-service/internal/controller"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/Opanpan/go-article-service/docs" // Import Swagger docs package
)

type router struct {
	router  *gin.Engine
	article *controller.ArticleController
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*") // Allow all origins
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		c.Next()
	}
}

func NewRouter(article *controller.ArticleController) *router {
	return &router{
		router:  gin.Default(),
		article: article,
	}
}

func (r *router) SetupRouter(port string) {

	r.router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	r.router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "OK",
		})
	})

	r.router.POST("/article", r.article.CreateArticle)

	r.router.PUT("/article/:id", r.article.UpdateArticle) //NOSONAR

	r.router.DELETE("/article/:id", r.article.DeleteArticle) //NOSONAR

	r.router.GET("/articles/:status/:limit/:offset", r.article.GetAllArticles)
	r.router.GET("/article/:id", r.article.GetArticleById) //NOSONAR

	r.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	r.router.Run(port)
}

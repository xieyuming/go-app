package routers

import (
	v1 "github.com/frankie/go-app/internal/models/api/v1"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	apiv1 := r.Group("/api/v1")
	{
		apiv1.GET("/tags", v1.NewTag().Get)
		apiv1.POST("/tags", v1.NewTag().Create)
		apiv1.PUT("/tags/:id", v1.NewTag().Update)
		apiv1.DELETE("/tags/:id", v1.NewTag().Delete)
		apiv1.PATCH("/tags/:id/state", v1.NewTag().Update)

		apiv1.GET("/articles", v1.NewArticle().Get)
		apiv1.GET("/articles/:id", v1.NewArticle().List)
		apiv1.POST("/articles", v1.NewArticle().Create)
		apiv1.PUT("/articles/:id", v1.NewArticle().Update)
		apiv1.DELETE("/articles/:id", v1.NewArticle().Delete)
		apiv1.PATCH("/articles/:id/state", v1.NewArticle().Update)
	}

	return r
}

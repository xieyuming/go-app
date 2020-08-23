package routers

import (
	v1 "github.com/frankie/go-app/internal/models/api/v1"
	"github.com/frankie/go-app/pkg/setting"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	gin.SetMode(setting.RunMode)

	apiv1 := r.Group("/api/v1")
	{
		//获取标签列表
		apiv1.GET("/tags", v1.GetTags)
		//新建标签
		apiv1.POST("/tags", v1.AddTag)
		//更新指定标签
		apiv1.PUT("/tags/:id", v1.EditTag)
		//删除指定标签
		apiv1.DELETE("/tags/:id", v1.DeleteTag)
	}

	return r
}

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	apiv1 := r.Group("/api/v1")
	{
		apiv1.GET("/tags")
		apiv1.POST("/tags")
		apiv1.PUT("/tags/:id")
		apiv1.DELETE("/tags/:id")
		apiv1.PATCH("/tags/:id/state")

		apiv1.GET("/articles")
		apiv1.GET("/articles/:id")
		apiv1.POST("/articles")
		apiv1.PUT("/articles/:id")
		apiv1.DELETE("/articles/:id")
		apiv1.PATCH("/articles/:id/state")
	}

	return r
}

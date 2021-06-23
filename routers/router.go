package routers

import (
	"github.com/gin-gonic/gin"

	"github.com/zhang21/learn_gin/middleware/jwt"
	"github.com/zhang21/learn_gin/pkg/setting"
	"github.com/zhang21/learn_gin/routers/api"
	v1 "github.com/zhang21/learn_gin/routers/api/v1"
)

// InitRouter initialize routers.
func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gin.SetMode(setting.RunMode)

	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "test",
		})
	})

	r.GET("/auth", api.GetAuth)

	apiv1 := r.Group("/api/v1")
	// add auth to api
	apiv1.Use(jwt.JWT())
	{
		// tags
		apiv1.GET("/tags", v1.GetTags)
		apiv1.POST("/tags", v1.AddTag)
		apiv1.PUT("/tags/:id", v1.EditTag)
		apiv1.DELETE("/tags/:id", v1.DeleteTag)
		// articles
		apiv1.GET("/articles", v1.GetArticles)
		apiv1.GET("/articles/:id", v1.GetArticle)
		apiv1.POST("/articles", v1.AddArticle)
		apiv1.PUT("/articles/:id", v1.EditArticle)
		apiv1.DELETE("/articles/:id", v1.DeleteArticle)
	}

	return r
}

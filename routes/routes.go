package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/inciner8r/blog_backend_go/controllers"
)

func Routes(router *gin.Engine) {
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	router.POST("/newBlog", controllers.CreateBlog)
}

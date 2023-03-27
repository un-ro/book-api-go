package routers

import (
	"book-api-go/controllers"
	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	r := gin.Default()

	bookRouter := r.Group("/books")
	{
		bookRouter.GET("/", controllers.GetBooks)
		bookRouter.GET("/:id", controllers.GetBook)
		bookRouter.POST("/", controllers.CreateBook)
		bookRouter.PUT("/:id", controllers.UpdateBook)
		bookRouter.DELETE("/:id", controllers.DeleteBook)
	}

	return r
}

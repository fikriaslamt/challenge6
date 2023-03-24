package routes

import (
	"ch6/controller"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	var r = gin.Default()

	r.GET("/books", controller.GetAllBook)
	r.GET("/books/:id", controller.GetBookById)
	r.POST("/books/add", controller.AddBook)
	r.PUT("/books/update/:id", controller.UpdateBook)
	r.DELETE("/books/delete/:id", controller.DeleteBook)

	return r
}

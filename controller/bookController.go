package controller

import (
	"ch6/model"
	"ch6/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllBook(c *gin.Context) {
	data := service.GetAllBook()

	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusCreated, gin.H{"book": data})
}
func GetBookById(c *gin.Context) {
	idParam := c.Param("id")

	result, err := service.GetBookById(idParam)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusCreated, gin.H{"book": result})
}
func AddBook(c *gin.Context) {
	var book = model.Book{}

	if err := c.ShouldBindJSON(&book); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	err := service.AddBook(book)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusCreated, gin.H{"book": "created"})
}
func UpdateBook(c *gin.Context) {
	idParam := c.Param("id")
	var book = model.Book{}

	if err := c.ShouldBindJSON(&book); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	err := service.UpdateBook(idParam, book)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusCreated, gin.H{"book": "updated"})
}
func DeleteBook(c *gin.Context) {
	idParam := c.Param("id")

	err := service.DeleteBook(idParam)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusCreated, gin.H{"book": "deleted"})
}

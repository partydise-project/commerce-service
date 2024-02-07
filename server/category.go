package server

import (
	"commerce-services/database"

	"github.com/gin-gonic/gin"
)

func CreateCategory(c *gin.Context) {
	var categoryCreateDTO database.CategoriaEvento

	if err := c.BindJSON(&categoryCreateDTO); err != nil {
		c.JSON(400, gin.H{"error in the data": err})
		return
	}

	if len(categoryCreateDTO.Planes) == 0 {
		c.JSON(400, gin.H{"error": "Must create a plan to create a category."})
		return
	}

	err := database.CreateCategory(&categoryCreateDTO)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, categoryCreateDTO)
}

func ReadCategoriesPublished(c *gin.Context) {
	categories, err := database.ReadCategoriesPublished()
	if err != nil {
		c.JSON(500, gin.H{"Error consulting the categories": err})
		return
	}

	c.JSON(200, categories)
}

func ReadCategoriesPublishedAdults(c *gin.Context) {
	categories, err := database.ReadCategoriesPublishedAdults()
	if err != nil {
		c.JSON(500, gin.H{"Error consulting the categories of adults": err})
		return
	}

	c.JSON(200, categories)
}

func ReadCategoriesPublishedChildrens(c *gin.Context) {
	categories, err := database.ReadCategoriesPublishedChildrens()
	if err != nil {
		c.JSON(500, gin.H{"Error consulting the categories of childrens": err})
		return
	}

	c.JSON(200, categories)
}

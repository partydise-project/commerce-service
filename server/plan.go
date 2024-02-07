package server

import (
	"commerce-services/database"

	"github.com/gin-gonic/gin"
)

func CreatePlanEvent(c *gin.Context) {
	var plaEvent database.PlanEvento
	if err := c.BindJSON(&plaEvent); err != nil {
		c.JSON(400, gin.H{"error in the data": err})
		return
	}

	if plaEvent.CategoriaEventoID == 0 {
		c.JSON(400, gin.H{"error": "Must send id of category"})
		return
	}

	err := database.CreatePlanEvent(&plaEvent)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, plaEvent)
}

func ReadPlansPublished(c *gin.Context) {
	plans, err := database.ReadPlansEventPublished(c.Param("id"))
	if err != nil {
		c.JSON(500, gin.H{"error consulting plans": err})
		return
	}

	c.JSON(200, plans)
}

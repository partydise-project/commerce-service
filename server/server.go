package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func WelcomeToCommerceServices(c *gin.Context) {
	c.JSON(http.StatusCreated, "Welcome to the commerce-services.")
}

func InicializeRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/", WelcomeToCommerceServices)

	//Category module.
	r.POST("category", CreateCategory)

	//PlanEvent module.
	r.POST("/plan-event", CreatePlanEvent)
	return r
}

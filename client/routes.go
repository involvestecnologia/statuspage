package client

import (
	"github.com/gin-gonic/gin"
)

//Router creates a new Controller and add all available endpoints to a Gin RouterGroup
func Router(clientService Service, router *gin.RouterGroup) {

	clientController := NewClientController(clientService)

	clientRouter := router.Group("/client")
	{
		clientRouter.POST("", clientController.Create)
		clientRouter.PATCH("/:clientRef", clientController.Update)
		clientRouter.GET("/:clientRef", clientController.Find)
		clientRouter.DELETE("/:clientRef", clientController.Delete)
	}

	clientsRouter := router.Group("/clients")
	{
		clientsRouter.GET("", clientController.List)
	}
}

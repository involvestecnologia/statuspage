package component

import (
	"github.com/gin-gonic/gin"
)

//Router creates a new Controller and add all available endpoints to a Gin RouterGroup
func Router(componentService Service, router *gin.RouterGroup) {

	componentController := NewComponentController(componentService)

	componentRouter := router.Group("/component")
	{
		componentRouter.POST("", componentController.Create)
		componentRouter.PATCH("/:ref", componentController.Update)
		componentRouter.GET("/:ref", componentController.Find)
		componentRouter.DELETE("/:ref", componentController.Delete)
	}

	componentsRouter := router.Group("/components")
	{
		componentsRouter.POST("", componentController.List)
	}
}

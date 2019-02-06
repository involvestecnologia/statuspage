package incident

import (
	"github.com/gin-gonic/gin"
)

//Router creates a new Controller and add all available endpoints to a Gin RouterGroup
func Router(incidentService Service, router *gin.RouterGroup) {

	incidentController := NewIncidentController(incidentService)

	incidentRouter := router.Group("/incident")
	{
		incidentRouter.POST("", incidentController.Create)
		incidentRouter.GET("/:componentRef", incidentController.Find)
	}

	incidentsRouter := router.Group("/incidents")
	{
		incidentsRouter.GET("", incidentController.List)
	}
}

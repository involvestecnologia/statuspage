package prometheus

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/involvestecnologia/statuspage/models"
)

//Controller contains all the available handlers
type Controller struct {
	service Service
}

//NewPrometheusController creates a new Controller
func NewPrometheusController(service Service) *Controller {
	return &Controller{service: service}
}

//Incoming it's the handler function for new Incidents from Prometheus AlertManager
func (prom *Controller) Incoming(c *gin.Context) {
	var incoming models.PrometheusIncomingWebhook
	if err := c.ShouldBindJSON(&incoming); err != nil {
		c.JSON(http.StatusBadRequest, "missing required parameter "+err.Error())
		c.AbortWithError(http.StatusBadGateway, err)
		return
	}
	err := prom.service.ProcessIncomingWebhook(incoming)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusCreated, "")
}

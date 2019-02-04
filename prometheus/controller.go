package prometheus

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/involvestecnologia/statuspage/models"
)

type prometheusController struct {
	service Service
}

func NewPrometheusController(service Service) *prometheusController {
	return &prometheusController{service: service}
}

func (prom *prometheusController) Incoming(c *gin.Context) {
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

package webhook

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/involvestecnologia/statuspage/errors"
	"github.com/involvestecnologia/statuspage/models"
)

//Controller contains all the available handlers
type Controller struct {
	service Service
}

//NewWebhookController creates a new Controller
func NewWebhookController(service Service) *Controller {
	return &Controller{service: service}
}

//Create it's the handler function for Webhook creation endpoints
func (ctrl *Controller) Create(c *gin.Context) {
	var newWebhook models.Webhook
	if err := c.ShouldBindJSON(&newWebhook); err != nil {
		c.JSON(http.StatusBadRequest, "Missing required parameter")
		return
	}
	if exists := ctrl.service.WebhookExists(newWebhook); exists {
		c.JSON(http.StatusConflict, "Webhook "+newWebhook.Name+" already exists")
		return
	}
	err := ctrl.service.Create(newWebhook)
	if err != nil {
		c.JSON(http.StatusInternalServerError, "")
		return
	}
	c.JSON(http.StatusCreated, "")
}

//Find it's the handler function to retrieve Webhooks
func (ctrl *Controller) Find(c *gin.Context) {
	id := c.Param("id")
	webhook, err := ctrl.service.FindWebhook(id)
	if err != nil {
		switch err.(type) {
		case *errors.ErrNotFound:
			c.AbortWithError(http.StatusNotFound, err)
			return
		case *errors.ErrAlreadyExists:
			c.AbortWithError(http.StatusBadRequest, err)
			return
		default:
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
	}
	c.JSON(http.StatusOK, webhook)

}

//Run it's the handler function to execute Webhooks
func (ctrl *Controller) Run(c *gin.Context) {
}

//Update it's the handler function for Webhook update endpoints
func (ctrl *Controller) Update(c *gin.Context) {
	r := strings.Split(c.Request.RequestURI, "/")
	webhookType := r[1]
	id := c.Param("id")

	var newWebhook models.Webhook
	if err := c.ShouldBindJSON(&newWebhook); err != nil {
		c.JSON(http.StatusBadRequest, "Missing required parameter")
		return
	}
	if webhookType != newWebhook.Type {
		c.JSON(http.StatusBadRequest, "You can't change a webhook's type")
		return
	}
	if exist := ctrl.service.WebhookExists(newWebhook); exist {
		c.JSON(http.StatusPreconditionFailed, "Webhook "+newWebhook.Name+" already exists")
		return
	}
	err := ctrl.service.Update(id, newWebhook)
	if err != nil {
		switch err.(type) {
		case *errors.ErrNotFound:
			c.AbortWithError(http.StatusNotFound, err)
			return
		case *errors.ErrAlreadyExists:
			c.AbortWithError(http.StatusBadRequest, err)
			return
		default:
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
	}
	c.JSON(http.StatusOK, "")

}

//Delete it's the handler function for Webhook deletion endpoints
func (ctrl *Controller) Delete(c *gin.Context) {
	id := c.Param("id")

	err := ctrl.service.Delete(id)
	if err != nil {
		switch err.(type) {
		case *errors.ErrNotFound:
			c.AbortWithError(http.StatusNotFound, err)
			return
		case *errors.ErrAlreadyExists:
			c.AbortWithError(http.StatusBadRequest, err)
			return
		default:
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
	}
	c.JSON(http.StatusOK, "")

}

//List it's the handler function for Webhook listing endpoints
func (ctrl *Controller) List(c *gin.Context) {
	r := strings.Split(c.Request.RequestURI, "/")
	webhookType := r[1]

	webhooks, err := ctrl.service.List(webhookType)
	if err != nil {
		switch err.(type) {
		case *errors.ErrNotFound:
			c.AbortWithError(http.StatusNotFound, err)
			return
		default:
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
	}
	c.JSON(http.StatusOK, webhooks)

}

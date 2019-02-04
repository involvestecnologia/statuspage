package webhook

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/involvestecnologia/statuspage/errors"
	"github.com/involvestecnologia/statuspage/models"
)

type WebhookController struct {
	service Service
}

func NewWebhookController(service Service) *WebhookController {
	return &WebhookController{service: service}
}

func (ctrl *WebhookController) Create(c *gin.Context) {
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

func (ctrl *WebhookController) Find(c *gin.Context) {
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

func (ctrl *WebhookController) Run(c *gin.Context) {

}

func (ctrl *WebhookController) Update(c *gin.Context) {
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

func (ctrl *WebhookController) Delete(c *gin.Context) {
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

func (ctrl *WebhookController) List(c *gin.Context) {
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

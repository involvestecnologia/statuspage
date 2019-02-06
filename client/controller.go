package client

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/involvestecnologia/statuspage/errors"
	"github.com/involvestecnologia/statuspage/models"
)

//Controller contains all the available handlers
type Controller struct {
	service Service
}

//NewClientController creates a new Controller
func NewClientController(service Service) *Controller {
	return &Controller{service: service}
}

//Create it's the handler function for Client creation endpoints
func (ctrl *Controller) Create(c *gin.Context) {
	var newClient models.Client
	if err := c.ShouldBindJSON(&newClient); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	ref, err := ctrl.service.CreateClient(newClient)
	if err != nil {
		switch err.(type) {
		case *errors.ErrClientNameAlreadyExists:
			c.AbortWithError(http.StatusBadRequest, err)
			return
		case *errors.ErrClientRefAlreadyExists:
			c.AbortWithError(http.StatusBadRequest, err)
			return
		case *errors.ErrInvalidRef:
			c.AbortWithError(http.StatusBadRequest, err)
			return
		default:
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
	}
	c.JSON(http.StatusCreated, ref)
}

//Update it's the handler function for Client update endpoints
func (ctrl *Controller) Update(c *gin.Context) {
	var client models.Client
	if err := c.ShouldBindJSON(&client); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	clientRef := c.Param("clientRef")
	err := ctrl.service.UpdateClient(clientRef, client)
	if err != nil {
		switch err.(type) {
		case *errors.ErrNotFound:
			c.AbortWithError(http.StatusNotFound, err)
			return
		case *errors.ErrInvalidRef:
			c.AbortWithError(http.StatusBadRequest, err)
			return
		default:
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
	}
	c.JSON(http.StatusOK, "")
}

//Find it's the handler function for filtered Client retrieve endpoints
func (ctrl *Controller) Find(c *gin.Context) {
	queryBy := c.DefaultQuery("search", "ref")
	qValue := c.Param("clientRef")
	client, err := ctrl.service.FindClient(map[string]interface{}{queryBy: qValue})
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
	c.JSON(http.StatusOK, client)
}

//Delete it's the handler function for Client deletion endpoints
func (ctrl *Controller) Delete(c *gin.Context) {
	clientID := c.Param("clientRef")
	err := ctrl.service.RemoveClient(clientID)
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
	c.JSON(http.StatusNoContent, "")
}

//List it's the handler function for Client listing endpoints
func (ctrl *Controller) List(c *gin.Context) {
	clients, err := ctrl.service.ListClients()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, clients)
}

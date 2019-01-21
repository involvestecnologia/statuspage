package client

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/involvestecnologia/statuspage/errors"
	"github.com/involvestecnologia/statuspage/models"
)

type ClientController struct {
	service Service
}

func NewClientController(service Service) *ClientController {
	return &ClientController{service: service}
}

func (ctrl *ClientController) Create(c *gin.Context) {
	var newClient models.Client
	if err := c.ShouldBindJSON(&newClient); err != nil {
		c.JSON(http.StatusBadRequest, "missing required parameter")
		return
	}
	ref, err := ctrl.service.CreateClient(newClient)
	if err != nil {
		if err.Error() == fmt.Sprintf(errors.ErrAlreadyExists, ref) {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}
		c.JSON(http.StatusInternalServerError, "")
		return
	}
	c.JSON(http.StatusCreated, "")
}

func (ctrl *ClientController) Update(c *gin.Context) {
	var client models.Client
	if err := c.ShouldBindJSON(&client); err != nil {
		c.JSON(http.StatusBadRequest, "missing required parameter")
		return
	}
	clientRef := c.Param("clientRef")
	err := ctrl.service.UpdateClient(clientRef, client)
	if err != nil {
		if err.Error() == errors.ErrNotFound {
			c.JSON(http.StatusNotFound, "")
			return
		}
		c.JSON(http.StatusInternalServerError, "")
		return
	}
	c.JSON(http.StatusOK, "")
}

func (ctrl *ClientController) Find(c *gin.Context) {
	queryBy := c.DefaultQuery("search", "ref")
	qValue := c.Param("clientRef")
	fmt.Printf("%s: %s\n", queryBy, qValue)
	client, err := ctrl.service.FindClient(map[string]interface{}{queryBy: qValue})
	if err != nil {
		if err.Error() == errors.ErrNotFound {
			c.JSON(http.StatusNotFound, "")
			return
		}
		c.JSON(http.StatusInternalServerError, "")
		return
	}
	c.JSON(http.StatusOK, client)
}

func (ctrl *ClientController) Delete(c *gin.Context) {
	clientID := c.Param("clientRef")
	err := ctrl.service.RemoveClient(clientID)
	if err != nil {
		if err.Error() == errors.ErrNotFound {
			c.JSON(http.StatusNotFound, "")
			return
		}
		c.JSON(http.StatusInternalServerError, "")
		return
	}
	c.JSON(http.StatusNoContent, "")
}

func (ctrl *ClientController) List(c *gin.Context) {
	clients, err := ctrl.service.ListClients()
	if err != nil {
		c.JSON(http.StatusInternalServerError, "")
		return
	}
	c.JSON(http.StatusOK, clients)
}
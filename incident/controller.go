package incident

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/involvestecnologia/statuspage/errors"
	"github.com/involvestecnologia/statuspage/models"
)

//Controller contains all the available handlers
type Controller struct {
	service Service
}

//NewIncidentController creates a new Controller
func NewIncidentController(service Service) *Controller {
	return &Controller{service: service}
}

//Create it's the handler function for Component creation endpoints
func (ctrl *Controller) Create(c *gin.Context) {
	var newIncident models.Incident
	if err := c.ShouldBindJSON(&newIncident); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	err := ctrl.service.CreateIncidents(newIncident)
	if err != nil {
		switch err.(type) {
		case *errors.ErrNotFound:
			c.AbortWithError(http.StatusNotFound, err)
			return
		case *errors.ErrIncidentStatusIgnored:
			c.AbortWithError(http.StatusPreconditionFailed, err)
			return
		default:
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
	}
	c.JSON(http.StatusCreated, "")
}

//Find it's the handler function for filtered Component retrieve endpoints
func (ctrl *Controller) Find(c *gin.Context) {
	queryBy := c.DefaultQuery("search", "component_ref")
	queryValue := c.Param("componentRef")
	incidents, err := ctrl.service.FindIncidents(map[string]interface{}{queryBy: queryValue})
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
	c.JSON(http.StatusOK, incidents)
}

//List it's the handler function for Component listing endpoints
func (ctrl *Controller) List(c *gin.Context) {
	mQ := c.Query("month")
	yQ := c.Query("year")
	rQ, err := strconv.ParseBool(c.DefaultQuery("unresolved", "false"))
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, &errors.ErrInvalidQuery{Message: errors.ErrInvalidQueryMessage})
		return
	}
	incidents, err := ctrl.service.ListIncidents(yQ, mQ, rQ)
	if err != nil {
		switch err.(type) {
		case *errors.ErrInvalidYear:
			c.AbortWithError(http.StatusBadRequest, err)
			return
		case *errors.ErrInvalidMonth:
			c.AbortWithError(http.StatusBadRequest, err)
			return
		default:
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

	}
	c.JSON(http.StatusOK, incidents)
}

package prometheus

import (
	"github.com/gin-gonic/gin"
	"github.com/involvestecnologia/statuspage/component"
	"github.com/involvestecnologia/statuspage/incident"
	"github.com/involvestecnologia/statuspage/mock"
	"github.com/stretchr/testify/assert"

	"testing"
)

func TestPrometheusRouter_PrometheusRouter(t *testing.T) {
	var result string
	router := gin.New()

	componentDAO := mock.NewMockComponentDAO()
	incidentDAO := mock.NewMockIncidentDAO()
	componentService := component.NewService(componentDAO, mock.NewComponentLogRepositoryMock())
	incidentService := incident.NewService(incidentDAO, componentService)

	Router(incidentService, componentService, router.Group("/v1"))
	r := router.Routes()

	for _, v := range r {
		result = v.Path
	}

	assert.Contains(t, result, "prometheus/incoming")
}

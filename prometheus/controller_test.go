package prometheus

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/involvestecnologia/statuspage/component"
	"github.com/involvestecnologia/statuspage/incident"
	"github.com/involvestecnologia/statuspage/mock"
	"github.com/stretchr/testify/assert"
)

var router = gin.New()

func init() {
	componentDAO := mock.NewMockComponentDAO()
	incidentDAO := mock.NewMockIncidentDAO()
	componentService := component.NewService(componentDAO, mock.NewComponentLogRepositoryMock())
	incidentService := incident.NewService(incidentDAO, componentService)
	Router(incidentService, componentService, router.Group("/v1"))
}

func performRequest(t *testing.T, r http.Handler, method, path string, body []byte) *httptest.ResponseRecorder {
	req, err := http.NewRequest(method, path, bytes.NewBuffer(body))
	if err != nil {
		t.Errorf("Failed to perform request: %s", err.Error())
	}
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func TestController_Incoming(t *testing.T) {
	newPrometheusMock := mock.PrometheusModel()

	ModelWithoutComponent, err := json.Marshal(newPrometheusMock["ModelWithoutComponent"])
	assert.Nil(t, err)
	resp := performRequest(t, router, "POST", "/v1/prometheus/incoming", ModelWithoutComponent)
	assert.Equal(t, http.StatusInternalServerError, resp.Code)

	ModelBlank, err := json.Marshal(newPrometheusMock["ModelBlank"])
	assert.Nil(t, err)
	resp = performRequest(t, router, "POST", "/v1/prometheus/incoming", ModelBlank)
	assert.Equal(t, http.StatusBadRequest, resp.Code)

	ModelComplete, err := json.Marshal(newPrometheusMock["ModelComplete"])
	assert.Nil(t, err)
	resp = performRequest(t, router, "POST", "/v1/prometheus/incoming", ModelComplete)
	assert.Equal(t, http.StatusCreated, resp.Code)
}

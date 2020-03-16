package incident_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/involvestecnologia/statuspage/component"
	"github.com/involvestecnologia/statuspage/incident"
	"github.com/involvestecnologia/statuspage/mock"

	"github.com/involvestecnologia/statuspage/models"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

const routerGroupName = "/test"
const failureRouterGroupName = "/failure"

var router = gin.Default()
var componentService = component.NewService(mock.NewMockComponentDAO(), mock.NewLogRepositoryMock())

func init() {
	incidentService := incident.NewService(mock.NewMockIncidentDAO(), componentService, mock.NewLogRepositoryMock())
	incidentFailureService := incident.NewService(mock.NewMockFailureIncidentDAO(), componentService, mock.NewLogRepositoryMock())
	incident.Router(incidentService, router.Group(routerGroupName))
	incident.Router(incidentFailureService, router.Group(failureRouterGroupName))
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

func TestController_Create(t *testing.T) {
	//Valid: new incident body
	body := []byte(`{"component_ref":"` + mock.ZeroTimeHex + `","status": 3,"description": "test", "occurrence_date": "` + time.Now().Format(time.RFC3339) + `"}`)
	resp := performRequest(t, router, "POST", routerGroupName+"/incident", body)

	assert.Equal(t, http.StatusCreated, resp.Code)

	var data string
	err := json.Unmarshal([]byte(resp.Body.String()), &data)
	assert.Nil(t, err)
	assert.NotNil(t, data)

	//Invalid: incident status is declined
	body = []byte(`{"component_ref":"` + mock.ZeroTimeHex + `","status": 2,"description": "test", "occurrence_date": "` + time.Now().Format(time.RFC3339) + `"}`)
	resp = performRequest(t, router, "POST", routerGroupName+"/incident", body)

	assert.Equal(t, http.StatusPreconditionFailed, resp.Code)

	//Invalid: incident body missing required parameter status
	body = []byte(`{"component_ref":"` + mock.ZeroTimeHex + `","description": "test", "occurrence_date": "` + time.Now().Format(time.RFC3339) + `"}`)
	resp = performRequest(t, router, "POST", routerGroupName+"/incident", body)

	assert.Equal(t, http.StatusBadRequest, resp.Code)

	//Invalid: incident missing required parameter occurrence_date
	body = []byte(`{"component_ref":"` + mock.ZeroTimeHex + `","status":1, "description": "test"}`)
	resp = performRequest(t, router, "POST", routerGroupName+"/incident", body)

	assert.Equal(t, http.StatusBadRequest, resp.Code)

	//Invalid: incident missing required parameter component_ref
	body = []byte(`{"status":1, "description": "test", "occurrence_date": "` + time.Now().Format(time.RFC3339) + `"}`)
	resp = performRequest(t, router, "POST", routerGroupName+"/incident", body)

	assert.Equal(t, http.StatusBadRequest, resp.Code)

	//Invalid: unknow component
	body = []byte(`{"component_ref":"Invalid Component Ref","status": 2,"description": "test", "occurrence_date": "` + time.Now().Format(time.RFC3339) + `"}`)
	resp = performRequest(t, router, "POST", routerGroupName+"/incident", body)

	assert.Equal(t, http.StatusNotFound, resp.Code)

	//Failure
	body = []byte(`{"component_ref":"` + mock.ZeroTimeHex + `","status": 2,"description": "test", "occurrence_date": "` + time.Now().Format(time.RFC3339) + `"}`)
	resp = performRequest(t, router, "POST", failureRouterGroupName+"/incident", body)

	assert.Equal(t, http.StatusInternalServerError, resp.Code)
}
func TestController_Find(t *testing.T) {
	//Valid: incident with exitent ref
	resp := performRequest(t, router, "GET", routerGroupName+"/incident/"+mock.ZeroTimeHex, nil)
	assert.Equal(t, http.StatusOK, resp.Code)

	var data []models.Incident
	err := json.Unmarshal([]byte(resp.Body.String()), &data)
	assert.Nil(t, err)
	assert.NotNil(t, data)

	//Invalid: inexistent component name
	resp = performRequest(t, router, "GET", routerGroupName+"/incident/test2", nil)

	assert.Equal(t, http.StatusNotFound, resp.Code)

	//Failure
	resp = performRequest(t, router, "GET", failureRouterGroupName+"/incident/"+mock.ZeroTimeHex, nil)
	assert.Equal(t, http.StatusInternalServerError, resp.Code)

}

func TestController_List(t *testing.T) {
	// Valid: no parameters
	resp := performRequest(t, router, "GET", routerGroupName+"/incidents", nil)

	assert.Equal(t, http.StatusOK, resp.Code)

	var data []models.Incident
	err := json.Unmarshal([]byte(resp.Body.String()), &data)
	assert.Nil(t, err)
	assert.NotNil(t, data)

	end := time.Now().Format(time.RFC3339)
	assert.Nil(t, err)
	start := time.Now().Add(-12 * time.Hour).Format(time.RFC3339)
	assert.Nil(t, err)
	// Valid: query parameters endDate startDate
	resp = performRequest(t, router, "GET", routerGroupName+"/incidents?startDate="+start+"&endDate="+end, nil)

	assert.Equal(t, http.StatusOK, resp.Code)

	err = json.Unmarshal([]byte(resp.Body.String()), &data)
	assert.Nil(t, err)
	assert.NotNil(t, data)

	// Valid: query parameters only start
	resp = performRequest(t, router, "GET", routerGroupName+"/incidents?startDate="+start, nil)

	assert.Equal(t, http.StatusOK, resp.Code)

	err = json.Unmarshal([]byte(resp.Body.String()), &data)
	assert.Nil(t, err)
	assert.NotNil(t, data)

	// Valid: query parameters only end
	resp = performRequest(t, router, "GET", routerGroupName+"/incidents?endDate="+end, nil)

	assert.Equal(t, http.StatusOK, resp.Code)

	err = json.Unmarshal([]byte(resp.Body.String()), &data)
	assert.Nil(t, err)
	assert.NotNil(t, data)

	// Invalid: query parameters unresolved
	resp = performRequest(t, router, "GET", routerGroupName+"/incidents?unresolved=test", nil)

	assert.Equal(t, http.StatusBadRequest, resp.Code)

	start = time.Now().Add(1 * time.Hour).Format(time.RFC3339)
	// Invalid: query parameters start
	resp = performRequest(t, router, "GET", routerGroupName+"/incidents?startDate="+start, nil)

	assert.Equal(t, http.StatusBadRequest, resp.Code)

	end = time.Now().Add(1 * time.Hour).Format(time.RFC3339)
	// Invalid: query parameters end
	resp = performRequest(t, router, "GET", routerGroupName+"/incidents?endDate="+end, nil)

	assert.Equal(t, http.StatusOK, resp.Code)

	start = time.Now().Add(-2 * time.Hour).Format(time.RFC3339)
	end = time.Now().Add(-3 * time.Hour).Format(time.RFC3339)
	// Invalid: query parameters start and end
	resp = performRequest(t, router, "GET", routerGroupName+"/incidents?startDate="+start+"&endDate="+end, nil)

	assert.Equal(t, http.StatusBadRequest, resp.Code)

	// Failure
	resp = performRequest(t, router, "GET", failureRouterGroupName+"/incidents", nil)

	assert.Equal(t, http.StatusInternalServerError, resp.Code)
}

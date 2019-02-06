package client_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/involvestecnologia/statuspage/client"
	"github.com/involvestecnologia/statuspage/component"
	"github.com/involvestecnologia/statuspage/mock"
	"github.com/involvestecnologia/statuspage/models"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

const routerGroupName = "/test"
const failureRouterGroupName = "/failure"

var router = gin.Default()
var compSvc = component.NewService(mock.NewMockComponentDAO())

func init() {
	clientService := client.NewService(mock.NewMockClientDAO(), compSvc)
	clientFailureService := client.NewService(mock.NewMockFailureClientDAO(), compSvc)
	client.Router(clientService, router.Group(routerGroupName))
	client.Router(clientFailureService, router.Group(failureRouterGroupName))
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
	//Valid: new client body
	body := []byte(`{"name": "test","resource_refs": []}`)
	resp := performRequest(t, router, "POST", routerGroupName+"/client", body)

	assert.Equal(t, http.StatusCreated, resp.Code)

	var data string
	err := json.Unmarshal([]byte(resp.Body.String()), &data)
	assert.Nil(t, err)
	assert.NotNil(t, data)

	//Invalid: client body with exitent name
	body = []byte(`{"name": "First Client","resource_refs": []}`)
	resp = performRequest(t, router, "POST", routerGroupName+"/client", body)

	assert.Equal(t, http.StatusBadRequest, resp.Code)

	//Invalid: client missing required parameter name
	body = []byte(`{"resource_refs": []}`)
	resp = performRequest(t, router, "POST", routerGroupName+"/client", body)

	assert.Equal(t, http.StatusBadRequest, resp.Code)

	//Invalid: client body with exitent ref
	body = []byte(`{"ref":"886e09000000000000000000","name": "test2","resource_refs": []}`)
	resp = performRequest(t, router, "POST", routerGroupName+"/client", body)

	assert.Equal(t, http.StatusBadRequest, resp.Code)

	//Invalid: client body with invalid component ref
	body = []byte(`{"ref":"test3","name": "test3","resource_refs": ["Invalid Ref","Another invalid Ref"]}`)
	resp = performRequest(t, router, "POST", routerGroupName+"/client", body)

	assert.Equal(t, http.StatusBadRequest, resp.Code)

	//Failure DAO
	body = []byte(`{"name": "test","resource_refs": []}`)
	resp = performRequest(t, router, "POST", failureRouterGroupName+"/client", body)

	assert.Equal(t, http.StatusInternalServerError, resp.Code)
}

func TestController_Update(t *testing.T) {
	//Valid: client with exitent ref
	body := []byte(`{"name": "test1","resource_refs": []}`)
	resp := performRequest(t, router, "PATCH", routerGroupName+"/client/886e09000000000000000000", body)

	assert.Equal(t, http.StatusOK, resp.Code)

	//Invalid: inexistent client ref
	body = []byte(`{"name": "test2","resource_refs": []}`)
	resp = performRequest(t, router, "PATCH", routerGroupName+"/client/test2", body)

	assert.Equal(t, http.StatusNotFound, resp.Code)

	//Invalid: missing name parameter
	body = []byte(`{"resource_refs": []}`)
	resp = performRequest(t, router, "PATCH", routerGroupName+"/client/886e09000000000000000000", body)

	assert.Equal(t, http.StatusBadRequest, resp.Code)

	//Valid: client with invalid component ref
	body = []byte(`{"name": "test1","resource_refs": ["Invalid Ref"]}`)
	resp = performRequest(t, router, "PATCH", routerGroupName+"/client/886e09000000000000000000", body)

	assert.Equal(t, http.StatusBadRequest, resp.Code)

	//Failure
	body = []byte(`{"name": "test1","resource_refs": []}`)
	resp = performRequest(t, router, "PATCH", failureRouterGroupName+"/client/886e09000000000000000000", body)
	assert.Equal(t, http.StatusInternalServerError, resp.Code)
}

func TestController_Find(t *testing.T) {
	//Valid: client with exitent ref
	resp := performRequest(t, router, "GET", routerGroupName+"/client/886e09000000000000000000", nil)
	assert.Equal(t, http.StatusOK, resp.Code)

	var data models.Client
	err := json.Unmarshal([]byte(resp.Body.String()), &data)
	assert.Nil(t, err)
	assert.NotNil(t, data)

	//Invalid: inexistent client ref
	resp = performRequest(t, router, "GET", routerGroupName+"/client/test2", nil)

	assert.Equal(t, http.StatusNotFound, resp.Code)

	//Valid: finding client by name
	resp = performRequest(t, router, "GET", routerGroupName+"/client/test?search=name", nil)
	assert.Equal(t, http.StatusOK, resp.Code)

	err = json.Unmarshal([]byte(resp.Body.String()), &data)
	assert.Nil(t, err)
	assert.NotNil(t, data)

	//Failure
	resp = performRequest(t, router, "GET", failureRouterGroupName+"/client/886e09000000000000000000", nil)
	assert.Equal(t, http.StatusInternalServerError, resp.Code)
}

func TestController_Delete(t *testing.T) {
	//Failure
	resp := performRequest(t, router, "DELETE", failureRouterGroupName+"/client/886e09000000000000000000", nil)

	assert.Equal(t, http.StatusInternalServerError, resp.Code)

	resp = performRequest(t, router, "DELETE", routerGroupName+"/client/886e09000000000000000000", nil)

	assert.Equal(t, http.StatusNoContent, resp.Code)

	resp = performRequest(t, router, "DELETE", routerGroupName+"/client/invalidRef", nil)

	assert.Equal(t, http.StatusNotFound, resp.Code)
}

func TestController_List(t *testing.T) {
	resp := performRequest(t, router, "GET", routerGroupName+"/clients", nil)

	assert.Equal(t, http.StatusOK, resp.Code)

	var data []models.Client
	err := json.Unmarshal([]byte(resp.Body.String()), &data)
	assert.Nil(t, err)
	assert.NotNil(t, data)

	resp = performRequest(t, router, "GET", failureRouterGroupName+"/clients", nil)

	assert.Equal(t, http.StatusInternalServerError, resp.Code)
}

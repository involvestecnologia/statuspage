package prometheus

import (
	"testing"

	"github.com/involvestecnologia/statuspage/component"
	"github.com/involvestecnologia/statuspage/incident"
	"github.com/involvestecnologia/statuspage/mock"
	"github.com/stretchr/testify/assert"
)

func NewServicesMock(failure bool, m string) Service {
	componentDAO := mock.NewMockComponentDAO()
	componentFailureDAO := mock.NewMockFailureComponentDAO()
	incidentDAO := mock.NewMockIncidentDAO()
	incidentFailureDAO := mock.NewMockFailureIncidentDAO()
	componentService := component.NewService(componentDAO, mock.NewComponentLogRepositoryMock())
	incidentService := incident.NewService(incidentDAO, componentService)

	if failure {
		if m == "incident" {
			componentService = component.NewService(componentDAO, mock.NewComponentLogRepositoryMock())
			incidentService = incident.NewService(incidentFailureDAO, componentService)
		}
		if m == "component" {
			componentService = component.NewService(componentFailureDAO, mock.NewComponentLogRepositoryMock())
			incidentService = incident.NewService(incidentDAO, componentService)
		}
	}

	return NewPrometheusService(incidentService, componentService)
}

func TestPrometheusService_Service(t *testing.T) {
	assert.Implements(t, (*Service)(nil), NewServicesMock(false, ""))
}

func TestPrometheusService_ProcessIncomingWebhookReturnNil(t *testing.T) {
	newPrometheusMock := mock.PrometheusModel()
	prometheusServiceMock := NewServicesMock(false, "")
	err := prometheusServiceMock.ProcessIncomingWebhook(newPrometheusMock["ModelComplete"])
	assert.Nil(t, err)

	err = prometheusServiceMock.ProcessIncomingWebhook(newPrometheusMock["ModelFiring"])
	assert.Nil(t, err)

	err = prometheusServiceMock.ProcessIncomingWebhook(newPrometheusMock["ModelUnstable"])
	assert.Nil(t, err)

	err = prometheusServiceMock.ProcessIncomingWebhook(newPrometheusMock["ModelResolved"])
	assert.Nil(t, err)
}

func TestPrometheusService_ProcessIncomingWebhookReturnIncidentErr(t *testing.T) {
	newPrometheusMock := mock.PrometheusModel()

	err := NewServicesMock(false, "incident").ProcessIncomingWebhook(newPrometheusMock["ModelCompleteWithBadStatus"])
	assert.Nil(t, err)

	err = NewServicesMock(true, "incident").ProcessIncomingWebhook(newPrometheusMock["ModelResolved"])
	assert.NotNil(t, err)

	err = NewServicesMock(true, "incident").ProcessIncomingWebhook(newPrometheusMock["ModelWithoutName"])
	assert.NotNil(t, err)

	err = NewServicesMock(false, "incident").ProcessIncomingWebhook(newPrometheusMock["ModelWithIncidentWithoutComponentRef"])
	assert.NotNil(t, err)
}

func TestPrometheusService_ProcessIncomingWebhookReturnComponentErr(t *testing.T) {
	newPrometheusMock := mock.PrometheusModel()
	err := NewServicesMock(true, "component").ProcessIncomingWebhook(newPrometheusMock["ModelWithoutRef"])
	assert.NotNil(t, err)

	err = NewServicesMock(false, "").ProcessIncomingWebhook(newPrometheusMock["ModelComponentRefAlreadyExists"])
	assert.NotNil(t, err)

	err = NewServicesMock(false, "").ProcessIncomingWebhook(newPrometheusMock["ModelComponentNameAlreadyExists"])
	assert.NotNil(t, err)
}

package prometheus

import (
	"time"

	"github.com/involvestecnologia/statuspage/component"
	"github.com/involvestecnologia/statuspage/errors"
	"github.com/involvestecnologia/statuspage/incident"
	"github.com/involvestecnologia/statuspage/models"
)

type prometheusService struct {
	incident  incident.Service
	component component.Service
}

//NewPrometheusService creates implementation of the Service interface
func NewPrometheusService(incident incident.Service, component component.Service) Service {
	return &prometheusService{incident: incident, component: component}
}

func (svc *prometheusService) ProcessIncomingWebhook(incoming models.PrometheusIncomingWebhook) error {
	for _, alerts := range incoming.Alerts {
		ref, err := svc.component.CreateComponent(alerts.Component)
		if svc.shouldFail(err) {
			return err
		}
		alerts.Incident.ComponentRef = ref
		if alerts.Incident.Date.IsZero() {
			alerts.Incident.Date = time.Now()
		}
		if err := svc.incident.CreateIncidents(alerts.Incident); err != nil {
			if svc.shouldFail(err) {
				return err
			}
		}
	}
	return nil
}

func (svc *prometheusService) shouldFail(err error) bool {
	switch err.(type) {
	case *errors.ErrComponentNameIsEmpty:
		return true
	case *errors.ErrComponentNameAlreadyExists:
		return false
	case *errors.ErrComponentRefAlreadyExists:
		return false
	case *errors.ErrIncidentStatusIgnored:
		return false
	case nil:
		return false
	default:
		return true
	}
}

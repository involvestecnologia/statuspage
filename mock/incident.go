package mock

import (
	"fmt"
	"time"

	"github.com/involvestecnologia/statuspage/errors"
	"github.com/involvestecnologia/statuspage/incident"
	"github.com/involvestecnologia/statuspage/models"
)

type MockIncidentDAO struct {
	incidents []models.Incident
}

func NewMockIncidentDAO() incident.Repository {
	return &MockIncidentDAO{
		incidents: []models.Incident{
			{
				ComponentRef: ZeroTimeHex,
				Description:  "status ok",
				Status:       1,
				Date:         time.Time{},
			},
			{
				ComponentRef: OneSecTimeHex,
				Description:  "status outage",
				Status:       3,
				Date:         time.Now(),
			},
		},
	}
}

func (m *MockIncidentDAO) Find(query map[string]interface{}) ([]models.Incident, error) {
	var incidents []models.Incident

	for _, i := range m.incidents {
		if i.ComponentRef == query["component_ref"] {
			incidents = append(incidents, i)
		}
	}
	if len(incidents) > 0 {
		return incidents, nil
	}
	return incidents, &errors.ErrNotFound{Message: errors.ErrNotFoundMessage}
}

func (m *MockIncidentDAO) FindOne(query map[string]interface{}) (models.Incident, error) {
	for _, i := range m.incidents {
		if i.ComponentRef == query["component_ref"] {
			return i, nil
		}
	}
	return models.Incident{}, &errors.ErrNotFound{Message: errors.ErrNotFoundMessage}
}

func (m *MockIncidentDAO) Update(incident models.Incident) error {
	for k, i := range m.incidents {
		if i.ComponentRef == incident.ComponentRef {
			m.incidents[k] = incident
			return nil
		}
	}
	return &errors.ErrNotFound{Message: errors.ErrNotFoundMessage}
}

func (m *MockIncidentDAO) List(start time.Time, end time.Time, unresolved bool) ([]models.Incident, error) {
	var inc []models.Incident
	for _, i := range m.incidents {
		if (i.Date.Before(end) && i.Date.Before(end)) || (start.IsZero() && end.IsZero()) {
			if unresolved && !i.Resolved {
				inc = append(inc, i)
				continue
			}
			inc = append(inc, i)
		}
	}
	return inc, nil
}

func (m *MockIncidentDAO) Insert(incident models.Incident) error {
	inc := []models.Incident{incident}
	m.incidents = append(inc, m.incidents...)
	return nil
}

type MockFailureIncidentDAO struct {
}

func NewMockFailureIncidentDAO() incident.Repository {
	return &MockFailureIncidentDAO{}
}

func (m *MockFailureIncidentDAO) Find(query map[string]interface{}) ([]models.Incident, error) {
	return []models.Incident{}, fmt.Errorf("Failure DAO")
}

func (m *MockFailureIncidentDAO) FindOne(query map[string]interface{}) (models.Incident, error) {
	return models.Incident{}, fmt.Errorf("Failure DAO")
}

func (m *MockFailureIncidentDAO) Update(i models.Incident) error {
	return fmt.Errorf("Failure DAO")
}

func (m *MockFailureIncidentDAO) List(start time.Time, end time.Time, unresolved bool) ([]models.Incident, error) {
	return []models.Incident{}, fmt.Errorf("Failure DAO")
}
func (m *MockFailureIncidentDAO) Insert(incident models.Incident) error {
	return fmt.Errorf("Failure DAO")
}

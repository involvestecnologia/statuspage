package mock

import (
	"fmt"
	"time"

	"github.com/involvestecnologia/statuspage/errors"
	"github.com/involvestecnologia/statuspage/incident"
	"github.com/involvestecnologia/statuspage/models"
)

type incidentDAO struct {
	incidents []models.Incident
}

// NewMockIncidentDAO creates a new Incident Repository Mock
func NewMockIncidentDAO() incident.Repository {
	return &incidentDAO{
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

func (m *incidentDAO) Find(query map[string]interface{}) ([]models.Incident, error) {
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
func (m *incidentDAO) FindOne(query map[string]interface{}) (models.Incident, error) {
	for _, i := range m.incidents {
		if i.ComponentRef == query["component_ref"] {
			return i, nil
		}
	}
	return models.Incident{}, &errors.ErrNotFound{Message: errors.ErrNotFoundMessage}
}
func (m *incidentDAO) Update(incident models.Incident) error {
	for k, i := range m.incidents {
		if i.ComponentRef == incident.ComponentRef {
			m.incidents[k] = incident
			return nil
		}
	}
	return &errors.ErrNotFound{Message: errors.ErrNotFoundMessage}
}
func (m *incidentDAO) List(start time.Time, end time.Time, unresolved bool) ([]models.Incident, error) {
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
func (m *incidentDAO) Insert(incident models.Incident) error {
	inc := []models.Incident{incident}
	m.incidents = append(inc, m.incidents...)
	return nil
}

type failureIncidentDAO struct {
}

// NewMockFailureIncidentDAO creates a new Incident Repository Mock that fails in every operation
func NewMockFailureIncidentDAO() incident.Repository {
	return &failureIncidentDAO{}
}

func (m *failureIncidentDAO) Find(query map[string]interface{}) ([]models.Incident, error) {
	return []models.Incident{}, fmt.Errorf("Failure DAO")
}
func (m *failureIncidentDAO) FindOne(query map[string]interface{}) (models.Incident, error) {
	return models.Incident{}, fmt.Errorf("Failure DAO")
}
func (m *failureIncidentDAO) Update(i models.Incident) error {
	return fmt.Errorf("Failure DAO")
}
func (m *failureIncidentDAO) List(start time.Time, end time.Time, unresolved bool) ([]models.Incident, error) {
	return []models.Incident{}, fmt.Errorf("Failure DAO")
}
func (m *failureIncidentDAO) Insert(incident models.Incident) error {
	return fmt.Errorf("Failure DAO")
}

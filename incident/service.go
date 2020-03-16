package incident

import (
	"time"

	"github.com/involvestecnologia/statuspage/component"
	"github.com/involvestecnologia/statuspage/errors"
	"github.com/involvestecnologia/statuspage/logs"
	"github.com/involvestecnologia/statuspage/models"
)

type incidentService struct {
	component component.Service
	repo      Repository
	log       logs.Log
}

//NewService creates implementation of the Service interface
func NewService(r Repository, component component.Service, l logs.Log) Service {
	return &incidentService{
		component: component,
		repo:      r,
		log:       l,
	}
}

func (s *incidentService) CreateIncidents(incident models.Incident) error {
	if _, err := s.component.FindComponent(map[string]interface{}{"ref": incident.ComponentRef}); err != nil {
		s.log.Error(models.LogFields{"component_ref": incident.ComponentRef}, err.Error())
		return err
	}

	// Look for unresolved related incidents
	openIncidents, err := s.GetUnresolvedIncidents(incident.ComponentRef)
	if err != nil {
		s.log.Error(models.LogFields{"component_ref": incident.ComponentRef}, err.Error())
		return err
	}
	hasOpenIncidents := len(openIncidents) > 0

	switch incident.Status {
	case models.IncidentStatusOK:
		incident.Resolved = true
		if hasOpenIncidents {
			s.closeIncidents(openIncidents...)
		}
		return s.repo.Insert(incident)
	case models.IncidentStatusUnstable, models.IncidentStatusOutage:
		if hasOpenIncidents && openIncidents[0].Status == models.IncidentStatusOutage {
			return &errors.ErrIncidentStatusIgnored{Message: errors.ErrIncidentStatusIgnoredMessage}
		}

		return s.repo.Insert(incident)
	}

	return &errors.ErrUnkownIncidentStatus{Message: errors.ErrUnkownIncidentStatusMessage}
}

func (s *incidentService) UpdateIncident(incident models.Incident) error {
	return s.repo.Update(incident)
}

func (s *incidentService) FindIncidents(query map[string]interface{}) ([]models.Incident, error) {
	return s.repo.Find(query)
}

func (s *incidentService) GetUnresolvedIncidents(componentRef string) ([]models.Incident, error) {
	incidents, err := s.repo.Find(map[string]interface{}{"component_ref": componentRef, "resolved": false})
	if _, ok := err.(*errors.ErrNotFound); ok {
		return incidents, nil
	}
	return incidents, err
}

func (s *incidentService) ListIncidents(queryParameters models.ListIncidentQueryParameters) ([]models.Incident, error) {
	var iComp []models.Incident
	var start time.Time
	var err error
	end := time.Now()

	if queryParameters.StartDate != "" {
		start, err = time.Parse(time.RFC3339, queryParameters.StartDate)
		if err != nil {
			s.log.Error(models.LogFields{"start_date": queryParameters.StartDate, "end_date": queryParameters.EndDate}, err.Error())
			return iComp, err
		}
	}

	if queryParameters.EndDate != "" {
		end, err = time.Parse(time.RFC3339, queryParameters.EndDate)
		if err != nil {
			s.log.Error(models.LogFields{"start_date": queryParameters.StartDate, "end_date": queryParameters.EndDate}, err.Error())
			return iComp, err
		}
	}

	if err := s.ValidateDate(start, end); err != nil {
		s.log.Error(models.LogFields{"start_date": start.String(), "end_date": end.String()}, err.Error())
		return iComp, err
	}

	return s.repo.List(start, end, queryParameters.Unresolved)
}

func (s *incidentService) ValidateDate(startDate, endDate time.Time) error {

	now := time.Now().Add(1 * time.Second)

	if startDate.After(endDate) || startDate.After(now) {
		return &errors.ErrInvalidDate{Message: errors.ErrInvalidDateMessage}
	}

	return nil
}

func (s *incidentService) closeIncidents(incidents ...models.Incident) {
	for _, openIncident := range incidents {
		openIncident.Resolved = true
		openIncident.Duration = time.Since(openIncident.Date)
		if err := s.UpdateIncident(openIncident); err != nil {
			s.log.Error(models.LogFields{"component_ref": openIncident.ComponentRef}, err.Error())
		}
	}
}

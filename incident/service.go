package incident

import (
	"time"

	"github.com/involvestecnologia/statuspage/component"
	"github.com/involvestecnologia/statuspage/errors"
	"github.com/involvestecnologia/statuspage/models"
)

type incidentService struct {
	component component.Service
	repo      Repository
}

//NewService creates implementation of the Service interface
func NewService(r Repository, component component.Service) Service {
	return &incidentService{
		component: component,
		repo:      r,
	}
}

func (s *incidentService) CreateIncidents(incident models.Incident) error {
	_, err := s.component.FindComponent(map[string]interface{}{"ref": incident.ComponentRef})
	if err != nil {
		return err
	}

	if incident.Status == models.IncidentStatusOK {
		// Certify that OK status are already resolved
		incident.Resolved = true
	}

	lastIncident, err := s.GetLastIncident(incident.ComponentRef)
	if err != nil {
		switch err.(type) {
		case *errors.ErrNotFound:
			//No previous incidents found, just create the new incident
			return s.repo.Insert(incident)
		default:
			return err
		}
	}

	if lastIncident.Status == models.IncidentStatusOK {
		unresolvedIncidents, err := s.GetUnresolvedIncidents(incident.ComponentRef, incident.Description)
		if err != nil {
			switch err.(type) {
			case *errors.ErrNotFound:
				//No unresolved incidents found, just create the new incident
				return s.repo.Insert(incident)
			default:
				return err
			}
		}
		for _, inc := range unresolvedIncidents {
			inc.Resolved = true
			inc.Duration = time.Since(lastIncident.Date)
			s.UpdateIncident(inc)   // #nosec
			s.repo.Insert(incident) // #nosec
		}
		return err
	}

	if incident.Status == models.IncidentStatusOK {
		//Last status was NOT OK and new status is OK.
		//Update resolved and duration from last, then create new incident
		lastIncident.Resolved = true
		lastIncident.Duration = time.Since(lastIncident.Date)
		s.UpdateIncident(lastIncident) // #nosec
		return s.repo.Insert(incident)
	}

	if incident.Status > lastIncident.Status {
		//Last status was NOT OK and new status is more critical.
		//Update last incident status.
		lastIncident.Status = incident.Status
		lastIncident.Description = incident.Description
		return s.UpdateIncident(lastIncident)
	}

	return &errors.ErrIncidentStatusIgnored{Message: errors.ErrIncidentStatusIgnoredMessage}
}

func (s *incidentService) UpdateIncident(incident models.Incident) error {
	return s.repo.Update(incident)
}

func (s *incidentService) FindIncidents(query map[string]interface{}) ([]models.Incident, error) {
	return s.repo.Find(query)
}

func (s *incidentService) GetLastIncident(componentRef string) (models.Incident, error) {
	return s.repo.FindOne(map[string]interface{}{"component_ref": componentRef})
}

func (s *incidentService) GetUnresolvedIncidents(componentRef, description string) ([]models.Incident, error) {
	return s.repo.Find(map[string]interface{}{"component_ref": componentRef, "description": description, "resolved": false})
}

func (s *incidentService) ListIncidents(queryParameters models.ListIncidentQueryParameters) ([]models.Incident, error) {
	var iComp []models.Incident
	var start time.Time
	var err error
	end := time.Now()

	if queryParameters.StartDate != "" {
		start, err = time.Parse(time.RFC3339, queryParameters.StartDate)
		if err != nil {
			return iComp, err
		}
	}

	if queryParameters.EndDate != "" {
		end, err = time.Parse(time.RFC3339, queryParameters.EndDate)
		if err != nil {
			return iComp, err
		}
	}

	if err := s.ValidateDate(start, end); err != nil {
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

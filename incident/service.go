package incident

import (
	"strconv"
	"time"

	"github.com/involvestecnologia/statuspage/component"
	"github.com/involvestecnologia/statuspage/errors"
	"github.com/involvestecnologia/statuspage/models"
)

type incidentService struct {
	component component.Service
	repo      Repository
}

func NewService(r Repository, component component.Service) Service {
	return &incidentService{
		component: component,
		repo:      r}
}

func (s *incidentService) CreateIncidents(incident models.Incident) error {
	_, err := s.component.FindComponent(map[string]interface{}{"ref": incident.ComponentRef})
	if err != nil {
		return err
	}

	lastIncident, err := s.GetLastIncident(incident.ComponentRef)
	if err != nil {
		return s.repo.Insert(incident)
	}

	if (incident.Status == models.IncidentStatusOK) && (lastIncident.Status != models.IncidentStatusOK) {
		lastIncident.Status = incident.Status
		lastIncident.Duration = incident.Duration
		lastIncident.Resolved = true
		return s.UpdateIncident(lastIncident)
	}

	if incident.Status > lastIncident.Status {
		lastIncident.Status = incident.Status
		lastIncident.Description = incident.Description
		lastIncident.Resolved = false
		return s.UpdateIncident(lastIncident)
	}

	return &errors.ErrIncidentStatusIgnored{errors.ErrIncidentStatusIgnoredMessage}
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

func (s *incidentService) ListIncidents(year string, month string) ([]models.Incident, error) {
	var iComp []models.Incident
	var start, end time.Time
	if year == "" && month == "" {
		return s.repo.List(start, end)
	}

	m, err := s.ValidateMonth(month)
	if err != nil && month != "" {
		return iComp, err
	}

	y, err := s.ValidateYear(year)
	if err != nil && year != "" {
		return iComp, err
	}

	if y == -1 {
		y = time.Now().Year()
	}

	if m == -1 {
		start = time.Date(y, 1, 1, 0, 0, 0, 0, time.UTC)
		end = time.Date(y+1, 1, 0, 23, 59, 59, 0, time.UTC)
		return s.repo.List(start, end)
	}

	start = time.Date(y, time.Month(m), 1, 0, 0, 0, 0, time.UTC)
	end = time.Date(y, time.Month(m+1), 0, 23, 59, 59, 0, time.UTC)

	return s.repo.List(start, end)
}

func (s *incidentService) ValidateMonth(month string) (int, error) {
	m, err := strconv.Atoi(month)
	if err != nil {
		return -1, err
	}
	valid := m > 0 && m < 13
	if !valid {
		return -1, &errors.ErrInvalidMonth{Message: errors.ErrInvalidMonthMessage}
	}
	return m, nil
}

func (s *incidentService) ValidateYear(year string) (int, error) {
	y, err := strconv.Atoi(year)
	if err != nil {
		return -1, err
	}
	valid := y > 0 && y <= time.Now().Year()
	if !valid {
		return -1, &errors.ErrInvalidYear{Message: errors.ErrInvalidYearMessage}
	}
	return y, nil
}

package incident

import (
	"time"

	"github.com/involvestecnologia/statuspage/models"
)

// Read implements the read action methods
type Read interface {
	Find(query map[string]interface{}) ([]models.Incident, error)
	FindOne(query map[string]interface{}) (models.Incident, error)
	List(startDt time.Time, endDt time.Time, unresolved bool) ([]models.Incident, error)
}

// Write implements the write action methods
type Write interface {
	Insert(incident models.Incident) error
	Update(incident models.Incident) error
}

// Repository describes the repository where the data will be written and read from
type Repository interface {
	Read
	Write
}

// Service describes the use case
type Service interface {
	CreateIncidents(incident models.Incident) error
	UpdateIncident(incident models.Incident) error
	GetLastIncident(componentRef string) (models.Incident, error)
	FindIncidents(query map[string]interface{}) ([]models.Incident, error)
	ListIncidents(year string, month string, unresolved bool) ([]models.Incident, error)
	ValidateMonth(monthArg string) (int, error)
	ValidateYear(yearArg string) (int, error)
}

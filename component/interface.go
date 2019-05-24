package component

import (
	"github.com/involvestecnologia/statuspage/models"
)

// Read implements the read action methods
type Read interface {
	Find(queryParam map[string]interface{}) (models.Component, error)
	FindAllWithLabel(label string) ([]models.Component, error)
	List() ([]models.Component, error)
	ListAllLabels() (models.ComponentLabels, error)
}

// Write implements the write action methods
type Write interface {
	Delete(componentRef string) error
	Insert(component models.Component) (string, error)
	Update(componentRef string, component models.Component) error
}

// Repository describes the repository where the data will be written and read from
type Repository interface {
	Read
	Write
}

// Service describes the use case
type Service interface {
	ComponentExists(map[string]interface{}) (models.Component, bool)
	CreateComponent(component models.Component) (string, error)
	FindComponent(queryParam map[string]interface{}) (models.Component, error)
	ListComponents(refs []string) ([]models.Component, error)
	RemoveComponent(componentRef string) error
	UpdateComponent(componentRef string, component models.Component) error
	ListAllLabels() (models.ComponentLabels, error)
	ListComponentsWithLabels(labels models.ComponentLabels) ([]models.Component, error)
}

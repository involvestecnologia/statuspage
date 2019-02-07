package mock

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/involvestecnologia/statuspage/component"
	"github.com/involvestecnologia/statuspage/errors"
	"github.com/involvestecnologia/statuspage/models"
)

type componentDAO struct {
	components []models.Component
}

// NewMockComponentDAO  creates a new Component Repository Mock
func NewMockComponentDAO() component.Repository {
	return &componentDAO{
		components: []models.Component{
			{
				Ref:     ZeroTimeHex,
				Name:    "first",
				Address: "",
			},
			{
				Ref:     bson.NewObjectIdWithTime(bson.Now()).Hex(),
				Name:    "first_comp_with_group",
				Address: "",
			},
			{
				Ref:     "Empty Component",
				Name:    "Empty test",
				Address: "",
			},
			{
				Ref:     bson.NewObjectIdWithTime(bson.Now()).Hex(),
				Name:    "test",
				Address: "",
			},
			{
				Ref:     bson.NewObjectIdWithTime(bson.Now()).Hex(),
				Name:    "last_comp_with_group",
				Address: "",
			},
			{
				Ref:     bson.NewObjectIdWithTime(bson.Now()).Hex(),
				Name:    "last",
				Address: "",
			},
		},
	}
}

func (m *componentDAO) List() ([]models.Component, error) {
	return m.components, nil
}
func (m *componentDAO) Find(q map[string]interface{}) (models.Component, error) {
	var c models.Component
	if keyValue, hasKey := q["ref"]; hasKey {
		for _, c := range m.components {
			if c.Ref == keyValue {
				return c, nil
			}
		}
	} else {
		if keyValue, hasKey := q["name"]; hasKey {
			for _, c := range m.components {
				if c.Name == keyValue {
					return c, nil
				}
			}
		} else {
			return c, &errors.ErrInvalidQuery{Message: errors.ErrInvalidQueryMessage}
		}
	}

	return c, &errors.ErrNotFound{Message: errors.ErrNotFoundMessage}
}
func (m *componentDAO) Insert(component models.Component) (string, error) {
	if component.Ref == "" {
		component.Ref = bson.NewObjectId().Hex()
	}
	m.components = append(m.components, component)
	return component.Ref, nil
}
func (m *componentDAO) Update(ref string, component models.Component) error {
	for k, comp := range m.components {
		if comp.Ref == ref {
			m.components[k].Name = component.Name
			m.components[k].Address = component.Address
			return nil
		}
	}
	return &errors.ErrNotFound{Message: errors.ErrNotFoundMessage}
}
func (m *componentDAO) Delete(ref string) error {
	for k, comp := range m.components {
		if comp.Ref == ref {
			m.components = append(m.components[:k], m.components[k+1:]...)
			return nil
		}
	}
	return &errors.ErrNotFound{Message: errors.ErrNotFoundMessage}
}
func (m *componentDAO) FindAllWithLabel(label string) ([]models.Component, error) {
	var comps []models.Component
	for _, c := range m.components {
		for _, l := range c.Labels {
			if l == label {
				comps = append(comps, c)
			}
		}
	}

	return comps, nil
}
func (m *componentDAO) ListAllLabels() (models.ComponentLabels, error) {
	cLabels := models.ComponentLabels{
		Labels: make([]string, 0),
	}
	labelExist := make(map[string]bool, 0)
	for _, c := range m.components {
		for _, l := range c.Labels {
			if !labelExist[l] {
				labelExist[l] = true
				cLabels.Labels = append(cLabels.Labels, l)
			}
		}
	}
	return cLabels, nil
}

type failureComponentDAO struct {
}

// NewMockFailureComponentDAO creates a new Component Repository Mock that fails in every operation
func NewMockFailureComponentDAO() component.Repository {
	return &failureComponentDAO{}
}
func (m *failureComponentDAO) List() ([]models.Component, error) {
	return []models.Component{}, fmt.Errorf("DAO Failure")
}
func (m *failureComponentDAO) ListAllLabels() (models.ComponentLabels, error) {
	return models.ComponentLabels{}, fmt.Errorf("DAO Failure")
}
func (m *failureComponentDAO) Find(q map[string]interface{}) (models.Component, error) {
	return models.Component{}, fmt.Errorf("DAO Failure")
}
func (m *failureComponentDAO) FindAllWithLabel(label string) ([]models.Component, error) {
	return []models.Component{}, fmt.Errorf("DAO Failure")
}
func (m *failureComponentDAO) Insert(component models.Component) (string, error) {
	return "", fmt.Errorf("DAO Failure")
}
func (m *failureComponentDAO) Update(ref string, component models.Component) error {
	return fmt.Errorf("DAO Failure")
}
func (m *failureComponentDAO) Delete(ref string) error {
	return fmt.Errorf("DAO Failure")
}

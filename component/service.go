package component

import (
	"github.com/involvestecnologia/statuspage/errors"
	"github.com/involvestecnologia/statuspage/models"
)

type componentService struct {
	repo Repository
}

//NewService creates implementation of the Service interface
func NewService(r Repository) Service {
	return &componentService{repo: r}
}

func (s *componentService) CreateComponent(component models.Component) (string, error) {
	if valid, err := s.isValidComponent(component); !valid {
		return component.Ref, err
	}

	if inUse, err := s.isComponentNameInUse(component); inUse {
		return component.Ref, err
	}

	if inUse, err := s.isComponentRefInUse(component); inUse {
		return component.Ref, err
	}

	return s.repo.Insert(component)
}

func (s *componentService) UpdateComponent(ref string, component models.Component) error {
	component.Ref = ref
	if valid, err := s.isValidComponent(component); !valid {
		return err
	}

	if inUse, err := s.isComponentNameInUse(component); inUse {
		return err
	}
	return s.repo.Update(ref, component)
}

func (s *componentService) FindComponent(queryParam map[string]interface{}) (models.Component, error) {
	if len(queryParam) == 0 {
		return models.Component{}, &errors.ErrInvalidQuery{Message: errors.ErrInvalidQueryMessage}
	}
	return s.repo.Find(queryParam)
}

func (s *componentService) ListComponents(refs []string) ([]models.Component, error) {
	if len(refs) == 0 {
		return s.repo.List()
	}
	comps := make([]models.Component, 0)
	for _, r := range refs {
		if c, exist := s.ComponentExists(map[string]interface{}{"ref": r}); exist {
			comps = append(comps, c)
		} else {
			return []models.Component{}, &errors.ErrNotFound{Message: errors.ErrNotFoundMessage}
		}
	}
	return comps, nil
}

func (s *componentService) RemoveComponent(id string) error {
	return s.repo.Delete(id)
}

func (s *componentService) ComponentExists(componentFields map[string]interface{}) (models.Component, bool) {
	c, err := s.repo.Find(componentFields)
	return c, err == nil
}

func (s *componentService) isValidComponent(c models.Component) (bool, error) {
	if c.Name == "" {
		return false, &errors.ErrComponentNameIsEmpty{Message: errors.ErrComponentNameIsEmptyMessage}
	}
	return true, nil
}

func (s *componentService) isComponentNameInUse(c models.Component) (bool, error) {
	if comp, exist := s.ComponentExists(map[string]interface{}{"name": c.Name}); exist && comp.Ref != c.Ref {
		return true, &errors.ErrComponentNameAlreadyExists{Message: errors.ErrComponentRefAlreadyExistsMessage}
	}
	return false, nil
}

func (s *componentService) isComponentRefInUse(c models.Component) (bool, error) {
	if _, exist := s.ComponentExists(map[string]interface{}{"ref": c.Ref}); exist {
		return true, &errors.ErrComponentRefAlreadyExists{Message: errors.ErrComponentRefAlreadyExistsMessage}
	}
	return false, nil
}

func (s *componentService) ListAllLabels() (models.ComponentLabels, error) {
	return s.repo.ListAllLabels()
}
func (s *componentService) ListComponentsWithLabels(cLabels models.ComponentLabels) ([]models.Component, error) {
	var components []models.Component
	for _, label := range cLabels.Labels {
		comps, err := s.repo.FindAllWithLabel(label)
		if err != nil {
			return nil, err
		}
		for _, c := range comps {
			components = append(components, c)
		}

	}
	return components, nil
}

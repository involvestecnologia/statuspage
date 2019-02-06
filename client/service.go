package client

import (
	"fmt"

	"github.com/involvestecnologia/statuspage/component"

	"github.com/involvestecnologia/statuspage/errors"
	"github.com/involvestecnologia/statuspage/models"
)

type clientService struct {
	repo      Repository
	component component.Service
}

//NewService creates implementation of the Service interface
func NewService(r Repository, component component.Service) Service {
	return &clientService{
		repo:      r,
		component: component,
	}
}

//CreateClient create a new client as long as the name passed it's not already in use by
//another client, all component ref do exists and in case a special Reference is passed
//it shouldn't be in use aswell.
func (s *clientService) CreateClient(client models.Client) (string, error) {
	if s.clientAlreadyExists(map[string]interface{}{"name": client.Name}) {
		return client.Name, &errors.ErrClientNameAlreadyExists{Message: errors.ErrClientNameAlreadyExistsMessage}
	}
	if client.Ref != "" {
		if s.clientAlreadyExists(map[string]interface{}{"ref": client.Ref}) {
			return client.Ref, &errors.ErrClientRefAlreadyExists{Message: errors.ErrClientRefAlreadyExistsMessage}
		}
	}
	for _, compRef := range client.Resources {
		if _, exists := s.component.ComponentExists(map[string]interface{}{"ref": compRef}); !exists {
			return client.Ref, &errors.ErrInvalidRef{Message: fmt.Sprintf(errors.ErrInvalidRefMessage, compRef)}
		}
	}

	return s.repo.Insert(client)
}

//UpdateClient updates a existent client as long as the name passed it's not already in use by
//another client and all component ref do exists.
func (s *clientService) UpdateClient(clientRef string, client models.Client) error {
	client.Ref = clientRef
	c, err := s.FindClient(map[string]interface{}{"name": client.Name})
	if (err == nil) && (c.Ref != clientRef) {
		return &errors.ErrClientNameAlreadyExists{Message: errors.ErrClientNameAlreadyExistsMessage}
	}
	for _, compRef := range client.Resources {
		if _, exists := s.component.ComponentExists(map[string]interface{}{"ref": compRef}); !exists {
			return &errors.ErrInvalidRef{Message: fmt.Sprintf(errors.ErrInvalidRefMessage, compRef)}
		}
	}
	return s.repo.Update(clientRef, client)
}

//FindClient search for a client that matches a attribute:value query map.
func (s *clientService) FindClient(queryParam map[string]interface{}) (models.Client, error) {
	if len(queryParam) == 0 {
		return models.Client{}, &errors.ErrInvalidQuery{Message: errors.ErrInvalidQueryMessage}
	}
	return s.repo.Find(queryParam)
}

//RemoveCliente deletes a clients that matches the passed reference.
func (s *clientService) RemoveClient(clientRef string) error {
	return s.repo.Delete(clientRef)
}

//ListClients retrieves all existent clients
func (s *clientService) ListClients() ([]models.Client, error) {
	return s.repo.List()
}

func (s *clientService) clientAlreadyExists(clientFields map[string]interface{}) bool {
	_, err := s.FindClient(clientFields)
	return err == nil
}

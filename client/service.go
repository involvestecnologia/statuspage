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

func NewService(r Repository, component component.Service) Service {
	return &clientService{
		repo:      r,
		component: component,
	}
}

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

func (s *clientService) UpdateClient(clientRef string, client models.Client) error {
	for _, compRef := range client.Resources {
		if _, exists := s.component.ComponentExists(map[string]interface{}{"ref": compRef}); !exists {
			return &errors.ErrInvalidRef{Message: fmt.Sprintf(errors.ErrInvalidRefMessage, compRef)}
		}
	}
	return s.repo.Update(clientRef, client)
}

func (s *clientService) FindClient(queryParam map[string]interface{}) (models.Client, error) {
	if len(queryParam) == 0 {
		return models.Client{}, &errors.ErrInvalidQuery{Message: errors.ErrInvalidQueryMessage}
	}
	return s.repo.Find(queryParam)
}

func (s *clientService) RemoveClient(clientID string) error {
	return s.repo.Delete(clientID)
}

func (s *clientService) ListClients() ([]models.Client, error) {
	return s.repo.List()
}

func (s *clientService) clientAlreadyExists(clientFields map[string]interface{}) bool {
	_, err := s.FindClient(clientFields)
	return err == nil
}

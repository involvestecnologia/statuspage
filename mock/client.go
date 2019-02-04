package mock

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/involvestecnologia/statuspage/client"
	"github.com/involvestecnologia/statuspage/errors"
	"github.com/involvestecnologia/statuspage/models"
)

type MockClientDAO struct {
	clients []models.Client
}

func NewMockClientDAO() client.Repository {
	return &MockClientDAO{
		clients: []models.Client{
			{
				Ref:  ZeroTimeHex,
				Name: "First Client",
				Resources: []string{
					bson.NewObjectIdWithTime(bson.Now()).Hex(),
					bson.NewObjectIdWithTime(bson.Now()).Hex(),
					bson.NewObjectIdWithTime(bson.Now()).Hex(),
				},
			},
			{
				Ref:  OneSecTimeHex,
				Name: "Last Client",
				Resources: []string{
					bson.NewObjectIdWithTime(bson.Now()).Hex(),
				},
			},
		},
	}
}

func (m *MockClientDAO) Delete(clientRef string) error {
	for i, c := range m.clients {
		if c.Ref == clientRef {
			m.clients = append(m.clients[:i], m.clients[i+1:]...)
			return nil
		}
	}
	return &errors.ErrNotFound{Message: errors.ErrNotFoundMessage}
}

func (m *MockClientDAO) Find(q map[string]interface{}) (models.Client, error) {
	if keyValue, hasKey := q["ref"]; hasKey {
		for _, c := range m.clients {
			if c.Ref == keyValue {
				return c, nil
			}
		}
	} else {
		if keyValue, hasKey := q["name"]; hasKey {
			for _, c := range m.clients {
				if c.Name == keyValue {
					return c, nil
				}
			}
		} else {
			return models.Client{}, &errors.ErrNotFound{Message: errors.ErrNotFoundMessage}
		}
	}

	return models.Client{}, &errors.ErrNotFound{Message: errors.ErrNotFoundMessage}
}

func (m *MockClientDAO) Insert(client models.Client) (string, error) {
	if client.Ref == "" {
		client.Ref = bson.NewObjectId().Hex()
	}
	m.clients = append(m.clients, client)
	return client.Ref, nil
}

func (m *MockClientDAO) List() ([]models.Client, error) {
	return m.clients, nil
}

func (m *MockClientDAO) Update(clientRef string, client models.Client) error {
	for i, c := range m.clients {
		if c.Ref == clientRef {
			m.clients[i].Name = client.Name
			m.clients[i].Resources = client.Resources
			return nil
		}
	}
	return &errors.ErrNotFound{Message: errors.ErrNotFoundMessage}
}

type MockFailureClientDAO struct {
}

func NewMockFailureClientDAO() client.Repository {
	return &MockFailureClientDAO{}
}

func (f *MockFailureClientDAO) Find(q map[string]interface{}) (models.Client, error) {
	return models.Client{}, fmt.Errorf("DAO Failure")
}
func (f *MockFailureClientDAO) Delete(clientRef string) error {
	return fmt.Errorf("DAO Failure")
}
func (f *MockFailureClientDAO) Insert(client models.Client) (string, error) {
	return "", fmt.Errorf("DAO Failure")
}
func (f *MockFailureClientDAO) List() ([]models.Client, error) {
	return []models.Client{}, fmt.Errorf("DAO Failure")
}
func (f *MockFailureClientDAO) Update(clientRef string, client models.Client) error {
	return fmt.Errorf("DAO Failure")
}

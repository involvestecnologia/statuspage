package mock

import (
	"fmt"

	"github.com/globalsign/mgo/bson"
	"github.com/involvestecnologia/statuspage/client"
	"github.com/involvestecnologia/statuspage/errors"
	"github.com/involvestecnologia/statuspage/models"
)

type clientDAO struct {
	clients []models.Client
}

//NewMockClientDAO creates a new Client Repository Mock
func NewMockClientDAO() client.Repository {
	return &clientDAO{
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

func (m *clientDAO) Delete(clientRef string) error {
	for i, c := range m.clients {
		if c.Ref == clientRef {
			m.clients = append(m.clients[:i], m.clients[i+1:]...)
			return nil
		}
	}
	return &errors.ErrNotFound{Message: errors.ErrNotFoundMessage}
}

func (m *clientDAO) Find(q map[string]interface{}) (models.Client, error) {
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

func (m *clientDAO) Insert(client models.Client) (string, error) {
	if client.Ref == "" {
		client.Ref = bson.NewObjectId().Hex()
	}
	m.clients = append(m.clients, client)
	return client.Ref, nil
}

func (m *clientDAO) List() ([]models.Client, error) {
	return m.clients, nil
}

func (m *clientDAO) Update(clientRef string, client models.Client) error {
	for i, c := range m.clients {
		if c.Ref == clientRef {
			m.clients[i].Name = client.Name
			m.clients[i].Resources = client.Resources
			return nil
		}
	}
	return &errors.ErrNotFound{Message: errors.ErrNotFoundMessage}
}

type failureClientDAO struct {
}

//NewMockFailureClientDAO creates a new Client Repository Mock that fails in every operation
func NewMockFailureClientDAO() client.Repository {
	return &failureClientDAO{}
}

func (f *failureClientDAO) Find(q map[string]interface{}) (models.Client, error) {
	return models.Client{}, fmt.Errorf("DAO Failure")
}
func (f *failureClientDAO) Delete(clientRef string) error {
	return fmt.Errorf("DAO Failure")
}
func (f *failureClientDAO) Insert(client models.Client) (string, error) {
	return "", fmt.Errorf("DAO Failure")
}
func (f *failureClientDAO) List() ([]models.Client, error) {
	return []models.Client{}, fmt.Errorf("DAO Failure")
}
func (f *failureClientDAO) Update(clientRef string, client models.Client) error {
	return fmt.Errorf("DAO Failure")
}

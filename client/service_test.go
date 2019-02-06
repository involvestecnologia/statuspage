package client_test

import (
	"testing"

	"github.com/globalsign/mgo/bson"
	"github.com/involvestecnologia/statuspage/client"
	"github.com/involvestecnologia/statuspage/component"
	"github.com/involvestecnologia/statuspage/mock"
	"github.com/involvestecnologia/statuspage/models"

	"github.com/stretchr/testify/assert"
)

func TestNewClientService(t *testing.T) {
	dao := mock.NewMockClientDAO()
	assert.Implements(t, (*client.Service)(nil), client.NewService(dao, component.NewService(mock.NewMockComponentDAO())))
}

func TestClientService_CreateClient(t *testing.T) {
	s := client.NewService(mock.NewMockClientDAO(), component.NewService(mock.NewMockComponentDAO()))
	c := models.Client{
		Name:      "test",
		Resources: make([]string, 0),
	}

	ref, err := s.CreateClient(c)
	c.Ref = ref

	assert.Nil(t, err)
	c2, err := s.FindClient(map[string]interface{}{"ref": c.Ref})
	if assert.Nil(t, err) && assert.NotNil(t, c2) {
		if assert.IsType(t, models.Client{}, c2) {
			assert.Equal(t, c, c2)
		}
	}

	_, err = s.CreateClient(c)
	assert.NotNil(t, err)

	c3 := models.Client{
		Name:      "test2",
		Resources: []string{"Invalid Ref", "Another invalid Ref"},
	}
	_, err = s.CreateClient(c3)
	assert.NotNil(t, err)

}
func TestClientService_FindClient(t *testing.T) {
	s := client.NewService(mock.NewMockClientDAO(), component.NewService(mock.NewMockComponentDAO()))

	c, err := s.FindClient(map[string]interface{}{"ref": mock.ZeroTimeHex})
	if assert.Nil(t, err) && assert.NotNil(t, c) {
		if assert.IsType(t, models.Client{}, c) {
			assert.Equal(t, mock.ZeroTimeHex, c.Ref)
		}
	}

	_, err = s.FindClient(map[string]interface{}{"ref": bson.NewObjectId().Hex()})
	assert.NotNil(t, err)

	c, err = s.FindClient(map[string]interface{}{"name": "First Client"})
	if assert.Nil(t, err) && assert.NotNil(t, c) {
		if assert.IsType(t, models.Client{}, c) {
			assert.Equal(t, "First Client", c.Name)
		}
	}

	_, err = s.FindClient(map[string]interface{}{"name": "test"})
	assert.NotNil(t, err)

	_, err = s.FindClient(map[string]interface{}{})
	assert.NotNil(t, err)
}
func TestClientService_ListClients(t *testing.T) {
	s := client.NewService(mock.NewMockClientDAO(), component.NewService(mock.NewMockComponentDAO()))

	c, err := s.ListClients()
	if assert.Nil(t, err) && assert.NotNil(t, c) {
		if assert.IsType(t, []models.Client{}, c) {
			assert.Equal(t, mock.ZeroTimeHex, c[0].Ref)
			assert.Equal(t, mock.OneSecTimeHex, c[len(c)-1].Ref)
		}
	}
}
func TestClientService_RemoveClient(t *testing.T) {
	s := client.NewService(mock.NewMockClientDAO(), component.NewService(mock.NewMockComponentDAO()))
	err := s.RemoveClient(mock.OneSecTimeHex)
	assert.Nil(t, err)

	_, err = s.FindClient(map[string]interface{}{"ref": mock.OneSecTimeHex})
	assert.NotNil(t, err)
}
func TestClientService_UpdateClient(t *testing.T) {
	s := client.NewService(mock.NewMockClientDAO(), component.NewService(mock.NewMockComponentDAO()))

	c, err := s.FindClient(map[string]interface{}{"ref": mock.ZeroTimeHex})
	assert.Nil(t, err)
	assert.NotNil(t, c)

	c.Name = "Modified First Client"

	err = s.UpdateClient(mock.ZeroTimeHex, c)
	assert.Nil(t, err)

	c2, err := s.FindClient(map[string]interface{}{"ref": mock.ZeroTimeHex})
	assert.Nil(t, err)
	assert.NotNil(t, c2)

	assert.Equal(t, c, c2)

	c.Name = "Last Client"
	err = s.UpdateClient(mock.ZeroTimeHex, c)
	assert.NotNil(t, err)

	c3 := models.Client{
		Ref:       c.Ref,
		Name:      "test3",
		Resources: []string{"Invalid Ref", "Another invalid Ref"},
	}
	err = s.UpdateClient(mock.ZeroTimeHex, c3)
	assert.NotNil(t, err)
}

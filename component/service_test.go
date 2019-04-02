package component_test

import (
	"testing"
	"time"

	"github.com/globalsign/mgo/bson"

	"github.com/involvestecnologia/statuspage/component"
	"github.com/involvestecnologia/statuspage/mock"
	"github.com/involvestecnologia/statuspage/models"

	"github.com/stretchr/testify/assert"
)

func TestNewComponentService(t *testing.T) {
	dao := mock.NewMockComponentDAO()
	s := component.NewService(dao, mock.NewComponentLogRepositoryMock())
	assert.Implements(t, (*component.Service)(nil), s)

}

func TestComponentService_ListComponents(t *testing.T) {
	s := component.NewService(mock.NewMockComponentDAO(), mock.NewComponentLogRepositoryMock())

	refs := make([]string, 0)
	c, err := s.ListComponents(refs)
	if assert.Nil(t, err) && assert.NotNil(t, c) {
		if assert.IsType(t, []models.Component{}, c) {
			assert.Equal(t, "first", c[0].Name)
			assert.Equal(t, "last", c[len(c)-1].Name)
		}
	}
}
func TestComponentService_FindComponent(t *testing.T) {
	s := component.NewService(mock.NewMockComponentDAO(), mock.NewComponentLogRepositoryMock())

	c, err := s.FindComponent(map[string]interface{}{"ref": mock.ZeroTimeHex})
	if assert.Nil(t, err) && assert.NotNil(t, c) {
		if assert.IsType(t, models.Component{}, c) {
			assert.Equal(t, "first", c.Name)
		}
	}

	c2, err := s.FindComponent(map[string]interface{}{"name": c.Name})
	if assert.Nil(t, err) && assert.NotNil(t, c) {
		if assert.IsType(t, models.Component{}, c) {
			assert.Equal(t, c, c2)
		}
	}

	_, err = s.FindComponent(map[string]interface{}{"ref": mock.OneSecTimeHex})
	assert.NotNil(t, err)

	_, err = s.FindComponent(map[string]interface{}{"name": mock.OneSecTimeHex})
	assert.NotNil(t, err)

	_, err = s.FindComponent(map[string]interface{}{"ref": "test"})
	assert.NotNil(t, err)

	_, err = s.FindComponent(map[string]interface{}{"name": ""})
	assert.NotNil(t, err)

	_, err = s.FindComponent(map[string]interface{}{"ref": ""})
	assert.NotNil(t, err)

	_, err = s.FindComponent(map[string]interface{}{})
	assert.NotNil(t, err)

}
func TestComponentService_CreateComponent(t *testing.T) {
	s := component.NewService(mock.NewMockComponentDAO(), mock.NewComponentLogRepositoryMock())

	c := models.Component{
		Ref:     bson.NewObjectIdWithTime(bson.Now().Add(5 * time.Second)).Hex(),
		Name:    "New Component",
		Address: "no-address",
		Labels:  []string{"test", "test2"},
	}

	ref, err := s.CreateComponent(c)
	c.Ref = ref
	assert.Nil(t, err)

	comp, err := s.FindComponent(map[string]interface{}{"name": c.Name})
	if assert.Nil(t, err) && assert.NotNil(t, comp) {
		if assert.IsType(t, models.Component{}, comp) {
			assert.Equal(t, c, comp)
		}
	}

	c.Ref = mock.ZeroTimeHex
	_, err = s.CreateComponent(c)
	assert.NotNil(t, err)

	c.Ref = ""
	c.Name = ""
	_, err = s.CreateComponent(c)
	assert.NotNil(t, err)

}
func TestComponentService_UpdateComponent(t *testing.T) {
	s := component.NewService(mock.NewMockComponentDAO(), mock.NewComponentLogRepositoryMock())

	currTime := bson.Now().String()
	c, err := s.FindComponent(map[string]interface{}{"ref": mock.ZeroTimeHex})
	if assert.Nil(t, err) && assert.NotNil(t, c) {
		c.Address = currTime
	}
	err = s.UpdateComponent(mock.ZeroTimeHex, c)
	assert.Nil(t, err)

	comp, err := s.FindComponent(map[string]interface{}{"ref": mock.ZeroTimeHex})
	if assert.Nil(t, err) && assert.NotNil(t, comp) {
		assert.Equal(t, c, comp)
	}

	err = s.UpdateComponent(mock.OneSecTimeHex, c)
	assert.NotNil(t, err)

	c.Name = ""
	err = s.UpdateComponent(mock.OneSecTimeHex, c)
	assert.NotNil(t, err)

}
func TestComponentService_RemoveComponent(t *testing.T) {
	s := component.NewService(mock.NewMockComponentDAO(), mock.NewComponentLogRepositoryMock())

	err := s.RemoveComponent(mock.ZeroTimeHex)
	assert.Nil(t, err)

	_, err = s.FindComponent(map[string]interface{}{"ref": mock.ZeroTimeHex})
	assert.NotNil(t, err)
}
func TestComponentService_componentExists(t *testing.T) {
	s := component.NewService(mock.NewMockComponentDAO(), mock.NewComponentLogRepositoryMock())

	c, err := s.FindComponent(map[string]interface{}{"ref": mock.ZeroTimeHex})
	if assert.Nil(t, err) && assert.NotNil(t, c) {
		if assert.IsType(t, models.Component{}, c) {
			assert.Equal(t, "first", c.Name)
		}
	}
	_, exists := s.ComponentExists(map[string]interface{}{"name": c.Name})
	assert.True(t, exists)

	_, exists = s.ComponentExists(map[string]interface{}{"name": bson.NewObjectIdWithTime(bson.Now()).String()})
	assert.False(t, exists)
}

func TestComponentService_ListAllLabels(t *testing.T) {
	s := component.NewService(mock.NewMockComponentDAO(), mock.NewComponentLogRepositoryMock())

	c := models.Component{
		Ref:     bson.NewObjectIdWithTime(bson.Now().Add(5 * time.Second)).Hex(),
		Name:    "New Component",
		Address: "no-address",
		Labels:  []string{"test", "test2"},
	}
	_, err := s.CreateComponent(c)
	assert.Nil(t, err)

	cLabels, err := s.ListAllLabels()
	assert.Nil(t, err)
	if assert.NotNil(t, cLabels) {
		assert.NotEmpty(t, cLabels.Labels)
	}

}
func TestComponentService_ListComponentsWithLabels(t *testing.T) {
	s := component.NewService(mock.NewMockComponentDAO(), mock.NewComponentLogRepositoryMock())

	c := models.Component{
		Name:    "New Component",
		Address: "no-address",
		Labels:  []string{"test", "test2"},
	}
	_, err := s.CreateComponent(c)
	assert.Nil(t, err)

	cLabels, err := s.ListAllLabels()
	assert.Nil(t, err)
	if assert.NotNil(t, cLabels) {
		assert.NotEmpty(t, cLabels.Labels)
	}

	comps, err := s.ListComponentsWithLabels(cLabels)
	if assert.Nil(t, err) && assert.NotNil(t, comps) {
		assert.IsType(t, []models.Component{}, comps)
	}
}

package client_test

import (
	"log"
	"testing"

	mgo "github.com/globalsign/mgo"
	"github.com/involvestecnologia/statuspage/client"
	"github.com/involvestecnologia/statuspage/db"
	"github.com/involvestecnologia/statuspage/models"
	"github.com/stretchr/testify/assert"
)

const validMongoArgs = "localhost"

var testSession *mgo.Session
var failureSession *mgo.Session
var c = models.Client{
	Name:      "Test Client",
	Ref:       "",
	Resources: make([]string, 0),
}

func init() {
	testSession = db.InitMongo(validMongoArgs)
	err := testSession.DB("status").DropDatabase()
	if err != nil {
		log.Panicf("%s\n", err)
	}
}

func TestClientMongoDB_Repository_NewMongoRepository(t *testing.T) {
	repo := client.NewMongoRepository(testSession)
	assert.Implements(t, (*client.Repository)(nil), repo)
}

func TestClientMongoDB_Repository_Insert(t *testing.T) {

	repo := client.NewMongoRepository(testSession)
	ref, err := repo.Insert(c)
	c.Ref = ref
	assert.Nil(t, err)
	c2, err := repo.Find(map[string]interface{}{"ref": c.Ref})
	if assert.Nil(t, err) && assert.NotNil(t, c2) {
		assert.Equal(t, c, c2)
	}

	repo = client.NewMongoRepository(failureSession)
	_, err = repo.Insert(c)
	assert.NotNil(t, err)

}

func TestClientMongoDB_Repository_Update(t *testing.T) {
	repo := client.NewMongoRepository(testSession)

	c.Name = "Updated Test Client"

	err := repo.Update(c.Ref, c)
	assert.Nil(t, err)

	c2, err := repo.Find(map[string]interface{}{"ref": c.Ref})
	if assert.Nil(t, err) && assert.NotNil(t, c2) {
		assert.Equal(t, c.Name, c2.Name)
	}

	err = repo.Update("Invalid Ref Client", c)
	assert.NotNil(t, err)

	repo = client.NewMongoRepository(failureSession)
	err = repo.Update(c.Ref, c)
	assert.NotNil(t, err)

}

func TestClientMongoDB_Repository_Find(t *testing.T) {
	repo := client.NewMongoRepository(testSession)
	c2, err := repo.Find(map[string]interface{}{"ref": c.Ref})
	if assert.Nil(t, err) && assert.NotNil(t, c2) {
		assert.Equal(t, c, c2)
	}

	c2, err = repo.Find(map[string]interface{}{"name": c.Name})
	if assert.Nil(t, err) && assert.NotNil(t, c2) {
		assert.Equal(t, c.Name, c2.Name)
	}

	_, err = repo.Find(map[string]interface{}{"ref": c.Name})
	assert.NotNil(t, err)

	_, err = repo.Find(map[string]interface{}{"name": "test"})
	assert.NotNil(t, err)

	repo = client.NewMongoRepository(failureSession)
	_, err = repo.Find(map[string]interface{}{"ref": c.Ref})
	assert.NotNil(t, err)
}

func TestClientMongoDB_Repository_Delete(t *testing.T) {
	repo := client.NewMongoRepository(testSession)
	c2, err := repo.Find(map[string]interface{}{"ref": c.Ref})
	if assert.Nil(t, err) && assert.NotNil(t, c2) {
		assert.Equal(t, c, c2)
	}

	err = repo.Delete(c.Ref)
	assert.Nil(t, err)

	_, err = repo.Find(map[string]interface{}{"ref": c.Ref})
	assert.NotNil(t, err)

	err = repo.Delete(c.Ref)
	assert.NotNil(t, err)

	err = repo.Delete(c.Name)
	assert.NotNil(t, err)

	repo = client.NewMongoRepository(failureSession)
	err = repo.Delete(c.Ref)
	assert.NotNil(t, err)
}

func TestClientMongoDB_Repository_List(t *testing.T) {
	repo := client.NewMongoRepository(testSession)

	clients, err := repo.List()
	assert.Nil(t, clients)
	assert.Nil(t, err)

	ref, err := repo.Insert(c)
	c.Ref = ref
	assert.Nil(t, err)

	clients, err = repo.List()
	if assert.Nil(t, err) && assert.NotNil(t, clients) {
		list := []models.Client{c}
		assert.IsType(t, list, clients)
		assert.Equal(t, list, clients)
	}

	repo = client.NewMongoRepository(failureSession)
	_, err = repo.List()
	assert.NotNil(t, err)
}

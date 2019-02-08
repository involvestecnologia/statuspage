package db

import (
	"os"
	"testing"

	mgo "github.com/globalsign/mgo"
	"github.com/stretchr/testify/assert"
)

const invalidMongoArgs = "invalid"

func TestDB_InitMongo(t *testing.T) {
	var sess *mgo.Session

	f := func() {
		InitMongo(invalidMongoArgs)
	}
	assert.Panics(t, f)

	validMongoArgs := os.Getenv("MONGO_URI")

	f = func() {
		InitMongo(validMongoArgs)
	}
	assert.NotPanics(t, f)

	newSession := InitMongo(validMongoArgs)
	assert.IsType(t, sess, newSession)

}

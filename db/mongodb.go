package db

import (
	"log"

	mgo "github.com/globalsign/mgo"
)

// InitMongo starts a mongo session
func InitMongo(arg string) *mgo.Session {
	session, err := mgo.Dial(arg)
	if err != nil {
		log.Panicf("%s\n", err.Error())
	}
	return session
}

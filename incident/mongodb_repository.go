package incident

import (
	"time"

	mgo "github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/involvestecnologia/statuspage/errors"
	"github.com/involvestecnologia/statuspage/models"
)

const databaseName = "status"

type mongoRepository struct {
	db *mgo.Session
}

func NewMongoRepository(session *mgo.Session) Repository {
	return &mongoRepository{db: session}
}

func (r *mongoRepository) Insert(incident models.Incident) (err error) {
	defer mongoFailure(&err)
	err = r.db.DB(databaseName).C("Incidents").Insert(incident)
	return err
}

func (r *mongoRepository) Find(queryParam map[string]interface{}) (incidents []models.Incident, err error) {
	defer mongoFailure(&err)
	err = r.db.DB(databaseName).C("Incidents").Find(queryParam).All(&incidents)
	if incidents == nil {
		return incidents, &errors.ErrNotFound{Message: errors.ErrNotFoundMessage}
	}
	return incidents, err
}

func (r *mongoRepository) FindOne(queryParam map[string]interface{}) (incident models.Incident, err error) {
	defer mongoFailure(&err)
	err = r.db.DB(databaseName).C("Incidents").Find(queryParam).One(&incident)
	if err != nil {
		return incident, &errors.ErrNotFound{Message: errors.ErrNotFoundMessage}
	}
	return incident, err
}

func (r *mongoRepository) Update(incident models.Incident) (err error) {
	defer mongoFailure(&err)
	err = r.db.DB(databaseName).C("Incidents").Update(bson.M{"component_ref": incident.ComponentRef}, incident)
	if err != nil {
		return &errors.ErrNotFound{Message: errors.ErrNotFoundMessage}
	}
	return err
}

func (r *mongoRepository) List(startDt time.Time, endDt time.Time) (incidents []models.Incident, err error) {
	defer mongoFailure(&err)
	findQ := bson.M{
		"date": bson.M{
			"$gt": startDt.Add(-(24 * time.Hour)),
			"$lt": endDt.Add(24 * time.Hour),
		},
	}

	err = r.db.DB(databaseName).C("Incidents").Find(findQ).All(&incidents)
	return incidents, err
}

func mongoFailure(e *error) {
	if r := recover(); r != nil {
		*e = &errors.ErrMongoFailuere{Message: errors.ErrMongoFailuereMessage}
	}
}

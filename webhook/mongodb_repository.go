package webhook

import (
	"log"

	mgo "github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/involvestecnologia/statuspage/errors"
	"github.com/involvestecnologia/statuspage/models"
)

type mongoRepository struct {
	db *mgo.Session
}

//NewMongoRepository creates a new Repository implementation using the MongoDB as database
func NewMongoRepository(session *mgo.Session) Repository {
	return &mongoRepository{db: session}
}

func (r *mongoRepository) GetCurrentSession() *mgo.Session {
	return r.db
}

func (r *mongoRepository) FindWebhook(id string) (models.Webhook, error) {
	var wh models.Webhook
	if !bson.IsObjectIdHex(id) {
		return wh, &errors.ErrAlreadyExists{Message: errors.ErrAlreadyExistsMessage}
	}
	webhookQ := bson.M{"_id": bson.ObjectIdHex(id)}
	err := r.db.DB("status").C("Webhooks").Find(webhookQ).One(&wh)
	if err != nil && err != mgo.ErrNotFound {
		log.Panicf("Error searching webhook: %s\n", err)
	}
	return wh, err
}

func (r *mongoRepository) FindWebhookByNameAndType(name string, webhookType string) (models.Webhook, error) {
	var wh models.Webhook
	webhookQ := bson.M{"name": name, "type": webhookType}
	err := r.db.DB("status").C("Webhooks").Find(webhookQ).One(&wh)
	if err != nil && err != mgo.ErrNotFound {
		log.Panicf("Error searching webhook: %s\n", err)
	}
	return wh, err
}

func (r *mongoRepository) ListWebhookByType(webhookType string) ([]models.Webhook, error) {
	var whs []models.Webhook
	webhookQ := bson.M{"type": webhookType}
	err := r.db.DB("status").C("Webhooks").Find(webhookQ).All(&whs)
	if err != nil && err != mgo.ErrNotFound {
		log.Panicf("Error searching webhooks: %s\n", err)
	}
	return whs, err
}

func (r *mongoRepository) CreateWebhook(webhook models.Webhook) error {
	return r.db.DB("status").C("Webhooks").Insert(webhook)
}

func (r *mongoRepository) UpdateWebhook(id string, webhook models.Webhook) error {
	if !bson.IsObjectIdHex(id) {
		return &errors.ErrAlreadyExists{Message: errors.ErrAlreadyExistsMessage}
	}
	webhookQ := bson.M{"_id": bson.ObjectIdHex(id)}
	return r.db.DB("status").C("Webhooks").Update(webhookQ, webhook)
}

func (r *mongoRepository) DeleteWebhook(id string) error {
	if !bson.IsObjectIdHex(id) {
		return &errors.ErrAlreadyExists{Message: errors.ErrAlreadyExistsMessage}
	}
	webhookQ := bson.M{"_id": bson.ObjectIdHex(id)}
	return r.db.DB("status").C("Webhooks").Remove(webhookQ)
}

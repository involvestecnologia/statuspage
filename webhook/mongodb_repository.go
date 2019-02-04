package webhook

import (
	"log"

	mgo "github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/involvestecnologia/statuspage/errors"
	"github.com/involvestecnologia/statuspage/models"
)

type MongoRepository struct {
	db *mgo.Session
}

func NewMongoRepository(session *mgo.Session) *MongoRepository {
	return &MongoRepository{db: session}
}

func (r *MongoRepository) GetCurrentSession() *mgo.Session {
	return r.db
}

func (r *MongoRepository) FindWebhook(id string) (models.Webhook, error) {
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

func (r *MongoRepository) FindWebhookByNameAndType(name string, webhookType string) (models.Webhook, error) {
	var wh models.Webhook
	webhookQ := bson.M{"name": name, "type": webhookType}
	err := r.db.DB("status").C("Webhooks").Find(webhookQ).One(&wh)
	if err != nil && err != mgo.ErrNotFound {
		log.Panicf("Error searching webhook: %s\n", err)
	}
	return wh, err
}

func (r *MongoRepository) ListWebhookByType(webhookType string) ([]models.Webhook, error) {
	var whs []models.Webhook
	webhookQ := bson.M{"type": webhookType}
	err := r.db.DB("status").C("Webhooks").Find(webhookQ).All(&whs)
	if err != nil && err != mgo.ErrNotFound {
		log.Panicf("Error searching webhooks: %s\n", err)
	}
	return whs, err
}

func (r *MongoRepository) CreateWebhook(webhook models.Webhook) error {
	return r.db.DB("status").C("Webhooks").Insert(webhook)
}

func (r *MongoRepository) UpdateWebhook(id string, webhook models.Webhook) error {
	if !bson.IsObjectIdHex(id) {
		return &errors.ErrAlreadyExists{Message: errors.ErrAlreadyExistsMessage}
	}
	webhookQ := bson.M{"_id": bson.ObjectIdHex(id)}
	return r.db.DB("status").C("Webhooks").Update(webhookQ, webhook)
}

func (r *MongoRepository) DeleteWebhook(id string) error {
	if !bson.IsObjectIdHex(id) {
		return &errors.ErrAlreadyExists{Message: errors.ErrAlreadyExistsMessage}
	}
	webhookQ := bson.M{"_id": bson.ObjectIdHex(id)}
	return r.db.DB("status").C("Webhooks").Remove(webhookQ)
}

package webhook

import (
	"fmt"

	"github.com/involvestecnologia/statuspage/errors"
	"github.com/involvestecnologia/statuspage/models"
)

type webhookService struct {
	repo Repository
}

//NewService creates implementation of the Service interface
func NewService(r Repository) Service {
	return &webhookService{repo: r}
}

type trigger func() error

var availableTriggers = map[string]trigger{
	"component.creation": componentCreationTrigger,
	"component.update":   componentUpdateTrigger,
	"component.deletion": componenDeletionTrigger,
	"incident.creation":  incidentCreationTrigger,
	"client.creation":    clientCreationTrigger,
	"client.update":      clientUpdateTrigger,
	"client.deletion":    clientDeletionTrigger,
}

func (s *webhookService) TriggerWebhook(element interface{}, action string) error {
	trigger := availableTriggers[action]
	if trigger == nil {
		return &errors.ErrTriggerUnavailable{Message: errors.ErrTriggerUnavailableMessage}
	}
	err := trigger()
	if err != nil {
		return err
	}
	return nil
}

func componentCreationTrigger() error {
	fmt.Println("A component was created ")
	return nil
}

func componentUpdateTrigger() error {
	fmt.Println("A component was modified ")
	return nil
}

func componenDeletionTrigger() error {
	fmt.Println("A component was removed ")
	return nil
}

func incidentCreationTrigger() error {
	fmt.Println("A incident was created ")
	return nil
}

func clientCreationTrigger() error {
	fmt.Println("A client was created ")
	return nil
}

func clientUpdateTrigger() error {
	fmt.Println("A client was modified ")
	return nil
}

func clientDeletionTrigger() error {
	fmt.Println("A client was removed ")
	return nil
}

func (s *webhookService) WebhookExists(webhook models.Webhook) bool {
	_, err := s.repo.FindWebhookByNameAndType(webhook.Name, webhook.Type)
	return err == nil
}

func (s *webhookService) Create(webhook models.Webhook) error {
	return s.repo.CreateWebhook(webhook)
}

func (s *webhookService) List(webhookType string) ([]models.Webhook, error) {
	return s.repo.ListWebhookByType(webhookType)
}

func (s *webhookService) Update(id string, webhook models.Webhook) error {
	return s.repo.UpdateWebhook(id, webhook)
}

func (s *webhookService) Delete(id string) error {
	return s.repo.DeleteWebhook(id)
}

func (s *webhookService) FindWebhook(id string) (models.Webhook, error) {
	return s.repo.FindWebhook(id)
}

func (s *webhookService) ListWebhookByType(webhookType string) ([]models.Webhook, error) {
	return s.repo.ListWebhookByType(webhookType)
}

func (s *webhookService) FindWebhookByNameAndType(name string, webhookType string) (models.Webhook, error) {
	return s.repo.FindWebhookByNameAndType(name, webhookType)
}

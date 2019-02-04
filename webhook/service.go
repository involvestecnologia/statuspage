package webhook

import (
	"fmt"

	"github.com/involvestecnologia/statuspage/errors"
	"github.com/involvestecnologia/statuspage/models"
)

type WebhookService struct {
	repo Repository
}

func NewService(r Repository) *WebhookService {
	return &WebhookService{repo: r}
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

func (s *WebhookService) TriggerWebhook(element interface{}, action string) error {
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

func (s *WebhookService) WebhookExists(webhook models.Webhook) bool {
	_, err := s.repo.FindWebhookByNameAndType(webhook.Name, webhook.Type)
	return err == nil
}

func (s *WebhookService) Create(webhook models.Webhook) error {
	return s.repo.CreateWebhook(webhook)
}

func (s *WebhookService) List(webhookType string) ([]models.Webhook, error) {
	return s.repo.ListWebhookByType(webhookType)
}

func (s *WebhookService) Update(id string, webhook models.Webhook) error {
	return s.repo.UpdateWebhook(id, webhook)
}

func (s *WebhookService) Delete(id string) error {
	return s.repo.DeleteWebhook(id)
}

func (s *WebhookService) FindWebhook(id string) (models.Webhook, error) {
	return s.repo.FindWebhook(id)
}

func (s *WebhookService) ListWebhookByType(webhookType string) ([]models.Webhook, error) {
	return s.repo.ListWebhookByType(webhookType)
}

func (s *WebhookService) FindWebhookByNameAndType(name string, webhookType string) (models.Webhook, error) {
	return s.repo.FindWebhookByNameAndType(name, webhookType)
}

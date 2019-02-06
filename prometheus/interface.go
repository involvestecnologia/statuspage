package prometheus

import "github.com/involvestecnologia/statuspage/models"

// Service describes the use case
type Service interface {
	ProcessIncomingWebhook(models.PrometheusIncomingWebhook) error
}

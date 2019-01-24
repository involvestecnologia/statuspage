package mock

import (
	"github.com/involvestecnologia/statuspage/models"
	"time"
)

func PrometheusModel() map[string]models.PrometheusIncomingWebhook {
	return map[string]models.PrometheusIncomingWebhook{
		"ModelComplete": models.PrometheusIncomingWebhook{
			Alerts: []models.PrometheusAlerts{
				models.PrometheusAlerts{
					Status: "RESOLVED",
					Incident: models.Incident{
						ComponentRef: "123123",
						Description:  "status ok",
						Status:       0,
					},
					Component: models.Component{
						Ref:     "123123",
						Name:    "CompleteModel",
						Address: "",
					},
					StartsAt:     time.Now(),
					EndsAt:       time.Now(),
					GeneratorURL: "ur.com",
				},
			},
		},
		"ModelWithoutName": models.PrometheusIncomingWebhook{
			Alerts: []models.PrometheusAlerts{
				models.PrometheusAlerts{
					Status: "RESOLVED",
					Incident: models.Incident{
						ComponentRef: ZeroTimeHex,
						Description:  "status ok",
						Status:       0,
					},
					Component: models.Component{
						Ref:     ZeroTimeHex,
						Address: "",
					},
					StartsAt:     time.Now(),
					EndsAt:       time.Now(),
					GeneratorURL: "ur.com",
				},
			},
		},
		"ModelWithoutRef": models.PrometheusIncomingWebhook{
			Alerts: []models.PrometheusAlerts{
				models.PrometheusAlerts{
					Status: "RESOLVED",
					Component: models.Component{
						Name:    "newComponent",
						Address: "",
					},
					StartsAt:     time.Now(),
					EndsAt:       time.Now(),
					GeneratorURL: "ur.com",
				},
			},
		},
		"ModelWithoutComponent": models.PrometheusIncomingWebhook{
			Alerts: []models.PrometheusAlerts{
				models.PrometheusAlerts{
					Status: "RESOLVED",
					Incident: models.Incident{
						ComponentRef: ZeroTimeHex,
						Description:  "status ok",
						Status:       0,
					},
					StartsAt:     time.Now(),
					EndsAt:       time.Now(),
					GeneratorURL: "ur.com",
				},
			},
		},
		"ModelComponentNameAlreadyExists": models.PrometheusIncomingWebhook{
			Alerts: []models.PrometheusAlerts{
				models.PrometheusAlerts{
					Status: "RESOLVED",
					Component:	models.Component{
						Ref:     "ZeroTimeHex",
						Name:    "first",
						Address: "",
					},
					StartsAt:     time.Now(),
					EndsAt:       time.Now(),
					GeneratorURL: "ur.com",
				},
			},
		},
		"ModelComponentRefAlreadyExists": models.PrometheusIncomingWebhook{
			Alerts: []models.PrometheusAlerts{
				models.PrometheusAlerts{
					Status: "RESOLVED",
					Component:	models.Component{
						Ref:     ZeroTimeHex,
						Name:    "RefTest",
						Address: "",
					},
					StartsAt:     time.Now(),
					EndsAt:       time.Now(),
					GeneratorURL: "ur.com",
				},
			},
		},
		"ModelBlank": models.PrometheusIncomingWebhook{},
	}
}

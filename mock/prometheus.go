package mock

import (
	"time"

	"github.com/involvestecnologia/statuspage/models"
)

//PrometheusModel return a map of PrometheusIncomingWebhook structures
// to be used in tests named by it's the test case model
func PrometheusModel() map[string]models.PrometheusIncomingWebhook {
	return map[string]models.PrometheusIncomingWebhook{
		"ModelComplete": {
			Alerts: []models.PrometheusAlerts{
				{
					Status: "RESOLVED",
					PrometheusLabel: models.PrometheusLabel{
						ComponentRef: "123123",
						Description:  "status ok",
						Status:       "1",
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
		"ModelFiring": {
			Alerts: []models.PrometheusAlerts{
				{
					Status: "firing",
					PrometheusLabel: models.PrometheusLabel{
						ComponentRef: "123123",
						Description:  "status ok",
						Status:       "3",
					},
					Component: models.Component{
						Ref:  "123123",
						Name: "CompleteModel",
						Labels: []string{
							"Firing",
						},
						Address: "",
					},
					StartsAt:     time.Now(),
					EndsAt:       time.Now(),
					GeneratorURL: "ur.com",
				},
			},
		},
		"ModelUnstable": {
			Alerts: []models.PrometheusAlerts{
				{
					Status: "firing",
					PrometheusLabel: models.PrometheusLabel{
						ComponentRef: "123123",
						Description:  "status unstable",
						Status:       "2",
					},
					Component: models.Component{
						Ref:  "123123",
						Name: "CompleteModel",
						Labels: []string{
							"Firing",
						},
						Address: "",
					},
					StartsAt:     time.Now(),
					EndsAt:       time.Now(),
					GeneratorURL: "ur.com",
				},
			},
		},
		"ModelResolved": {
			Alerts: []models.PrometheusAlerts{
				{
					Status: "resolved",
					PrometheusLabel: models.PrometheusLabel{
						ComponentRef: "123123",
						Description:  "status ok",
						Status:       "1",
					},
					Component: models.Component{
						Ref:  "123123",
						Name: "CompleteModel",
						Labels: []string{
							"Resolved",
						},
						Address: "",
					},
					StartsAt:     time.Now(),
					EndsAt:       time.Now(),
					GeneratorURL: "ur.com",
				},
			},
		},
		"ModelWithoutName": {
			Alerts: []models.PrometheusAlerts{
				{
					Status: "RESOLVED",
					PrometheusLabel: models.PrometheusLabel{
						ComponentRef: ZeroTimeHex,
						Description:  "status ok",
						Status:       "1",
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
		"ModelWithoutRef": {
			Alerts: []models.PrometheusAlerts{
				{
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
		"ModelWithoutComponent": {
			Alerts: []models.PrometheusAlerts{
				{
					Status: "RESOLVED",
					PrometheusLabel: models.PrometheusLabel{
						ComponentRef: ZeroTimeHex,
						Description:  "status ok",
						Status:       "1",
					},
					StartsAt:     time.Now(),
					EndsAt:       time.Now(),
					GeneratorURL: "ur.com",
				},
			},
		},
		"ModelComponentNameAlreadyExists": {
			Alerts: []models.PrometheusAlerts{
				{
					Status: "RESOLVED",
					Component: models.Component{
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
		"ModelWithIncidentWithoutComponentRef": {
			Alerts: []models.PrometheusAlerts{
				{
					Status: "firing",
					PrometheusLabel: models.PrometheusLabel{
						Description: "status nonexistent",
						Status:      "-1",
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
		"ModelComponentRefAlreadyExists": {
			Alerts: []models.PrometheusAlerts{
				{
					Status: "RESOLVED",
					Component: models.Component{
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
		"ModelCompleteWithBadStatus": {
			Alerts: []models.PrometheusAlerts{
				{
					Status: "RESOLVED",
					PrometheusLabel: models.PrometheusLabel{
						ComponentRef: "123123",
						Description:  "status ok",
						Status:       "1",
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
		"ModelBlank": {},
	}
}

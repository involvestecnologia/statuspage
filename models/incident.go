package models

import (
	"time"
)

// Incident represents some event categorized by one of three possible status
type Incident struct {
	Status       int           `json:"status" binding:"required"`
	Resolved     bool          `json:"resolved"`
	Description  string        `json:"description,omitempty"`
	Date         time.Time     `json:"occurrence_date" binding:"required"`
	Duration     time.Duration `json:"duration"`
	ComponentRef string        `json:"component_ref" bson:"component_ref" binding:"required"`
}

const (
	// IncidentStatusOK means "This status is fully operational."
	IncidentStatusOK = 1
	// IncidentStatusUnstable means "You received reports of a problem."
	IncidentStatusUnstable = 2
	// IncidentStatusOutage means "A severe problem happened"
	IncidentStatusOutage = 3
)

package models

// Component represents a server, database, application or any resource that will be monitored
type Component struct {
	Ref     string   `json:"ref,omitempty"`
	Name    string   `json:"name" binding:"required"`
	Labels  []string `json:"labels,omitempty"`
	Address string   `json:"address"`
}

//ComponentRefs is a convenience structure that identifies an list of Component references
type ComponentRefs struct {
	Refs []string `json:"refs,omitempty"`
}

//ComponentLabels is a convenience structure that identifies an list of Component by labels
type ComponentLabels struct {
	Labels []string `json:"labels,omitempty" binding:"required"`
}

type ComponentLog struct {
	ComponentName string
	Address       string
	LogMessage    string
}

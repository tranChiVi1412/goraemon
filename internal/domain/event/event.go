package event

import "time"

type Event interface {
	Name() string
	OccurredAt() time.Time
	Payload() map[string]interface{}
}

type ProjectCreatedEvent struct {
	ProjectID string
	CreatedAt time.Time
}

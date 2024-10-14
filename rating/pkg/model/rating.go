package model

type RecordID string
type RecordType string

const (
	RecordTypeFilm = RecordType("film")
)

type UserID string

// Defines the value of a rating record.
type RatingValue int

// Defines an individual rating created by a user for some record.
type Rating struct {
	RecordID   RecordID    `json:"recordId"`
	RecordType RecordType  `json:"recordType"`
	UserID     UserID      `json:"userId"`
	Value      RatingValue `json:"value"`
}

// RatingEvent defines an event containing rating information.
type RatingEvent struct {
	UserID     UserID          `json:"userId"`
	RecordID   RecordID        `json:"recordId"`
	RecordType RecordType      `json:"recordType"`
	Value      RatingValue     `json:"value"`
	ProviderID string          `json:"providerId"`
	EventType  RatingEventType `json:"eventType"`
}

// RatingEventType defines the type of a rating event.
type RatingEventType string

// Rating event types.
const (
	RatingEventTypePut    = "put"
	RatingEventTypeDelete = "delete"
)
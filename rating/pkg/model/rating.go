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

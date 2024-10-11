package model

import "github.com/samjove/cinephile/metadata/pkg/model"

type FilmDetails struct {
	Rating   *float64       `json:"rating,omitEmpty"`
	Metadata model.Metadata `json:"metadata"`
}
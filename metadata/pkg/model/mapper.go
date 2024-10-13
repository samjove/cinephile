package model

import "github.com/samjove/cinephile/gen"

// MetadataToProto coverts a Metadata struct into a generated proto counterpart.
func MetadataToProto (m *Metadata) *gen.Metadata {
	return &gen.Metadata{
		Id: m.ID,
		Title: m.Title,
		Description: m.Description,
		Director: m.Director,
	}
}

// MetadataFromProto converts a generated proto counter into a Metadata struct.
func MetadataFromProto (m *gen.Metadata) *Metadata {
	return &Metadata{
		ID: m.Id,
		Title: m.Title,
		Description: m.Description,
		Director: m.Director,
	}
}
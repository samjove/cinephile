package memory

import (
	"context"
	"sync"

	"github.com/samjove/cinephile/metadata/internal/repository"
	model "github.com/samjove/cinephile/metadata/pkg"
)

// Repository defines a memory film metadata repository.
type Repository struct {
	sync.RWMutex
	data map[string]*model.Metadata
}

// New creates a new repository
func New() *Repository {
	return &Repository{data: map[string]*model.Metadata{}}
}

// Get retrieves film metadata by id.
func (r *Repository) Get(_ context.Context, id string) (*model.Metadata, error) {
	r.RLock()
	defer r.RUnlock()
	m, ok := r.data[id]
	if !ok {
		return nil, repository.ErrNotFound
	}
	return m, nil
}

// Put adds metadata for a given film id.
func (r *Repository) Put(_ context.Context, id string, m *model.Metadata) error {
	r.Lock()
	defer r.Unlock()
	r.data[id] = m
	return nil
}
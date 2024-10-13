package grpc

import (
	"context"
	"errors"

	"github.com/samjove/cinephile/film/internal/controller/film"
	"github.com/samjove/cinephile/gen"
	"github.com/samjove/cinephile/metadata/pkg/model"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Handler defines a movie gRPC handler.
type Handler struct {
	gen.UnimplementedFilmServiceServer
	ctrl *film.Controller
}

// New creates a new movie gRPC handler.
func New(ctrl *film.Controller) *Handler {
	return &Handler{ctrl: ctrl}
}

// GetMovieDetails returns moviie details by id.
func (h *Handler) GetMovieDetails(ctx context.Context, req *gen.GetFilmDetailsRequest) (*gen.GetFilmDetailsResponse, error) {
	if req == nil || req.FilmId == "" {
		return nil, status.Errorf(codes.InvalidArgument, "nil req or empty id")
	}
	m, err := h.ctrl.Get(ctx, req.FilmId)
	if err != nil && errors.Is(err, film.ErrNotFound) {
		return nil, status.Errorf(codes.NotFound, err.Error())
	} else if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return &gen.GetFilmDetailsResponse{
		FilmDetails: &gen.FilmDetails{
			Metadata: model.MetadataToProto(&m.Metadata),
			Rating:   *m.Rating,
		},
	}, nil
}
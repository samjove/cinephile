package grpc

import (
	"context"

	"github.com/samjove/cinephile/gen"
	"github.com/samjove/cinephile/internal/grpcutil"
	"github.com/samjove/cinephile/metadata/pkg/model"
	"github.com/samjove/cinephile/pkg/discovery"
)

// Gateway defines a film metadata gRPC gateway.
type Gateway struct {
	registry discovery.Registry
}

// New creates a new gRPC gateway for a film metadata service.
func New(registry discovery.Registry) *Gateway {
	return &Gateway{registry}
}

// Get returns film metadata by a film id.
func (g *Gateway) Get(ctx context.Context, id string) (*model.Metadata, error) {
	conn, err := grpcutil.ServiceConnection(ctx, "metadata", g.registry)
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	client := gen.NewMetadataServiceClient(conn)
	resp, err := client.GetMetadata(ctx, &gen.GetMetadataRequest{FilmId: id})
	if err != nil {
		return nil, err
	}
	return model.MetadataFromProto(resp.Metadata), nil
}
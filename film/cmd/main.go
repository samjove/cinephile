package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/samjove/cinephile/film/internal/controller/film"
	metadatagateway "github.com/samjove/cinephile/film/internal/gateway/metadata/http"
	ratinggateway "github.com/samjove/cinephile/film/internal/gateway/rating/http"
	grpchandler "github.com/samjove/cinephile/film/internal/handler/grpc"
	"github.com/samjove/cinephile/gen"
	"github.com/samjove/cinephile/pkg/discovery"
	"github.com/samjove/cinephile/pkg/discovery/consul"
)

const serviceName = "film"

func main() {
	var port int
	flag.IntVar(&port, "port", 8083, "API handler port")
	flag.Parse()
	log.Printf("Starting the film service on port %d", port)
	registry, err := consul.NewRegistry("localhost:8500")
	if err != nil {
		panic(err)
	}
	ctx := context.Background()
	instanceID := discovery.GenerateInstanceID(serviceName)
	if err := registry.Register(ctx, instanceID, serviceName, fmt.Sprintf("localhost:%d", port)); err != nil {
		panic(err)
	}
	go func() {
		for {
			if err := registry.ReportHealthyState(instanceID, serviceName); err != nil {
				log.Println("Failed to report healthy state: " + err.Error())
			}
			time.Sleep(1 * time.Second)
		}
	}()
	defer registry.Deregister(ctx, instanceID, serviceName)
	metadataGateway := metadatagateway.New(registry)
	ratingGateway := ratinggateway.New(registry)
	ctrl := film.New(ratingGateway, metadataGateway)
	h := grpchandler.New(ctrl)
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	srv := grpc.NewServer()
	reflection.Register(srv)
	gen.RegisterFilmServiceServer(srv, h)
	if err := srv.Serve(lis); err != nil {
		panic(err)
	}
}
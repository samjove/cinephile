package main

import (
	"log"
	"net/http"

	"github.com/samjove/cinephile/film/internal/controller/film"
	metadatagateway "github.com/samjove/cinephile/film/internal/gateway/metadata/http"
	ratinggateway "github.com/samjove/cinephile/film/internal/gateway/rating/http"
	httphandler "github.com/samjove/cinephile/film/internal/handler/http"
)

func main() {
	log.Println("Starting the film service")
	metadataGateway := metadatagateway.New("localhost:8081")
	ratingGateway := ratinggateway.New("localhost:8082")
	ctrl := film.New(ratingGateway, metadataGateway)
	h := httphandler.New(ctrl)
	http.Handle("/film", http.HandlerFunc(h.GetFilmDetails))
	if err := http.ListenAndServe(":8083", nil); err != nil {
		panic(err)
	}
}
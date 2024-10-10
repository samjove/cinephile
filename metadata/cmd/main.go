package main

import (
	"log"
	"net/http"

	"github.com/samjove/cinephile/metadata/internal/controller/metadata"
	handler "github.com/samjove/cinephile/metadata/internal/handler/http"
	"github.com/samjove/cinephile/metadata/internal/repository/memory"
)

func main() {
	log.Println("Starting the film metadata service")
	repo := memory.New()
	c := metadata.New(repo)
	h := handler.New(c)
	http.Handle("/metadata", http.HandlerFunc(h.GetMetadata))
	if err := http.ListenAndServe(":8081", nil); err != nil {
		panic(err)
	}
}
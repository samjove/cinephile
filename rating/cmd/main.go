package main

import (
	"log"
	"net/http"

	"github.com/samjove/cinephile/rating/internal/controller/rating"
	httphandler "github.com/samjove/cinephile/rating/internal/handler/http"
	"github.com/samjove/cinephile/rating/internal/repository/memory"
)

func main() {
	log.Println("Starting the rating service")
	repo := memory.New()
	c := rating.New(repo)
	h := httphandler.New(c)
	http.Handle("/rating", http.HandlerFunc(h.Handle))
	if err := http.ListenAndServe(":8082", nil); err != nil {
		panic(err)
	}
}
package main

import (
	"log"
	"net/http"

	httpapi "public-podcast-episode-room/internal/http"
	"public-podcast-episode-room/internal/seed"
	"public-podcast-episode-room/internal/store"
)

func main() {
	s := store.New(seed.DefaultRecords())
	log.Fatal(http.ListenAndServe(":8080", httpapi.Router(s)))
}

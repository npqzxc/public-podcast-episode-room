package httpapi

import (
	"net/http"

	"public-podcast-episode-room/internal/store"
)

func Router(s *store.Store) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", DashboardPage(s))
	mux.HandleFunc("/records", RecordsPage(s))
	mux.HandleFunc("/create", CreatePage())
	mux.HandleFunc("/api/dashboard", DashboardHandler(s))
	mux.HandleFunc("/api/records", RecordsHandler(s))
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("web/static"))))
	return mux
}

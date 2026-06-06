package httpapi

import (
	"html/template"
	"net/http"

	"public-podcast-episode-room/internal/store"
)

func renderTemplate(w http.ResponseWriter, file string, data any) {
	t := template.Must(template.ParseFiles("web/templates/" + file))
	_ = t.Execute(w, data)
}

func DashboardPage(s *store.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w, "dashboard.html", map[string]any{
			"Title":   "Podcast Episode Room",
			"Records": s.List(),
		})
	}
}

func RecordsPage(s *store.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w, "records.html", map[string]any{
			"Title":   "Podcast Episode Room",
			"Records": s.List(),
		})
	}
}

func CreatePage() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w, "create.html", map[string]any{"Title": "Podcast Episode Room"})
	}
}

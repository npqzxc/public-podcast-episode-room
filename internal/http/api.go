package httpapi

import (
	"encoding/json"
	"net/http"

	"public-podcast-episode-room/internal/store"
)

func DashboardHandler(s *store.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		records := s.List()
		json.NewEncoder(w).Encode(map[string]any{
			"totals": []map[string]any{
				{"label": "总记录", "value": len(records)},
				{"label": "进行中", "value": len(records)},
			},
			"records": records,
		})
	}
}

func RecordsHandler(s *store.Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			var payload map[string]string
			_ = json.NewDecoder(r.Body).Decode(&payload)
			json.NewEncoder(w).Encode(s.Create(payload["title"], payload["owner"], payload["note"]))
			return
		}
		json.NewEncoder(w).Encode(s.List())
	}
}

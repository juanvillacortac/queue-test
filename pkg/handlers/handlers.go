package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func respondWithError(w http.ResponseWriter, code int, err error) {
	respondWithPayload(w, code, map[string]any{
		"error": err.Error(),
	})
}

func respondWithPayload(w http.ResponseWriter, code int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(payload)
}

func RegisterHandlers(r chi.Router, prefix string) {
	r.Route(prefix, func(r chi.Router) {
		r.Route("/auth", authRouter)
		r.Route("/clients", clientsRouter)
		r.Route("/queue", queueRouter)
		r.Route("/history", historyRouter)
	})
}

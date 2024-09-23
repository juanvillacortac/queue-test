package handlers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/juanvillacortac/bank-queue/pkg/auth"
	"github.com/juanvillacortac/bank-queue/pkg/database"
	"github.com/juanvillacortac/bank-queue/pkg/repositories"
)

func historyRouter(r chi.Router) {
	r.Use(auth.AuthMiddleware)

	r.Get("/", getHistoryHandler)
}

func getHistoryHandler(w http.ResponseWriter, r *http.Request) {
	repo := repositories.NewSQLHistoryRepository(database.Instance)
	entries, err := repo.GetHistoryEntries()
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, err)
		return
	}
	respondWithPayload(w, http.StatusOK, entries)
}

package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/juanvillacortac/bank-queue/pkg/auth"
	"github.com/juanvillacortac/bank-queue/pkg/database"
	"github.com/juanvillacortac/bank-queue/pkg/models"
	"github.com/juanvillacortac/bank-queue/pkg/repositories"
)

var ErrorClientNotFound = fmt.Errorf("cliente no encontrado")

type CreateClientRequest struct {
	DPI        string            `json:"dpi"`
	Name       string            `json:"name"`
	ClientType models.ClientType `json:"clientType"`
}

func clientsRouter(r chi.Router) {
	r.Use(auth.AuthMiddleware)

	r.Post("/", createClientHandler)
	r.Get("/", getAllClientsHandler)
}

func createClientHandler(w http.ResponseWriter, r *http.Request) {
	var request CreateClientRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		log.Println("Error decoding JSON: ", err)
	}
	repo := repositories.NewSQLClientRepository(database.Instance)
	client, err := repo.CreateClient(models.Client{
		DPI:        request.DPI,
		Name:       request.Name,
		ClientType: request.ClientType,
	})
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, err)
		return
	}
	respondWithPayload(w, http.StatusOK, client)
}

func getAllClientsHandler(w http.ResponseWriter, r *http.Request) {
	repo := repositories.NewSQLClientRepository(database.Instance)
	clients, err := repo.GetAllClients()
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, err)
		return
	}
	respondWithPayload(w, http.StatusOK, clients)
}

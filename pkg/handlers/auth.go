package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/juanvillacortac/bank-queue/pkg/auth"
	"github.com/juanvillacortac/bank-queue/pkg/database"
	"github.com/juanvillacortac/bank-queue/pkg/repositories"
)

var (
	ErrorUnauthorized = fmt.Errorf("Correo o contraseña inválida")
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func authRouter(r chi.Router) {
	r.Post("/login", loginHandler)

	authRouter := r.With(auth.AuthMiddleware)

	authRouter.HandleFunc("/logout", logoutHandler)
	authRouter.Get("/whoami", whoamiHandler)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	var loginRequest LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&loginRequest); err != nil {
		log.Println("Error decoding JSON: ", err)
	}
	repo := repositories.NewSQLUserRepository(database.Instance)
	user, err := repo.GetUser(loginRequest.Email)
	if err != nil || !user.ValidatePassword(loginRequest.Password) {
		log.Println(err)
		respondWithError(w, http.StatusUnauthorized, ErrorUnauthorized)
		return
	}
	claims := auth.NewUserClaims(*user)
	token, err := auth.GenerateJWT(claims)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, err)
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    token,
		HttpOnly: true,
		Secure:   true,
		Path:     "/",
		Expires:  time.Unix(claims.ExpiresAt, 0),
	})
	w.Header().Set("x-token", token)
	respondWithPayload(w, http.StatusOK, map[string]any{
		"user":  user,
		"token": token,
	})
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		HttpOnly: true,
		Secure:   true,
		Path:     "/",
		Expires:  time.Unix(0, 0),
	})
}

func whoamiHandler(w http.ResponseWriter, r *http.Request) {
	user, _ := auth.GetUserFromContext(r.Context())
	respondWithPayload(w, http.StatusOK, user)
}

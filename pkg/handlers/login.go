package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/IgnacioAgustinCabral/notes-api/pkg/auth"
	"github.com/IgnacioAgustinCabral/notes-api/pkg/db"
	"github.com/IgnacioAgustinCabral/notes-api/pkg/payloads"
	"github.com/jackc/pgx/v5"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	request := payloads.LoginRequest{}

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	var hashedPass string
	err = db.Conn.QueryRow(
		context.Background(),
		`SELECT u.password FROM "user".user u WHERE u.username = ($1)`,
		request.Username,
	).Scan(&hashedPass)

	if errors.Is(err, pgx.ErrNoRows) {
		http.Error(w, "Username not found", http.StatusUnauthorized)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashedPass), []byte(request.Password))
	if err != nil {
		http.Error(w, "Username or password incorrect", http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	jwt, err := auth.GenerateJWT(request.Username)
	if err != nil {
		http.Error(w, "Error generating JWT", http.StatusInternalServerError)
		return
	}
	response := map[string]string{
		"msg": "Login successful",
		"jwt": jwt,
	}
	json.NewEncoder(w).Encode(response)
}

package handlers

import (
	"context"
	"encoding/json"
	"github.com/IgnacioAgustinCabral/notes-api/pkg/db"
	"github.com/IgnacioAgustinCabral/notes-api/pkg/payloads"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func Register(w http.ResponseWriter, r *http.Request) {
	request := payloads.RegisterRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	var hashedPassword []byte
	hashedPassword, err = bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Failed to hash password", http.StatusInternalServerError)
		return
	}

	_, err = db.Conn.Exec(context.Background(),
		`INSERT INTO "user".user (username, email, password) VALUES ($1, $2, $3)`,
		request.Username,
		request.Email,
		string(hashedPassword),
	)
	if err != nil {
		response := map[string]string{"msg": "Error while creating user, username or email already exists"}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	response := map[string]string{"msg": "User created successfully"}
	json.NewEncoder(w).Encode(response)
}

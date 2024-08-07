package handlers

import (
	"context"
	"encoding/json"
	"github.com/IgnacioAgustinCabral/notes-api/pkg/db"
	"github.com/IgnacioAgustinCabral/notes-api/pkg/payloads"
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

	var id int

	err = db.Conn.QueryRow(context.Background(),
		`INSERT INTO "user".user (username, email, password) VALUES ($1, $2, $3) RETURNING id`,
		request.Username,
		request.Email,
		request.Password,
	).Scan(&id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	response := map[string]string{"msg": "User created successfully"}
	json.NewEncoder(w).Encode(response)
}
package note

import (
	"context"
	"encoding/json"
	"github.com/IgnacioAgustinCabral/notes-api/pkg/db"
	"github.com/IgnacioAgustinCabral/notes-api/pkg/middleware"
	"github.com/IgnacioAgustinCabral/notes-api/pkg/payloads/note"
	"net/http"
)

func CreateNote(w http.ResponseWriter, r *http.Request) {
	request := note.NoteRequest{}
	claims := middleware.GetClaims(r)

	json.NewDecoder(r.Body).Decode(&request)

	defer r.Body.Close()

	_, err := db.Conn.Exec(context.Background(),
		`INSERT INTO "note".note (title,content,user_id) VALUES ($1,$2,$3)`,
		request.Title,
		request.Content,
		claims.ID)
	if err != nil {
		response := map[string]string{"msg": "Error while creating note, try again"}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response)
		return
	}
	w.WriteHeader(http.StatusOK)
	resp := map[string]string{"msg": "Note saved successfully"}
	json.NewEncoder(w).Encode(resp)
}

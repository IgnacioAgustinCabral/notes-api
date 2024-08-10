package main

import (
	"fmt"
	"github.com/IgnacioAgustinCabral/notes-api/pkg/db"
	"github.com/IgnacioAgustinCabral/notes-api/pkg/handlers/note"
	"github.com/IgnacioAgustinCabral/notes-api/pkg/handlers/user"
	"github.com/IgnacioAgustinCabral/notes-api/pkg/middleware"
	"net/http"
)

func main() {
	db.Init()
	defer db.Close()

	mux := http.NewServeMux()

	mux.HandleFunc("POST /register", user.Register)
	mux.HandleFunc("POST /login", user.Login)
	mux.Handle("POST /notes", middleware.AuthMiddleware(http.HandlerFunc(note.CreateNote)))

	err := http.ListenAndServe(":9090", mux)
	if err != nil {
		fmt.Println(err.Error())
	}
}

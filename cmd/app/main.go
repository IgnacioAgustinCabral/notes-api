package main

import (
	"fmt"
	"github.com/IgnacioAgustinCabral/notes-api/pkg/db"
	"github.com/IgnacioAgustinCabral/notes-api/pkg/handlers"
	"net/http"
)

func main() {
	db.Init()
	defer db.Close()

	mux := http.NewServeMux()

	mux.HandleFunc("POST /register", handlers.Register)

	err := http.ListenAndServe(":9090", mux)
	if err != nil {
		fmt.Println(err.Error())
	}
}

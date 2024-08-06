package main

import (
	"fmt"
	"github.com/IgnacioAgustinCabral/notes-api/pkg/db"
	"net/http"
)

func main() {
	db.Init()
	defer db.Close()

	mux := http.NewServeMux()

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Println(err.Error())
	}
}

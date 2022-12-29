package main

import (
	"net/http"

	"github.com/al33h/go-test/database"
	"github.com/al33h/go-test/web"
	_ "github.com/lib/pq"
)

func main() {
	database.GetConnection()
	handles()
}

func handles() {
	http.HandleFunc("/hello", web.Consultar)
	http.ListenAndServe(":8080", nil)
}

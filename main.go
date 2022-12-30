package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/al33h/go-test/connection"
	"github.com/al33h/go-test/internal"
	"github.com/al33h/go-test/models"
	_ "github.com/lib/pq"
)

func main() {
	connection.GetConnection()
	handles()
}

func handles() {
	http.HandleFunc("/consultar/new", new)
	http.HandleFunc("/consultar/get", findAll)
	http.HandleFunc("/consultar/get/{id}", findById)
	http.HandleFunc("/consultar/delete/{id}", deleteById)

	fmt.Println("Server is on!")
	http.ListenAndServe(":8080", nil)
}

func new(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(`método Http não aceito!`)
		return
	}

	if r.Method == http.MethodPost {
		consulta := &models.ConsultaRequest{}

		err := json.NewDecoder(r.Body).Decode(consulta)
		if err != nil {
			w.WriteHeader(http.StatusBadGateway)
			json.NewEncoder(w).Encode(err.Error())
			return
		}

		consultaRealizada, err := internal.Calcular(consulta)
		if err != nil {
			w.WriteHeader(http.StatusBadGateway)
			json.NewEncoder(w).Encode(err.Error())
			return
		}

		connection.DB.Create(&consultaRealizada)

		ConsultaResponse := consultaRealizada.ToConsultaResponse()

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(ConsultaResponse)
	}
}

func findAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")

	if r.Method == http.MethodGet {
		var consultaRealizadas []internal.ConsultaRealizada
		connection.DB.Find(&consultaRealizadas)
		json.NewEncoder(w).Encode(consultaRealizadas)
	}

}

func findById(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

func deleteById(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

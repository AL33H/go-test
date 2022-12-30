package consultarealizada

import (
	"encoding/json"
	"net/http"

	"github.com/al33h/go-test/config"
	"github.com/al33h/go-test/domain"
	"github.com/al33h/go-test/models"
)

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

		consultaResponse, err := CalcularTodasInfo(consulta)

		if err != nil {
			w.WriteHeader(http.StatusBadGateway)
			json.NewEncoder(w).Encode(err.Error())
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(consultaResponse)
	}
}

func findAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")

	if r.Method == http.MethodGet {
		var consultaRealizadas []domain.ConsultaRealizada
		config.DB.Find(&consultaRealizadas)
		json.NewEncoder(w).Encode(consultaRealizadas)
	}

}

func findById(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

func deleteById(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

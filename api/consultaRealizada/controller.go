package consultarealizada

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/al33h/go-test/domain"
	"github.com/al33h/go-test/models"
	"github.com/al33h/go-test/repository"
	"github.com/go-playground/validator"
)

func headerContentTypeJSON(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
}

func new(w http.ResponseWriter, r *http.Request) {
	headerContentTypeJSON(w)
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(`método http não aceito!`)
		return
	}

	consulta := &models.ConsultaRequest{}

	err := json.NewDecoder(r.Body).Decode(consulta)
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	err = validator.New().Struct(consulta)

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

func findAll(w http.ResponseWriter, r *http.Request) {
	headerContentTypeJSON(w)

	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(`método http não aceito!`)
		return
	}

	var consultaRealizadas []domain.ConsultaRealizada
	var consultaResponses []models.ConsultaResponse

	consultaRealizadas = repository.GetAll()

	for _, consulta := range consultaRealizadas {
		consultaResponses = append(consultaResponses, consulta.ToConsultaResponse())
	}

	json.NewEncoder(w).Encode(consultaResponses)

}

func findById(w http.ResponseWriter, r *http.Request) {
	headerContentTypeJSON(w)

	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(`método http não aceito!`)
		return
	}

	id, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/consultar/"))
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		json.NewEncoder(w).Encode("Codigo inválido")
		return
	}
	consultaRealizada := repository.GetById(id)
	json.NewEncoder(w).Encode(consultaRealizada.ToConsultaResponse())

}

func deleteById(w http.ResponseWriter, r *http.Request) {
	headerContentTypeJSON(w)

	if r.Method != http.MethodDelete {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(`método http não aceito!`)
		return
	}

	id, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/consultar/delete/"))
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		json.NewEncoder(w).Encode(err.Error())
		return
	}
	repository.DeleteById(id)

}

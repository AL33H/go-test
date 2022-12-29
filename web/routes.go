package web

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/al33h/go-test/database"
	"github.com/al33h/go-test/internal"
	"github.com/al33h/go-test/models"
)

func Consultar(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		consulta := &models.ConsultaRequest{}
		json.NewDecoder(r.Body).Decode(consulta)
		consultaRealizada := internal.Calcular(consulta)

		fmt.Println(consultaRealizada)
		database.DB.Create(&consultaRealizada)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(201)
		json.NewEncoder(w).Encode(consultaRealizada)
	}
}

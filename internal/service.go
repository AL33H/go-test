package internal

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/al33h/go-test/models"
)

func Calcular(consulta *models.ConsultaRequest) ConsultaRealizada {
	consultaRealizada := ConsultaRealizada{}
	consultaRealizada.ToConsultaRealizada(*consulta)
	return consultaRealizada
}

func ConsultaExternalCEP(cep string) models.ConsultaExternalCEP {
	url := "https://viacep.com.br/ws/" + cep + "/json/"

	resp, err := http.Get(url)
	if err != nil || resp.StatusCode != http.StatusOK {
		log.Println("Não foi possível requisitar para o viaCep!")
		log.Println(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Não foi possível requisitar para o viaCep!")
		log.Println(err)
	}

	var consultaExternalCEP models.ConsultaExternalCEP
	json.Unmarshal(body, &consultaExternalCEP)

	return consultaExternalCEP
}

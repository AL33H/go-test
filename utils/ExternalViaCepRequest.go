package utils

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/al33h/go-test/models"
)

// Requisita para API do viacep
func ConsultaExternalCEP(cep string) (models.ConsultaExternalCEP, error) {
	url := "https://viacep.com.br/ws/" + cep + "/json/"

	resp, err := http.Get(url)
	if err != nil || resp.StatusCode != http.StatusOK {
		return models.ConsultaExternalCEP{}, errors.New("erro ao requisitar na viacep o cep: " + cep)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return models.ConsultaExternalCEP{}, errors.New("não foi possivel pegar o payload de retorno do viaCep para o cep " + cep)
	}

	var consultaExternalCEP models.ConsultaExternalCEP
	json.Unmarshal(body, &consultaExternalCEP)

	return consultaExternalCEP, nil
}

package consultarealizada

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/al33h/go-test/domain"
	"github.com/al33h/go-test/models"
)

func init() {
	log.SetFlags(log.Flags() | log.Lshortfile)
	log.SetOutput(os.Stdout)
	log.SetPrefix("ERROR: ")
}

// Calcula VlTotalFrete, DataPrevistaEntrega
func CalcularTodasInfo(consulta *models.ConsultaRequest) (domain.ConsultaRealizada, error) {
	consultaRealizada := domain.ConsultaRealizada{}
	consultaRealizada.ToConsultaRealizada(*consulta)
	err := Calcular(consultaRealizada)
	if err != nil {
		return consultaRealizada, err
	}

	return consultaRealizada, nil
}

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

// Calcula informações para consultaRealizada
func Calcular(consulta domain.ConsultaRealizada) error {
	cepOrigem, err := ConsultaExternalCEP(consulta.CepOrigem)
	if err != nil {
		log.Println(err)
		return err
	}
	cepDestino, err := ConsultaExternalCEP(consulta.CepDestino)
	if err != nil {
		log.Println(err)
		return err
	}

	CalcularValorFrete(cepOrigem, cepDestino, consulta)
	CalcularPrevisaoEntrega(cepOrigem, cepDestino, consulta)

	return nil
}

// Calcula VlTotalFrete
func CalcularValorFrete(origem models.ConsultaExternalCEP, destino models.ConsultaExternalCEP, consulta domain.ConsultaRealizada) {
	desconto := 1.0

	if origem.Uf == destino.Uf {
		desconto = 0.25
	}

	if origem.Ddd == destino.Ddd {
		desconto = 0.5
	}

	consulta.VlTotalFrete = consulta.Peso * desconto

}

// Calcula DataPrevistaEntrega
func CalcularPrevisaoEntrega(origem models.ConsultaExternalCEP, destino models.ConsultaExternalCEP, consulta domain.ConsultaRealizada) {
	prazo := 10

	if origem.Uf == destino.Uf {
		prazo = 1
	}

	if origem.Ddd == destino.Ddd {
		prazo = 3
	}

	consulta.DataPrevistaEntrega = consulta.DataConsulta.AddDate(0, 0, prazo)

}

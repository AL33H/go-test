package internal

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/al33h/go-test/models"
	"gorm.io/gorm"
)

func init() {
	log.SetFlags(log.Flags() | log.Lshortfile)
	log.SetOutput(os.Stdout)
	log.SetPrefix("ERROR: ")
}

// Calcula VlTotalFrete, DataPrevistaEntrega
func Calcular(consulta *models.ConsultaRequest) (ConsultaRealizada, error) {
	consultaRealizada := ConsultaRealizada{}
	consultaRealizada.ToConsultaRealizada(*consulta)
	err := consultaRealizada.Calcular()
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

type ConsultaRealizada struct {
	gorm.Model          `json:"-"`
	NomeDestinatario    string
	CepOrigem           string
	CepDestino          string
	Peso                float64
	VlTotalFrete        float64
	DataPrevistaEntrega time.Time
	DataConsulta        time.Time
}

// converte ConsultaRequest to consultaRealizada
func (c *ConsultaRealizada) ToConsultaRealizada(consulta models.ConsultaRequest) error {
	c.NomeDestinatario = consulta.Nome
	c.CepOrigem = consulta.CepOrigem
	c.CepDestino = consulta.CepDestino
	c.Peso = consulta.Peso
	c.DataConsulta = time.Now()

	return nil
}

// converte	ConsultaRealizada to consultaResponse
func (c *ConsultaRealizada) ToConsultaResponse() models.ConsultaResponse {
	return models.ConsultaResponse{
		NomeDestinatario:    c.NomeDestinatario,
		CepOrigem:           c.CepOrigem,
		CepDestino:          c.CepDestino,
		Peso:                c.Peso,
		VlTotalFrete:        c.VlTotalFrete,
		DataConsulta:        c.DataConsulta.Format(time.RFC822),
		DataPrevistaEntrega: c.DataPrevistaEntrega.Format(time.RFC822),
	}
}

// Calcula informações para consultaRealizada
func (c *ConsultaRealizada) Calcular() error {
	cepOrigem, err := ConsultaExternalCEP(c.CepOrigem)
	if err != nil {
		log.Println(err)
		return err
	}
	cepDestino, err := ConsultaExternalCEP(c.CepDestino)
	if err != nil {
		log.Println(err)
		return err
	}

	c.CalcularValorFrete(cepOrigem, cepDestino)
	c.CalcularPrevisaoEntrega(cepOrigem, cepDestino)

	return nil
}

// Calcula VlTotalFrete
func (c *ConsultaRealizada) CalcularValorFrete(origem models.ConsultaExternalCEP, destino models.ConsultaExternalCEP) {
	desconto := 1.0

	if origem.Uf == destino.Uf {
		desconto = 0.25
	}

	if origem.Ddd == destino.Ddd {
		desconto = 0.5
	}

	c.VlTotalFrete = c.Peso * desconto

}

// Calcula DataPrevistaEntrega
func (c *ConsultaRealizada) CalcularPrevisaoEntrega(origem models.ConsultaExternalCEP, destino models.ConsultaExternalCEP) {
	prazo := 10

	if origem.Uf == destino.Uf {
		prazo = 1
	}

	if origem.Ddd == destino.Ddd {
		prazo = 3
	}

	c.DataPrevistaEntrega = c.DataConsulta.AddDate(0, 0, prazo)

}

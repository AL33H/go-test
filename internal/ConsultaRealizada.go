package internal

import (
	"time"

	"github.com/al33h/go-test/models"
	"gorm.io/gorm"
)

type ConsultaRealizada struct {
	gorm.Model
	NomeDestinatario    string    `json:"nomeDestinatario"`
	CepOrigem           string    `json:"CepOrigem"`
	CepDestino          string    `json:"CEPDestino"`
	Peso                float64   `json:"peso"`
	VlTotalFrete        float64   `json:"vlTotalFrete"`
	DataPrevistaEntrega time.Time `json:"dataPrevistaEntrega"`
	DataConsulta        time.Time `json:"dataConsulta"`
}

func (c *ConsultaRealizada) ToConsultaRealizada(consulta models.ConsultaRequest) {
	c.NomeDestinatario = consulta.Nome
	c.CepOrigem = consulta.CepOrigem
	c.CepDestino = consulta.CepDestino
	c.Peso = consulta.Peso
	c.DataConsulta = time.Now()

	c.Calcular()

}

func (c *ConsultaRealizada) Calcular() {
	cepOrigem := ConsultaExternalCEP(c.CepOrigem)
	cepDestino := ConsultaExternalCEP(c.CepDestino)

	c.CalcularValorFrete(cepOrigem, cepDestino)
	c.CalcularPrevisaoEntrega(cepOrigem, cepDestino)

}

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

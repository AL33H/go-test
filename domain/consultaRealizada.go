package domain

import (
	"time"

	"github.com/al33h/go-test/models"
	"gorm.io/gorm"
)

type ConsultaRealizada struct {
	gorm.Model
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

package models

type ConsultaRequest struct {
	Nome       string  `json:"nome" validate:"required"`
	CepOrigem  string  `json:"cepOrigem" validate:"required"`
	CepDestino string  `json:"cepDestino" validate:"required"`
	Peso       float64 `json:"peso" validate:"required"`
}

type ConsultaResponse struct {
	NomeDestinatario    string  `json:"nome"`
	CepOrigem           string  `json:"cepOrigem"`
	CepDestino          string  `json:"cepDestino"`
	Peso                float64 `json:"peso"`
	VlTotalFrete        float64 `json:"vlTotalFrete"`
	DataPrevistaEntrega string  `json:"dataPrevistaEntrega"`
	DataConsulta        string  `json:"dataConsulta"`
}

type ConsultaExternalCEP struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

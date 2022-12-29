package models

type ConsultaRequest struct {
	Nome       string  `json:"nome"`
	CepOrigem  string  `json:"cepOrigem"`
	CepDestino string  `json:"cepDestino"`
	Peso       float64 `json:"peso"`
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

package consultarealizada

import (
	"log"
	"os"
	"sync"

	"github.com/al33h/go-test/domain"
	"github.com/al33h/go-test/models"
	"github.com/al33h/go-test/repository"
	"github.com/al33h/go-test/utils"
)

func init() {
	log.SetFlags(log.Flags() | log.Lshortfile)
	log.SetOutput(os.Stdout)
	log.SetPrefix("ERROR: ")
}

// Calcula VlTotalFrete, DataPrevistaEntrega
func CalcularTodasInfo(consulta *models.ConsultaRequest) (models.ConsultaResponse, error) {
	consultaRealizada := domain.ConsultaRealizada{}
	consultaRealizada.ToConsultaRealizada(*consulta)
	err := Calcular(&consultaRealizada)

	if err != nil {
		return models.ConsultaResponse{}, err
	}

	repository.Create(consultaRealizada)
	consultaResponse := consultaRealizada.ToConsultaResponse()

	return consultaResponse, nil
}

// Calcula informações para consultaRealizada
func Calcular(consulta *domain.ConsultaRealizada) error {
	cepOrigem, err := utils.ConsultaExternalCEP(consulta.CepOrigem)
	if err != nil {
		log.Println(err)
		return err
	}
	cepDestino, err := utils.ConsultaExternalCEP(consulta.CepDestino)
	if err != nil {
		log.Println(err)
		return err
	}

	var WaitGroup sync.WaitGroup
	WaitGroup.Add(2)

	go func(origem models.ConsultaExternalCEP, destino models.ConsultaExternalCEP, consulta *domain.ConsultaRealizada) {

		prazo := 10

		if origem.Uf == destino.Uf {
			prazo = 1
		}

		if origem.Ddd == destino.Ddd {
			prazo = 3
		}

		consulta.DataPrevistaEntrega = consulta.DataConsulta.AddDate(0, 0, prazo)
		WaitGroup.Done()

	}(cepOrigem, cepDestino, consulta)

	go func(origem models.ConsultaExternalCEP, destino models.ConsultaExternalCEP, consulta *domain.ConsultaRealizada) {
		desconto := 1.0

		if origem.Uf == destino.Uf {
			desconto = 0.25
		}

		if origem.Ddd == destino.Ddd {
			desconto = 0.5
		}

		consulta.VlTotalFrete = consulta.Peso * desconto
		WaitGroup.Done()
	}(cepOrigem, cepDestino, consulta)

	WaitGroup.Wait()

	return nil
}

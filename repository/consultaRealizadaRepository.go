package repository

import (
	"github.com/al33h/go-test/config"
	"github.com/al33h/go-test/domain"
)

func Create(consulta domain.ConsultaRealizada) domain.ConsultaRealizada {
	config.DB.Create(&consulta)
	return consulta
}

func GetAll() []domain.ConsultaRealizada {
	var consultaRealizadas []domain.ConsultaRealizada
	config.DB.Find(&consultaRealizadas)
	return consultaRealizadas
}

func GetById(id int) domain.ConsultaRealizada {
	var consultaRealizadas domain.ConsultaRealizada
	config.DB.First(&consultaRealizadas, id)
	return consultaRealizadas
}

func DeleteById(id int) {
	var consultaRealizadas domain.ConsultaRealizada
	config.DB.First(&consultaRealizadas, id)
	config.DB.Delete(&consultaRealizadas)
}

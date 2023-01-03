package consultarealizada

import (
	"net/http"
)

func RouteConsultaRealizada() {
	http.HandleFunc("/consultar/new", new)
	http.HandleFunc("/consultar/all", findAll)
	http.HandleFunc("/consultar/", findById)
	http.HandleFunc("/consultar/delete/", deleteById)
}

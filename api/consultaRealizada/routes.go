package consultarealizada

import (
	"fmt"
	"net/http"
)

func RouteConsultaRealizada() {
	http.HandleFunc("/consultar/new", new)
	http.HandleFunc("/consultar/get", findAll)
	http.HandleFunc("/consultar/get/{id}", findById)
	http.HandleFunc("/consultar/delete/{id}", deleteById)

	fmt.Println("Server is on!")
	http.ListenAndServe(":8080", nil)
}

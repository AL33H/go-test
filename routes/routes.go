package routes

import (
	"fmt"
	"net/http"

	consultarealizada "github.com/al33h/go-test/api/consultaRealizada"
)

func ConfigRoutes() {
	consultarealizada.RouteConsultaRealizada()
	fmt.Println("Server is on!")
	http.ListenAndServe(":8080", nil)
}

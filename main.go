package main

import (
	"github.com/al33h/go-test/config"
	"github.com/al33h/go-test/routes"
	_ "github.com/lib/pq"
)

func main() {
	config.GetConnection()
	routes.ConfigRoutes()
}

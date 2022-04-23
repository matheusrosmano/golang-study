package main

import (
	"net/http"
	"v1/golang-web/routes"
)

func main() {
	routes.CarregaRotas()
	http.ListenAndServe(":8000", nil)
}
package main

import (
	"awesomeProject/Login/gitalura/routes"
	"net/http"
)

func main() {
	routes.CarregaRotas()
	http.ListenAndServe(":8000", nil)
}

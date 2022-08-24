package main

import (
	"awesomeProject/Login/ProjetoGolang-fiber-postgres/routes"
	"net/http"
)

func main() {
	routes.CarregaRotas()
	http.ListenAndServe(":8000", nil)
}

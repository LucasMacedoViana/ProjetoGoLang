package main

import (
	"awesomeProject/Login/ProjetoGolang-fiber-postgres/routes"
	"net/http"
	"github.com/gofiber/fiber/v2"

)

func main() {
	app := fiber.New()
	app.CarregarRotas()
	routes.CarregaRotas()
	//http.ListenAndServe(":8000", nil)
	app.Listen(":3000")
}

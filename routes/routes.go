package routes

import (
	"awesomeProject/Login/gitalura/controllers"
	"net/http"
	//"github.com/gofiber/fiber/v2"
)

func CarregaRotas() {
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/login", controllers.Login)
	http.HandleFunc("/new", controllers.New)
	http.HandleFunc("/insert", controllers.Insert)
	http.HandleFunc("/delete", controllers.Delete)
	http.HandleFunc("/edit", controllers.Edit)
	http.HandleFunc("/update", controllers.Update)

	//app.Get("/", controllers.Index)
}
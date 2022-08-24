package controllers

import (
	"awesomeProject/Login/ProjetoGolang-fiber-postgres/models"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	todosOsUsuarios := models.BuscaTodosOsUsuarios()
	temp.ExecuteTemplate(w, "Index", todosOsUsuarios)
}

func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nome := r.FormValue("nome")
		email := r.FormValue("email")
		permissao := r.FormValue("permissao")
		squad := r.FormValue("squad")

		squadConvertidaParaInt, err := strconv.Atoi(squad)
		if err != nil {
			log.Println("Erro na conversão da Squad: ", err)
		}

		models.CriaNovoUsuario(nome, email, permissao, squadConvertidaParaInt)
	}
	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	idDoUsuario := r.URL.Query().Get("id")
	models.DeletaUsuario(idDoUsuario)
	http.Redirect(w, r, "/", 301)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	idDoUsuario := r.URL.Query().Get("id")
	usuario := models.EditaUsuario(idDoUsuario)
	temp.ExecuteTemplate(w, "Edit", usuario)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		nome := r.FormValue("nome")
		email := r.FormValue("email")
		permissao := r.FormValue("permissao")
		squad := r.FormValue("squad")

		idConvertidaParaInt, err := strconv.Atoi(id)
		if err != nil {
			log.Println("Erro na convesão do ID para int:", err)
		}

		squadConvertidaParaInt, err := strconv.Atoi(squad)
		if err != nil {
			log.Println("Erro na convesão da squad para int:", err)
		}

		models.AtualizaUsuario(idConvertidaParaInt, nome, email, permissao, squadConvertidaParaInt)

	}
	http.Redirect(w, r, "/", 301)
}
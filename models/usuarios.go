package models

import "awesomeProject/Login/ProjetoGolang-fiber-postgres/db"

type Usuario struct {
	Id         int
	Nome       string
	Email  	   string
	Permissao  string
	Squad      int
}

func BuscaTodosOsUsuarios() []Usuario {
	db := db.ConectaComBancoDeDados()

	selectDeTodosOsUsuarios, err := db.Query("select * from usuario")
	if err != nil {
		panic(err.Error())
	}

	u := Usuario{}
	usuarios := []Usuario{}

	for selectDeTodosOsUsuarios.Next() {
		var id, squad int
		var nome, email, permissao string

		err = selectDeTodosOsUsuarios.Scan(&id, &nome, &email, &permissao, &squad)
		if err != nil {
			panic(err.Error())
		}

		u.Id = id
		u.Nome = nome
		u.Email = email
		u.Permissao = permissao
		u.Squad = squad

		usuarios = append(usuarios, u)
	}
	defer db.Close()
	return usuarios
}
func CriaNovoUsuario(nome, email, permissao string, squad int) {
	db := db.ConectaComBancoDeDados()

	insereDadosNoBanco, err := db.Prepare("insert into usuario(nome, email, permissao, squad) values($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error())
	}

	insereDadosNoBanco.Exec(nome, email, permissao, squad)
	defer db.Close()

}

func DeletaUsuario(id string) {
	db := db.ConectaComBancoDeDados()

	deletarOUsuario, err := db.Prepare("delete from usuario where id=$1")
	if err != nil {
		panic(err.Error())
	}

	deletarOUsuario.Exec(id)
	defer db.Close()

}

func EditaUsuario(id string) Usuario {
	db := db.ConectaComBancoDeDados()

	usuarioDoBanco, err := db.Query("select * from usuario where id=$1", id)
	if err != nil {
		panic(err.Error())
	}

	usuarioParaAtualizar := Usuario{}

	for usuarioDoBanco.Next() {
		var id, squad int
		var nome, email, permissao string

		err = usuarioDoBanco.Scan(&id, &nome, &email, &permissao, &squad)
		if err != nil {
			panic(err.Error())
		}
		usuarioParaAtualizar.Id = id
		usuarioParaAtualizar.Nome = nome
		usuarioParaAtualizar.Email = email
		usuarioParaAtualizar.Permissao = permissao
		usuarioParaAtualizar.Squad = squad
	}
	defer db.Close()
	return usuarioParaAtualizar
}

func AtualizaUsuario(id int, nome, email, permissao string, squad int) {
	db := db.ConectaComBancoDeDados()

	AtualizaUsuario, err := db.Prepare("update usuario set nome=$1, email=$2, permissao=$3, squad=$4 where id=$5")
	if err != nil {
		panic(err.Error())
	}
	AtualizaUsuario.Exec(nome, email, permissao, squad, id)
	defer db.Close()
}
package controllers

import (
	"api/src/database"
	"api/src/model"
	"api/src/repositories"
	"api/src/responses"
	"encoding/json"
	"io/ioutil"
	"errors"
	"net/http"
)

// CriarUsuario insere um usuário no banco de dados
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	corpoRequest, erro := ioutil.ReadAll(r.Body)
	erro = errors.New("Deu Erro")
	if erro != nil {
		responses.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var usuario model.Usuario
	if erro = json.Unmarshal(corpoRequest, &usuario); erro != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := database.Conectar()
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repo := repositories.NovoRepositorioUsuarios(db)
	usuarioID, erro := repo.Criar(usuario)

	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	usuario.ID = usuarioID

	responses.JSON(w, http.StatusCreated, usuario)
}

// BuscarUsuarios Busca todos os usuários salvos no banco de dados
func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando todos os usuários"))
}

// BuscarUsuario busca apenas um usuário salvo no banco de dados
func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando um usuário"))
}

// AtualizarUsuario atualiza apenas um usuário salvo no banco de dados
func AtualizandoUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Atualizando usuário"))
}

// DeletarUsuario deleta apenas um usuário salvo no banco de dados
func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deletando usuário"))
}

	package controllers

	import "net/http"

	// CriarUsuario insere um usuário no banco de dados
	func CriarUsuario(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Criando usuário"))
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

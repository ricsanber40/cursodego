package main

import (
	"banco_sql/manipulador"
	"banco_sql/repo"
	"fmt"
	"net/http"
)

func init() {
	fmt.Println("Vamos começar a subir o servidor...")
	_, err := repo.AbreConexaoComBancoDeDadosSQL()
	if err != nil {
		fmt.Println("Parando a carga do servidor. Erro ao abrir o banco de dados: ", err.Error())
		return
	}
}

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Olá Mundo!")
	})

	http.HandleFunc("/funcao", manipulador.Funcao)
	http.HandleFunc("/ola", manipulador.Ola)
	http.HandleFunc("/local/", manipulador.Local)
	fmt.Println("Estou escutando, comandante...")
	http.ListenAndServe(":8181", nil)
}

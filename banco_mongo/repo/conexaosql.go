package repo

import (
	"github.com/jmoiron/sqlx"
	/*
		github.com/lib/pq não é usado diretamente pela aplicação
	*/
	_ "github.com/lib/pq"
)

//Db armazena a conexão com o banco de dados
var Db *sqlx.DB

//AbreConexaoComBancoDeDadosSQL funcao que abre a conexao com o banco MYSQL
func AbreConexaoComBancoDeDadosSQL() (err error) {
	err = nil
	Db, err = sqlx.Open("postgres", "host=localhost dbname=cursodego user=postgres password=postgres sslmode=disable")
	if err != nil {
		return
	}
	err = Db.Ping()
	if err != nil {
		return
	}
	return
}

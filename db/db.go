package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func ConectarComBancoDeDados() *sql.DB {
	conexao := "user=postgres dbname=Loja password=12345 host=localhost sslmode=disable"
	db, err := sql.Open("postgres", conexao)
	if err != nil {
		panic(err.Error())
	}
	return db

}

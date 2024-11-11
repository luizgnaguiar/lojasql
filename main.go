package main

import (
	//"database/sql"
	"database/sql"
	"html/template"
	"net/http"

	//"github\luizgnaguiar\loja\db"
	//"C:\Users\DEV\go\src\github\luizgnaguiar\loja\db"
	_ "github.com/lib/pq"
)

type Produto struct {
	Nome, Descricao string
	Preco           float64
	Quantidade, id  int
}

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)

}
func ConectarComBancoDeDados() *sql.DB {
	conexao := "user=postgres dbname=Loja password=12345 host=localhost sslmode=disable"
	db, err := sql.Open("postgres", conexao)
	if err != nil {
		panic(err.Error())
	}
	return db

}
func index(w http.ResponseWriter, r *http.Request) {
	db := ConectarComBancoDeDados()
	selectDeTodosProdutos, err := db.Query("select * from produtos")
	if err != nil {
		panic(err.Error())
	}
	p := Produto{}
	produtos := []Produto{}

	for selectDeTodosProdutos.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = selectDeTodosProdutos.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}
		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade

		produtos = append(produtos, p)
	}
	temp.ExecuteTemplate(w, "Index", produtos)
	defer db.Close()
}

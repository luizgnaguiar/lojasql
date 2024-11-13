package models

import "github/luizgnaguiar/lojasql/db"

type Produto struct {
	Nome, Descricao string
	Preco           float64
	Quantidade, Id  int
}

func BuscarTodosOsProdutos() []Produto {
	db := db.ConectarComBancoDeDados()
	selectDeTodosProdutos, err := db.Query("select * from produtos order by id asc")
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
		p.Id = id
		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade

		produtos = append(produtos, p)
	}
	defer db.Close()
	return produtos
}
func CriaNovoProduto(nome, descricao string, preco float64, quantidade int) {
	db := db.ConectarComBancoDeDados()

	insertDadosNoBanco, err := db.Prepare("insert into produtos(nome, descricao, preco, quantidade) values($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error())
	}
	insertDadosNoBanco.Exec(nome, descricao, preco, quantidade)

	defer db.Close()
}
func DeletaOProduto(id string) {
	db := db.ConectarComBancoDeDados()
	deletarOProduto, err := db.Prepare("delete from produtos where id=$1")

	if err != nil {
		panic(err.Error())
	}

	deletarOProduto.Exec(id)
	defer db.Close()
}
func EditaProduto(id string) Produto {
	db := db.ConectarComBancoDeDados()

	produtoDoBanco, err := db.Query("select * from produtos where id=$1", id)

	if err != nil {
		panic(err.Error())
	}
	produtoParaAtualizar := Produto{}

	for produtoDoBanco.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = produtoDoBanco.Scan(&id, &nome, &descricao, &preco, &quantidade)

		if err != nil {
			panic(err.Error())
		}

		produtoParaAtualizar.Id = id
		produtoParaAtualizar.Nome = nome
		produtoParaAtualizar.Descricao = descricao
		produtoParaAtualizar.Preco = preco
		produtoParaAtualizar.Quantidade = quantidade
	}
	defer db.Close()
	return produtoParaAtualizar
}
func AtualizarProduto(id, quantidade int, nome, descricao string, preco float64) {
	db := db.ConectarComBancoDeDados()

	AtualizarProduto, err := db.Prepare("update produtos set nome=$1, descricao=$2, preco=$3, quantidade=$4 where id=$5")
	if err != nil {
		panic(err.Error())
	}
	AtualizarProduto.Exec(nome, descricao, preco, quantidade, id)
	defer db.Close()
}

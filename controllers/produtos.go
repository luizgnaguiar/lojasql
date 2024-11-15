package controllers

import (
	"github/luizgnaguiar/lojasql/models"
	"log"
	"net/http"
	"strconv"
	"text/template"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	todosOsProdutos := models.BuscarTodosOsProdutos()
	temp.ExecuteTemplate(w, "Index", todosOsProdutos)
}
func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "New", nil)
}
func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		precoConvertido, err := strconv.ParseFloat(preco, 64)
		if err != nil {
			log.Println("Erro na conversao do preco", err)
		}

		quantidadeConvertidaParaInt, err := strconv.Atoi(quantidade)
		if err != nil {
			log.Println("Erro na conversao da quantidade", err)
		}
		models.CriaNovoProduto(nome, descricao, precoConvertido, quantidadeConvertidaParaInt)
	}
	http.Redirect(w, r, "/", 301)
}
func Delete(w http.ResponseWriter, r *http.Request) {
	idDoProduto := r.URL.Query().Get("id")
	models.DeletaOProduto(idDoProduto)
	http.Redirect(w, r, "/", 301)
}
func Edit(w http.ResponseWriter, r *http.Request) {
	idDoProduto := r.URL.Query().Get("id")
	produto := models.EditaProduto(idDoProduto)
	temp.ExecuteTemplate(w, "Edit", produto)
}
func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		//O retorno e uma string, e preciso converter

		idConvertidoParaInt, err := strconv.Atoi(id)
		if err != nil {
			log.Println("Erro na conversao do ID para int", err)
		}

		precoConvertidoParaFloat, err := strconv.ParseFloat(preco, 64)
		if err != nil {
			log.Panicln("Erro na conversao do preco para Float64", err)
		}

		quantidadeConvertidaParaInt, err := strconv.Atoi(quantidade)
		if err != nil {
			log.Panicln("Erro na conversao da quantidade para int", err)
		}

		models.AtualizarProduto(idConvertidoParaInt, quantidadeConvertidaParaInt, nome, descricao, precoConvertidoParaFloat)
	}
	http.Redirect(w, r, "/", 301)
}

package controllers

import (
	"log"
	"net/http"
	"strconv"
	"text/template"
	"v1/golang-web/models"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	produtos := models.ConsultaTodosProdutos()
	temp.ExecuteTemplate(w, "Index", produtos)

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
			log.Println("Erro na conversão do preço", err)
		}

		quantidadeConvertida, err := strconv.Atoi(quantidade)

		if err != nil {
			log.Println("Erro na conversão da quantidade", err)
		}

		models.CriarNovoProduto(nome, descricao, precoConvertido, quantidadeConvertida)
	}
	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	idProduto := r.URL.Query().Get("id")
	models.DeletaProduto(idProduto)
	http.Redirect(w, r, "/", 301)
}
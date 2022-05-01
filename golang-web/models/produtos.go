package models

import (
	"log"
	"v1/golang-web/db"
)

type Produto struct {
	Id int
	Nome, Descricao string
	Preco float64
	Quantidade int
}

func ConsultaTodosProdutos() []Produto {
	db := db.ConectaDb()

	selectProd, err := db.Query("select * from produtos")

	if err != nil {
		panic(err.Error())
	}

	p:= Produto{}
	produtos := []Produto{}

	for selectProd.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err := selectProd.Scan(&id, &nome, &descricao, &preco, &quantidade)

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

func CriarNovoProduto(nome, descricao string, preco float64, quantidade int) {
	db := db.ConectaDb()

	insereDados, err := db.Prepare("insert into produtos(nome, descricao, preco, quantidade) values ($1, $2, $3, $4)")

	if err != nil {
		panic(err.Error())
	}

	insereDados.Exec(nome, descricao, preco, quantidade)
	defer db.Close()
}

func DeletaProduto(id string) {
	db := db.ConectaDb()

	deletarOProduto, err := db.Prepare("delete from produtos where id = $1")

	if err != nil {
		panic(err.Error())
	}
	deletarOProduto.Exec(id)
	defer db.Close()
}

func ConsultarProduto(id string) Produto {
	db := db.ConectaDb()

	produtoDoBanco, err := db.Query("select * from produtos where id = $1", id)

	if err != nil {
		panic(err.Error())
	}

	produtoParaAtualizar := Produto{}

	for produtoDoBanco.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err := produtoDoBanco.Scan(&id, &nome, &descricao, &preco, &quantidade)

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

func AtualizaProduto(id int, nome string, descricao string, preco float64, quantidade int) {
	db := db.ConectaDb()

	atualizaProduto, err := db.Prepare("update produtos set nome = $1, descricao = $2, preco = $3, quantidade = $4 where id=$5")
	if err != nil {
		log.Fatal(err.Error())
	}

	atualizaProduto.Exec(nome, descricao, preco, quantidade, id)
	defer db.Close()
}
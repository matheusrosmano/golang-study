package models

import(
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
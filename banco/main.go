package main

import (
	"fmt"
	"v1/banco/contas"
)

func PagarBoleto(conta verificarConta, valor float64) {
	conta.Sacar(valor)
}

type verificarConta interface {
	Sacar(valor float64) string
}

func main() {
	contaDoDenis := contas.ContaPoupanca{}
	contaDoDenis.Depositar(200)
	PagarBoleto(&contaDoDenis, 115)
	fmt.Println(contaDoDenis.ObterSaldo())

	contaDaLuiza := contas.ContaCorrente{}
	contaDaLuiza.Depositar(1000)
	PagarBoleto(&contaDaLuiza, 240)
	fmt.Println(contaDaLuiza.ObterSaldo())
}

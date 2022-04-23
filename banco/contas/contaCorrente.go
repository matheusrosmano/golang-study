package contas

import(
	"v1/banco/clientes"
)

type ContaCorrente struct {
	Titular clientes.Titular
	NumeroAgencia, NumeroConta int
	saldo float64
}

func (c *ContaCorrente) Sacar(valor float64) string {
	podeSacar := c.saldo >= valor && valor >= 0

	if podeSacar {
		c.saldo -= valor
		return "Saque realizado com sucesso"
	} else {
		return "saldo insuficiente"
	}
}

func (c *ContaCorrente) Transferir(valor float64, contaDestino *ContaCorrente) bool {
	if c.saldo >= valor && valor > 0{
		c.saldo -= valor
		contaDestino.Depositar(valor)
		return true
	}
	return false
}

func (c *ContaCorrente) Depositar(valor float64) (string, float64) {
	if valor > 0 {
		c.saldo += valor
		return "Depósito realizado com sucesso.", c.saldo
	}
	return "Favor informar um valor válido.", c.saldo
}

func (c *ContaCorrente) ObterSaldo() float64 {
	return c.saldo
}
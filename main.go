package main

import (
	"fmt"
	"contas"
	"clientes"
)

func PagarBoleto(conta verificarConta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
}

type verificarConta interface {
	Sacar(valor float64) string
}

func main() {
	contaExemplo := contas.ContaCorrente{
		Titular: clientes.Titular{
			Nome: "teste",
			CPF: "123.456.789",
			Profissao: "desenvolvedor"},
		NumeroAgencia: 123,
		NumeroConta: 456}
	contaExemplo.Depositar(100)

	PagarBoleto(&contaExemplo, 60)

	fmt.Println(contaExemplo.ObterSaldo())
	 
}

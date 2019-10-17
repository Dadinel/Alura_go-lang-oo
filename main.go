package main

import (
	"clientes"
	"contas"
	"fmt"
)

func pagarBoleto(conta verificarConta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
}

type verificarConta interface {
	Sacar(valor float64) string
}

func main() {
	clienteBruno := clientes.Titular{"Bruno", "123.123.123.12", "Desenvolvedor"}
	contaDoBruno := contas.ContaCorrente{Titular: clienteBruno, NumeroAgencia: 123, NumeroConta: 123456}

	fmt.Println(contaDoBruno)

	contaExemplo := contas.ContaCorrente{}
	// contaExemplo.Saldo = -100

	fmt.Println(contaExemplo)
	fmt.Println(contaExemplo.Depositar(500))
	fmt.Println(contaExemplo.ObterSaldo())

	contaDoDenis := contas.ContaPoupanca{}
	contaDaPati := contas.ContaCorrente{}

	fmt.Println(contaDoDenis)
	fmt.Println(contaDaPati)

	contaDoDenis.Depositar(100)
	pagarBoleto(&contaDoDenis, 60)
	fmt.Println(contaDoDenis.ObterSaldo())

	contaDaLuisa := contas.ContaCorrente{}
	contaDaLuisa.Depositar(500)
	pagarBoleto(&contaDaLuisa, 400)
	fmt.Println(contaDaLuisa.ObterSaldo())
}

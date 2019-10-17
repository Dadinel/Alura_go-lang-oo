package main

import (
	"fmt"

	"contas"
)

func main() {
	// var contaDoGuilherme ContaCorrente
	// contaDoGuilherme := ContaCorrente{}

	// contaDoGuilherme.titular = "Guilherme"
	// contaDoGuilherme.numeroAgencia = 589
	// contaDoGuilherme.numeroConta = 123456
	// contaDoGuilherme.saldo = 125.50

	// contaDoGuilherme := ContaCorrente{titular: "Guilherme"}

	//Criação das variáveis
	contaDoGuilherme := contas.ContaCorrente{Titular: "Guilherme", NumeroAgencia: 589, NumeroConta: 123456, Saldo: 125.50}
	contaDaBruna := contas.ContaCorrente{"Bruna", 222, 111222, 200}

	fmt.Println(contaDoGuilherme)
	fmt.Println(contaDaBruna)

	var contaDaCris *contas.ContaCorrente
	contaDaCris = new(contas.ContaCorrente)
	contaDaCris.Titular = "Cris"
	contaDaCris.Saldo = 500

	fmt.Println(*contaDaCris)
	fmt.Println(contaDaCris)

	//Método saque
	contaDaSilvia := contas.ContaCorrente{}
	contaDaSilvia.Titular = "Silvia"
	contaDaSilvia.Saldo = 500

	fmt.Println(contaDaSilvia.Sacar(600))
	status, valor := contaDaSilvia.Depositar(250)
	fmt.Println(status, valor)
	fmt.Println(contaDaSilvia.Saldo)

	fmt.Println("\n\n")

	contaDoGustavo := contas.ContaCorrente{Titular: "Gustavo", Saldo: 100}

	statusTransf := contaDaSilvia.Transferencia(200, &contaDoGustavo)
	fmt.Println(statusTransf)
	fmt.Println(contaDaSilvia)
	fmt.Println(contaDoGustavo)
}

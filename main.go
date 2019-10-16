package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo

	if podeSacar {
		c.saldo -= valorDoSaque
		return "Saque realizado com sucesso"
	} else {
		return "Saldo insuficiente"
	}
}

func main() {
	// var contaDoGuilherme ContaCorrente
	// contaDoGuilherme := ContaCorrente{}

	// contaDoGuilherme.titular = "Guilherme"
	// contaDoGuilherme.numeroAgencia = 589
	// contaDoGuilherme.numeroConta = 123456
	// contaDoGuilherme.saldo = 125.50

	// contaDoGuilherme := ContaCorrente{titular: "Guilherme"}

	//Criação das variáveis
	contaDoGuilherme := ContaCorrente{titular: "Guilherme", numeroAgencia: 589, numeroConta: 123456, saldo: 125.50}
	contaDaBruna := ContaCorrente{"Bruna", 222, 111222, 200}

	fmt.Println(contaDoGuilherme)
	fmt.Println(contaDaBruna)

	var contaDaCris *ContaCorrente
	contaDaCris = new(ContaCorrente)
	contaDaCris.titular = "Cris"
	contaDaCris.saldo = 500

	fmt.Println(*contaDaCris)
	fmt.Println(contaDaCris)

	//Método saque
	contaDaSilvia := ContaCorrente{}
	contaDaSilvia.titular = "Silvia"
	contaDaSilvia.saldo = 500

	fmt.Println(contaDaSilvia.Sacar(600))
	fmt.Println(contaDaSilvia.saldo)
}

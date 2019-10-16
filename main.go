package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func main() {
	// var contaDoGuilherme ContaCorrente
	// contaDoGuilherme := ContaCorrente{}

	// contaDoGuilherme.titular = "Guilherme"
	// contaDoGuilherme.numeroAgencia = 589
	// contaDoGuilherme.numeroConta = 123456
	// contaDoGuilherme.saldo = 125.50

	// contaDoGuilherme := ContaCorrente{titular: "Guilherme"}

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
}

package main

import (
	"fmt"
	"my-bank/clientes"
	"my-bank/contas"
)

func main() {
	contaDoDiegues := contas.ContaCorrente{
		Titular: clientes.Titular{
			Nome:      "Diegues",
			CPF:       "11111111111",
			Profissao: "Desenvolvedor",
		},
		NumeroAgencia: 589,
		NumeroConta:   123456,
	}
	contaDoDiegues.Depositar(125.50)

	contaDoFulano := contas.ContaCorrente{
		Titular: clientes.Titular{
			Nome: "Fulano",
		},
		NumeroAgencia: 234,
		NumeroConta:   333444,
	}
	contaDoFulano.Depositar(200.14)

	beltrano := clientes.Titular{
		Nome: "Beltrano",
	}

	var contaDoBeltrano *contas.ContaCorrente = new(contas.ContaCorrente)
	contaDoBeltrano.Titular = beltrano
	contaDoBeltrano.Depositar(500.10)

	var contaDoBeltrano2 *contas.ContaCorrente = new(contas.ContaCorrente)
	contaDoBeltrano2.Titular = beltrano
	contaDoBeltrano2.Depositar(500.10)

	fmt.Println(contaDoDiegues, contaDoFulano)
	fmt.Println(contaDoBeltrano)
	fmt.Println(&contaDoBeltrano, &contaDoBeltrano2) //exibe o endereco da memoria para o ponteito
	fmt.Println(contaDoBeltrano.Titular)
	fmt.Println(contaDoBeltrano == contaDoBeltrano2)   //false, pois sao enderecos de memoria diferentes
	fmt.Println(*contaDoBeltrano == *contaDoBeltrano2) //true, pois esta comparando o conteudo

	fmt.Println(contaDoBeltrano.Sacar(100.10))
	fmt.Println(contaDoBeltrano.ObterSaldo())

	status, valor := contaDoBeltrano.Depositar(500)
	fmt.Println("--> ", status, valor)
	fmt.Println(contaDoBeltrano.ObterSaldo())

	status2 := contaDoDiegues.Transferir(50, contaDoBeltrano)
	fmt.Println(status2)
	fmt.Println(contaDoDiegues.ObterSaldo())
	fmt.Println(contaDoBeltrano.ObterSaldo())

	fmt.Println("---------")

	contaDoDenis := contas.ContaPoupanca{}
	contaDoDenis.Depositar(100.0)
	fmt.Println(contaDoDenis.ObterSaldo())

	PagarBoleto(&contaDoDenis, 60)
	fmt.Println(contaDoDenis.ObterSaldo())

	fmt.Println("--------")
	contaDaLuiza := contas.ContaCorrente{}
	contaDaLuiza.Depositar(250.0)
	fmt.Println(contaDaLuiza.ObterSaldo())
	PagarBoleto(&contaDaLuiza, 220)
	fmt.Println(contaDaLuiza.ObterSaldo())
}

func PagarBoleto(conta contas.IVerificarConta, valorDoBoleto float64) {
	fmt.Println("Pagando boleto!")
	conta.Sacar(valorDoBoleto)
}

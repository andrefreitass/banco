package main

import (
	"banco/boleto"
	"banco/clientes"
	"banco/contas"
	"fmt"
)

func main() {

	clienteAndre := clientes.Titular{Nome: "Andre", Cpf: "01234567890", Profissao: "Analista Sistemas"}
	contaCorrente := contas.ContaCorrente{Titular: clienteAndre, NumeroAgencia: 389, NumeroConta: 10203040}
	contaCorrente.Depositar(1000)
	fmt.Println(contaCorrente.ObterSaldo())
	boleto.PagarBoleto(&contaCorrente, 500)
	fmt.Println(contaCorrente.ObterSaldo())

	clientePop := clientes.Titular{Nome: "Andre Pop", Cpf: "01234567890", Profissao: "Analista Sistemas"}
	contaPop := contas.ContaPoupanca{Titular: clientePop, NumeroAgencia: 389, NumeroConta: 10203040, Operacao: 1}
	contaPop.Depositar(1000)
	fmt.Println(contaPop.ObterSaldo())
	boleto.PagarBoleto(&contaPop, 500)
	fmt.Println(contaPop.ObterSaldo())

}

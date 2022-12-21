package boleto

func PagarBoleto(conta verificarConta, valorBoleto float64) {
	conta.Sacar(valorBoleto)

}

type verificarConta interface {
	Sacar(valor float64) string
}

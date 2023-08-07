package contas

type IVerificarConta interface {
	Sacar(valor float64) string
}

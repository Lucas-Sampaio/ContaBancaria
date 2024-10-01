package domain

type IContaRepository interface {
	Criar(conta *Conta) error
	ObterConta(agencia int, numero int64) (*Conta, error)
	Atualizar(conta *Conta) error
}

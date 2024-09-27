package domain

type IContaRepository interface{
	Criar(conta *Conta) error
}
package domain

import seedwork "github.com/Lucas-Sampaio/ContaBancaria/internal/Domain/SeedWork"

type Conta struct {
	ID              seedwork.ID
	Numero          int64
	Agencia         int
	SaldoDisponivel float64
	Status          StatusConta
}

func CriarConta(agencia int, numeroConta int64) *Conta {
	conta := &Conta{
		ID:      seedwork.NewId(),
		Numero:  numeroConta,
		Agencia: agencia,
		Status:  Ativa,
	}
	return conta
}

package usecase

import (
	domain "github.com/Lucas-Sampaio/ContaBancaria/internal/Domain/ContaAggregate"
	"github.com/Lucas-Sampaio/ContaBancaria/internal/Infra/database"
)

type CriarContaUseCase struct {
	Uow database.UnitOfWork
}

type CriarContaInput struct {
	agencia     int
	numeroConta int64
}

func NewCriarContaUsecase(uow database.UnitOfWork) *CriarContaUseCase {
	return &CriarContaUseCase{
		Uow: uow,
	}
}

func (usecase *CriarContaUseCase) Execute(input CriarContaInput) error {
	conta := domain.CriarConta(input.agencia, input.numeroConta)
	if err := usecase.Uow.ContaRepository().Criar(conta); err != nil {
		return err
	}
	return nil
}

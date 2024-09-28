package usecase

import (
	domain "github.com/Lucas-Sampaio/ContaBancaria/internal/Domain/ContaAggregate"
	"github.com/Lucas-Sampaio/ContaBancaria/internal/Infra/database"
)

type CriarContaUseCase struct {
	uow database.IUnitOfWork
}

type CriarContaInput struct {
	Agencia     int
	NumeroConta int64
}

func NewCriarContaUsecase(uow database.IUnitOfWork) *CriarContaUseCase {
	return &CriarContaUseCase{
		uow: uow,
	}
}

func (usecase *CriarContaUseCase) Execute(input CriarContaInput) error {
	conta := domain.CriarConta(input.Agencia, input.NumeroConta)
	if err := usecase.uow.ContaRepository().Criar(conta); err != nil {
		return err
	}
	return nil
}

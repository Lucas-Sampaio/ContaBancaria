package usecase

import (
	domain "github.com/Lucas-Sampaio/ContaBancaria/internal/Domain/ContaAggregate"
	"github.com/Lucas-Sampaio/ContaBancaria/internal/Infra/database"
)

type ObterContasUseCase struct {
	uow database.IUnitOfWork
}

func NewObterContasUseCase(uow database.IUnitOfWork) *ObterContasUseCase {
	return &ObterContasUseCase{
		uow: uow,
	}
}

func (usecase *ObterContasUseCase) Execute() ([]domain.Conta, error) {

	contaRepo := usecase.uow.ContaRepository()
	return contaRepo.ObterContas()
}

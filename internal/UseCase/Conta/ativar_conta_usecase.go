package usecase

import (
	"github.com/Lucas-Sampaio/ContaBancaria/internal/Infra/database"
	utils "github.com/Lucas-Sampaio/ContaBancaria/internal/Utils"
	"github.com/go-playground/validator/v10"
)

type AtivarContaUseCase struct {
	uow database.IUnitOfWork
}

type AtivarContaInput struct {
	AgenciaNumeroConta string `validate:"required"`
}

func NewAtivarContaUsecase(uow database.IUnitOfWork) *AtivarContaUseCase {
	return &AtivarContaUseCase{
		uow: uow,
	}
}

func (usecase *AtivarContaUseCase) Execute(input AtivarContaInput) error {

	err := input.validate()
	if err != nil {
		return err
	}
	agencia, numero, err := utils.ObterConta(input.AgenciaNumeroConta)
	if err != nil {
		return err
	}

	contaRepo := usecase.uow.ContaRepository()
	conta, err := contaRepo.ObterConta(agencia, numero)

	if err != nil {
		return err
	}

	err = conta.Ativar()
	if err != nil {
		return err
	}

	err = contaRepo.Atualizar(conta)
	return err
}

// Função para validar o input
func (input *AtivarContaInput) validate() error {
	validate := validator.New()
	return validate.Struct(input)
}

package usecase

import (
	"github.com/Lucas-Sampaio/ContaBancaria/internal/Infra/database"
	utils "github.com/Lucas-Sampaio/ContaBancaria/internal/Utils"
	"github.com/go-playground/validator/v10"
)

type DesativarContaUseCase struct {
	uow database.IUnitOfWork
}

type DesativarContaInput struct {
	AgenciaNumeroConta string `validate:"required"`
}

func NewDesativarContaUsecase(uow database.IUnitOfWork) *DesativarContaUseCase {
	return &DesativarContaUseCase{
		uow: uow,
	}
}

func (usecase *DesativarContaUseCase) Execute(input DesativarContaInput) error {

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

	err = conta.Desativar()
	if err != nil {
		return err
	}

	err = contaRepo.Atualizar(conta)
	return err
}

// Função para validar o input
func (input *DesativarContaInput) validate() error {
	validate := validator.New()
	return validate.Struct(input)
}

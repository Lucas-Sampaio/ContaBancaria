package usecase

import (
	"github.com/Lucas-Sampaio/ContaBancaria/internal/Infra/database"
	utils "github.com/Lucas-Sampaio/ContaBancaria/internal/Utils"
	"github.com/go-playground/validator/v10"
)

type BloquearContaUseCase struct {
	uow database.IUnitOfWork
}

type BloquearContaInput struct {
	AgenciaNumeroConta string `validate:"required"`
}

func NewBloquearContaUsecase(uow database.IUnitOfWork) *BloquearContaUseCase {
	return &BloquearContaUseCase{
		uow: uow,
	}
}

func (usecase *BloquearContaUseCase) Execute(input BloquearContaInput) error {

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

	err = conta.Bloquear()
	if err != nil {
		return err
	}

	err = contaRepo.Atualizar(conta)
	return err
}

// Função para validar o input
func (input *BloquearContaInput) validate() error {
	validate := validator.New()
	return validate.Struct(input)
}

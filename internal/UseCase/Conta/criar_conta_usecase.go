package usecase

import (
	"errors"
	"fmt"

	"os"

	domain "github.com/Lucas-Sampaio/ContaBancaria/internal/Domain/ContaAggregate"
	"github.com/Lucas-Sampaio/ContaBancaria/internal/Infra/database"
	"github.com/go-playground/validator/v10"
)

type CriarContaUseCase struct {
	uow database.IUnitOfWork
}

type CriarContaInput struct {
	Agencia     int   `validate:"required,gt=0"`
	NumeroConta int64 `validate:"required,gt=0"`
}

func NewCriarContaUsecase(uow database.IUnitOfWork) *CriarContaUseCase {
	return &CriarContaUseCase{
		uow: uow,
	}
}

func (usecase *CriarContaUseCase) Execute(input CriarContaInput) (*domain.Conta, error) {

	err := input.validate()

	if err != nil {
		var format string = err.Error()
		fmt.Fprintf(os.Stdout, format, []any{}...)
		return nil, err
	}

	contaRepo := usecase.uow.ContaRepository()
	conta, _ := contaRepo.ObterConta(input.Agencia, input.NumeroConta)
	if conta != nil {
		return nil, errors.New("conta ja existente")
	}

	conta = domain.CriarConta(input.Agencia, input.NumeroConta)

	if err := contaRepo.Criar(conta); err != nil {
		return nil, err
	}
	return conta, nil
}

// Função para validar o input
func (input *CriarContaInput) validate() error {
	validate := validator.New()
	return validate.Struct(input)
}

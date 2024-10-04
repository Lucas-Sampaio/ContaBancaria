package repositories

import (
	"errors"

	domain "github.com/Lucas-Sampaio/ContaBancaria/internal/Domain/ContaAggregate"
	"gorm.io/gorm"
)

type ContaRepository struct {
	DB *gorm.DB
}

func NewContaRepository(db *gorm.DB) *ContaRepository {
	return &ContaRepository{DB: db}
}

func (contaRepository *ContaRepository) Criar(conta *domain.Conta) error {
	return contaRepository.DB.Create(conta).Error
}
func (contaRepository *ContaRepository) ObterConta(agencia int, numero int64) (*domain.Conta, error) {

	var conta domain.Conta
	result := contaRepository.DB.Where(&domain.Conta{Agencia: agencia, Numero: numero}).First(&conta)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("conta não encontrada")
	} else if result.Error != nil {
		return nil, result.Error
	} else {
		return &conta, nil
	}
}

func (contaRepository *ContaRepository) Atualizar(conta *domain.Conta) error {
	return contaRepository.DB.Save(conta).Error
}

func (contaRepository *ContaRepository) ObterContas() ([]domain.Conta, error) {

	var contas []domain.Conta
	result := contaRepository.DB.Find(&contas)
	if result.Error != nil {
		return nil, result.Error
	}
	return contas, nil
}

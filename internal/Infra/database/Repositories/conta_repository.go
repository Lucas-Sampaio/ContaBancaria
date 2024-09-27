package repositories

import (
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

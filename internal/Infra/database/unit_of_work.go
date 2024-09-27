package database

import (
	domain "github.com/Lucas-Sampaio/ContaBancaria/internal/Domain/ContaAggregate"
	repositories "github.com/Lucas-Sampaio/ContaBancaria/internal/Infra/database/Repositories"
	"gorm.io/gorm"
)

type unitOfWork struct {
	db          *gorm.DB
	transaction *gorm.DB
	contaRepo   domain.IContaRepository
}

func NewUnitOfWork(db *gorm.DB) UnitOfWork {
	return &unitOfWork{
		db: db,
	}
}

func (u *unitOfWork) Begin() error {
	u.transaction = u.db.Begin()
	// Reinicializa os repositórios usando a transação.
	u.contaRepo = repositories.NewContaRepository(u.transaction)
	return u.transaction.Error
}

func (u *unitOfWork) Commit() error {
	if u.transaction != nil {
		return u.transaction.Commit().Error
	}
	return nil
}

func (u *unitOfWork) Rollback() error {
	if u.transaction != nil {
		return u.transaction.Rollback().Error
	}
	return nil
}

func (u *unitOfWork) ContaRepository() domain.IContaRepository {
	if u.contaRepo == nil {
		u.contaRepo = repositories.NewContaRepository(u.db)
	}
	return u.contaRepo
}

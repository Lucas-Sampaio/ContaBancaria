package database

import (
	domain "github.com/Lucas-Sampaio/ContaBancaria/internal/Domain/ContaAggregate"
	repositories "github.com/Lucas-Sampaio/ContaBancaria/internal/Infra/database/Repositories"
	"gorm.io/gorm"
)

type UnitOfWork struct {
	db          *gorm.DB
	transaction *gorm.DB
	contaRepo   domain.IContaRepository
}

func NewUnitOfWork(db *gorm.DB) *UnitOfWork {
	return &UnitOfWork{
		db: db,
	}
}

func (u *UnitOfWork) Begin() error {
	u.transaction = u.db.Begin()
	u.contaRepo = repositories.NewContaRepository(u.transaction)
	return u.transaction.Error
}

func (u *UnitOfWork) Commit() error {
	if u.transaction != nil {
		return u.transaction.Commit().Error
	}
	return nil
}

func (u *UnitOfWork) Rollback() error {
	if u.transaction != nil {
		return u.transaction.Rollback().Error
	}
	return nil
}

func (u *UnitOfWork) ContaRepository() domain.IContaRepository {
	if u.contaRepo == nil {
		u.contaRepo = repositories.NewContaRepository(u.db)
	}
	return u.contaRepo
}

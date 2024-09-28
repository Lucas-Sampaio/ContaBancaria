package database

import domain "github.com/Lucas-Sampaio/ContaBancaria/internal/Domain/ContaAggregate"

type IUnitOfWork interface {
	Begin() error
	Commit() error
	Rollback() error
	ContaRepository() domain.IContaRepository
}

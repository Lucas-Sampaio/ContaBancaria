package database

import domain "github.com/Lucas-Sampaio/ContaBancaria/internal/Domain/ContaAggregate"

type UnitOfWork interface {
	Begin() error
	Commit() error
	Rollback() error
	ContaRepository() domain.IContaRepository
}

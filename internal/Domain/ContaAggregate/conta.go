package domain

import (
	"errors"
)

type Conta struct {
	ID              int         `gorm:"primaryKey;autoIncrement"`
	Agencia         int         `gorm:"type:int;not null"`
	Numero          int64       `gorm:"type:bigint;not null"`
	SaldoDisponivel float64     `gorm:"type:decimal(10,4);not null"`
	Status          StatusConta `gorm:"type:smallint;not null"`
}

func CriarConta(agencia int, numeroConta int64) *Conta {
	conta := &Conta{
		Numero:  numeroConta,
		Agencia: agencia,
		Status:  Ativa,
	}
	return conta
}

func (c *Conta) Desativar() error {

	if c.Status == Desativada {
		return errors.New("Conta ja encontra-se desativada")
	}

	if c.SaldoDisponivel > 0 {
		return errors.New("Conta nao pode ser desativada porque possui saldo disponivel")
	}

	if c.SaldoDisponivel < 0 {
		return errors.New("Conta nao pode ser desativada porque possui saldo negativado")
	}
	c.Status = Desativada
	return nil
}

func (c *Conta) Bloquear() error {
	if c.Status == Bloqueada {
		return errors.New("Conta ja encontra-se bloqueada")
	}
	if c.Status == Desativada {
		return errors.New("Conta desativada nao pode ser bloqueada")
	}
	c.Status = Bloqueada
	return nil
}

func (c *Conta) Ativar() error {
	c.Status = Ativa
	return nil
}

package domain

type StatusConta int

const (
	Ativa      StatusConta = iota + 1 // 1
	Desativada                        // 2
	Bloqueada                         // 3
)

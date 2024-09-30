package utils

import (
	"errors"
	"strconv"
	"strings"
)

// obtem uma conta atraves de uma string do tipo agencia-conta ex: 1-9872
func ObterConta(agenciaConta string) (int, int64, error) {
	parts := strings.Split(agenciaConta, "-")
	msgErro := "Agencia e conta invalida"

	if len(parts) == 2 {
		// Acessa os valores separados
		agencia, err := strconv.Atoi(parts[0])
		if err != nil {
			return 0, 0, errors.New(msgErro)
		}
		numero, err := strconv.Atoi(parts[1])
		// Converte a string para int
		if err != nil {
			return 0, 0, errors.New(msgErro)
		}
		return agencia, int64(numero), nil
	} else {
		return 0, 0, errors.New(msgErro)
	}
}

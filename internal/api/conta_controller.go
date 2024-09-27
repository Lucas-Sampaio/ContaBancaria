package api

import (
	"net/http"

	"github.com/Lucas-Sampaio/ContaBancaria/internal/Infra/database"
)

type ContaController struct {
	Uow database.UnitOfWork
}

func NewContaController(uow database.UnitOfWork) *ContaController {
	return &ContaController{
		Uow: uow,
	}
}

func (controller *ContaController) Create(w http.ResponseWriter, r *http.Request) {

}

package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/Lucas-Sampaio/ContaBancaria/internal/Infra/database"
	usecase "github.com/Lucas-Sampaio/ContaBancaria/internal/UseCase/Conta"
)

type ContaController struct {
	uow database.IUnitOfWork
}

func NewContaController(uow database.IUnitOfWork) *ContaController {
	return &ContaController{
		uow: uow,
	}
}

func (controller *ContaController) Create(resp http.ResponseWriter, req *http.Request) {
	var dto usecase.CriarContaInput
	err := json.NewDecoder(req.Body).Decode(&dto)
	if err != nil {
		http.Error(resp, err.Error(), http.StatusBadRequest)
		return
	}
	usecase := usecase.NewCriarContaUsecase(controller.uow)
	err = usecase.Execute(dto)
	if err != nil {
		http.Error(resp, err.Error(), http.StatusBadRequest)
		return
	}
	resp.WriteHeader(http.StatusCreated)
}

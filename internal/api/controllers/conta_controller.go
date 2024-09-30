package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/Lucas-Sampaio/ContaBancaria/internal/Infra/database"
	usecase "github.com/Lucas-Sampaio/ContaBancaria/internal/UseCase/Conta"
	errors_api "github.com/Lucas-Sampaio/ContaBancaria/internal/api/errors"
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
		errors_api.SendErrorResponse(resp, http.StatusInternalServerError, "", err)
		return
	}
	usecase := usecase.NewCriarContaUsecase(controller.uow)
	conta, err := usecase.Execute(dto)
	if err != nil {
		errors_api.SendErrorResponse(resp, http.StatusBadRequest, "", err)
		return
	}

	resp.WriteHeader(http.StatusCreated)
	json.NewEncoder(resp).Encode(conta)
}

package controllers

import (
	"encoding/json"
	"errors"
	"net/http"

	errors_api "github.com/Lucas-Sampaio/ContaBancaria/internal/Api/errors"
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

func (controller *ContaController) Criar(resp http.ResponseWriter, req *http.Request) {
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

func (controller *ContaController) Desativar(resp http.ResponseWriter, req *http.Request) {

	agenciaConta := req.PathValue("agenciaConta")
	var dto usecase.DesativarContaInput
	if agenciaConta == "" {
		errors_api.SendErrorResponse(resp, http.StatusNotFound, "", errors.New("agencia e conta nao informada"))
		return
	}
	dto.AgenciaNumeroConta = agenciaConta

	usecase := usecase.NewDesativarContaUsecase(controller.uow)
	err := usecase.Execute(dto)
	if err != nil {
		errors_api.SendErrorResponse(resp, http.StatusBadRequest, "", err)
		return
	}

	resp.WriteHeader(http.StatusNoContent)
}

func (controller *ContaController) Bloquear(resp http.ResponseWriter, req *http.Request) {

	agenciaConta := req.PathValue("agenciaConta")
	var dto usecase.BloquearContaInput
	if agenciaConta == "" {
		errors_api.SendErrorResponse(resp, http.StatusNotFound, "", errors.New("agencia e conta nao informada"))
		return
	}
	dto.AgenciaNumeroConta = agenciaConta

	usecase := usecase.NewBloquearContaUsecase(controller.uow)
	err := usecase.Execute(dto)
	if err != nil {
		errors_api.SendErrorResponse(resp, http.StatusBadRequest, "", err)
		return
	}

	resp.WriteHeader(http.StatusNoContent)
}

func (controller *ContaController) Ativar(resp http.ResponseWriter, req *http.Request) {

	agenciaConta := req.PathValue("agenciaConta")
	var dto usecase.AtivarContaInput
	if agenciaConta == "" {
		errors_api.SendErrorResponse(resp, http.StatusNotFound, "", errors.New("agencia e conta nao informada"))
		return
	}
	dto.AgenciaNumeroConta = agenciaConta

	usecase := usecase.NewAtivarContaUsecase(controller.uow)
	err := usecase.Execute(dto)
	if err != nil {
		errors_api.SendErrorResponse(resp, http.StatusBadRequest, "", err)
		return
	}

	resp.WriteHeader(http.StatusNoContent)
}

func (controller *ContaController) ObterContas(resp http.ResponseWriter, req *http.Request) {
	usecase := usecase.NewObterContasUseCase(controller.uow)
	contas, err := usecase.Execute()
	if err != nil {
		errors_api.SendErrorResponse(resp, http.StatusBadRequest, "", err)
		return
	}

	resp.WriteHeader(http.StatusOK)
	json.NewEncoder(resp).Encode(contas)
}

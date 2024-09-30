package errors_api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
)

// Definição da estrutura de resposta de erro
type ErrorResponse struct {
	Status int         `json:"status"`
	Error  ErrorDetail `json:"error"`
}

type ErrorDetail struct {
	Code    string       `json:"code"`
	Message string       `json:"message"`
	Details []FieldError `json:"details,omitempty"`
}

type FieldError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

// Função auxiliar para enviar resposta de erro
func SendErrorResponse(w http.ResponseWriter, status int, code string, erro error) {
	response := getErrorResponse(status, code, erro)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(response)
}

func getErrorResponse(status int, code string, err error) ErrorResponse {
	response := ErrorResponse{
		Status: status,
		Error: ErrorDetail{
			Code: code,
		},
	}

	// Verifica se o erro é do tipo `validator.ValidationErrors`
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		response.Error.Details = getFieldDetails(validationErrors)
		response.Error.Message = "Existe campos Invalidos"
	} else {
		response.Error.Message = err.Error()
	}
	return response
}

func getFieldDetails(validationErrors validator.ValidationErrors) []FieldError {

	var details []FieldError

	for _, fieldError := range validationErrors {
		var message string
		switch fieldError.Tag() {
		case "required":
			message = fmt.Sprintf("O campo '%s' é obrigatório.\n", fieldError.Field())
		case "min":
			message = fmt.Sprintf("O campo '%s' deve ter pelo menos %s caracteres.\n", fieldError.Field(), fieldError.Param())
		case "max":
			message = fmt.Sprintf("O campo '%s' deve ter no máximo %s caracteres.\n", fieldError.Field(), fieldError.Param())
		case "gt":
			message = fmt.Sprintf("O campo '%s' deve ser maior que %s.\n", fieldError.Field(), fieldError.Param())
		case "gte":
			message = fmt.Sprintf("O campo '%s' deve ser maior ou igual que %s.\n", fieldError.Field(), fieldError.Param())
		default:
			message = fmt.Sprintf("Erro no campo '%s': %s\n", fieldError.Field(), fieldError.Error())
		}

		details = append(details, FieldError{
			Message: message,
			Field:   fieldError.Field(),
		})
	}

	return details
}

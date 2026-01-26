package system

import (
	e "github.com/Ivan-Martins-DevProjects/PayHub/internal/appErrors"
)

// Env Errors
var InternalEnvError = &e.InternalError{
	InternalCode:    "INTERNAL_ENV_ERROR",
	InternalMessage: "Erro interno ao manipular .env",
}

var UpdateEnvError = &e.InternalError{
	InternalCode:    "UPDATE_ENV_ERROR",
	InternalMessage: "Erro ao atualizar arquivo .env",
}

var OpenEnvError = &e.InternalError{
	InternalCode:    "OPEN_ENV_ERROR",
	InternalMessage: "Erro ao abrir arquivo .env",
}

var CreateEnvError = &e.InternalError{
	InternalCode:    "CREATE_ENV_ERROR",
	InternalMessage: "Erro ao criar arquivo .env",
}

var CreateKeysError = &e.InternalError{
	InternalCode:    "CREATE_ENV_ERROR",
	InternalMessage: "Erro ao criar arquivo .keys",
}

var EnvNotFound = &e.InternalError{
	InternalCode:    "ENV_NOT_FOUND",
	InternalMessage: "Arquivo .env não encontrado",
}

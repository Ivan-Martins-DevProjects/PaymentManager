package system

import (
	e "github.com/Ivan-Martins-DevProjects/PayHub/internal/appErrors"
)

var SecretMissing = &e.InternalError{
	InternalCode:    "SECRET_KEY_NOT_FOUND",
	InternalMessage: "SECRET_KEY não encontrada",
}

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
	InternalCode:    "CREATE_KEYS_ERROR",
	InternalMessage: "Erro ao criar arquivo .keys",
}

var EnvNotFound = &e.InternalError{
	InternalCode:    "ENV_NOT_FOUND",
	InternalMessage: "Arquivo .env não encontrado",
}

var DeleteKeysError = &e.InternalError{
	InternalCode:    "DELETE_KEYS_ERROR",
	InternalMessage: "Erro ao criar arquivo .keys",
}

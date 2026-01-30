package repository

import (
	e "github.com/Ivan-Martins-DevProjects/PayHub/internal/appErrors"
)

var GatewayAlreadyExists = &e.InternalError{
	InternalCode:    "GATEWAY_ALREADY_EXISTS",
	InternalMessage: "Gateway já cadastrado",
}

var InternalDBError = &e.InternalError{
	InternalCode:    "POSTGRES_INTERNAL_ERROR",
	InternalMessage: "Erro interno ao registrar informações no banco de dados",
}

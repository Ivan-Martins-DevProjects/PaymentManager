package init

import (
	e "github.com/Ivan-Martins-DevProjects/PayHub/internal/appErrors"
)

var SecretMissing = &e.InternalError{
	InternalCode:    "SECRET_KEY_NOT_FOUND",
	InternalMessage: "SECRET_KEY não encontrada",
}

// /////////////////////////////////////////////////////////////////////

// Init Errors
var InternalInitError = &e.InternalError{
	InternalCode:    "INIT_INTERNAL_ERROR",
	InternalMessage: "Erro interno com a função init",
}

var GatewayNameAlreadyExists = &e.InternalError{
	InternalCode:    "NAME_ALREADY_EXISTS",
	InternalMessage: "Já existe um Gateway com esse nome, por favor altere e tente novamente",
}

// /////////////////////////////////////////////////////////////////////

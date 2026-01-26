package security

import (
	e "github.com/Ivan-Martins-DevProjects/PayHub/internal/appErrors"
)

var EncryptError = &e.InternalError{
	InternalCode:    "ENCRYPT_KEYS_ERROR",
	InternalMessage: "Erro ao criptografar chaves de API",
}

var DecryptError = &e.InternalError{
	InternalCode:    "DECRYPT_KEYS_ERROR",
	InternalMessage: "Erro ao descriptografar chaves de API",
}

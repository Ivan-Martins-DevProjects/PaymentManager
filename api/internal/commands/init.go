package commands

import (
	"fmt"

	security "github.com/Ivan-Martins-DevProjects/PayHub/internal/security"
	system "github.com/Ivan-Martins-DevProjects/PayHub/internal/system"
)

func FuncInit() error {
	filesConfig, err := system.CreateGatewayConfig()
	if err != nil {
		return err
	}

	for _, config := range filesConfig {
		for name, gateway := range config.Gateways {
			key := gateway.Secrets.Api_Key
			secret := gateway.Secrets.Secret_key

			hashedKey, err := security.EncryptKey(key, secret)
			if err != nil {
				return fmt.Errorf("Erro ao encriptografar ApiKey: %v", err)
			}

			found, err := system.NameAlreadyExists(".env", name)
			if err != nil {
				return fmt.Errorf("erro ao validar se Gateway já existe: %v", err)
			}
			if found {
				return fmt.Errorf("Já existe um Gateway com esse nome, por favor altere e tente novamente\nGateway: %v", name)
			}

			err = system.CreateEnv(name, hashedKey)
			if err != nil {
				return fmt.Errorf("Erro ao registrar key criptografada: %s", err)
			}
		}
	}
	return nil
}

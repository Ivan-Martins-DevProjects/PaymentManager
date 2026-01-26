package init

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"

	e "github.com/Ivan-Martins-DevProjects/PayHub/internal/appErrors"
	security "github.com/Ivan-Martins-DevProjects/PayHub/internal/security"
	system "github.com/Ivan-Martins-DevProjects/PayHub/internal/system"
)

var flagSecret string

var PayHubInit = &cobra.Command{
	Use:   "init",
	Short: "Inicialize sua configuração a partir do seus arquivos YAML",
	Run: func(cmd *cobra.Command, args []string) {
		err := funcInit(flagSecret)
		if err != nil {
			fmt.Printf("Erro com a aplicação: %v\n", err)
			return
		}

		fmt.Printf("Criação dos arquivos de configuração concluída\n")
	},
}

func init() {
	PayHubInit.Flags().StringVarP(&flagSecret, "secret", "s", "", "Senha de criptografia das chaves de API")
}

func funcInit(mainSecret string) error {
	filesConfig, err := system.CreateGatewayConfig()
	if err != nil {
		return err
	}

	var secret string
	if mainSecret == "" {
		var createPassKey string
		secret, err = getSecretFromEnv()
		if err != nil {
			var appErr *e.InternalError
			if errors.As(err, &appErr) {
				switch appErr.Code() {
				case "ENV_NOT_FOUND":
					fmt.Print("Arquivo .env não encontrado, digite a SECRET_KEY para criptografar as chaves de API:\n")
					fmt.Scanln(&createPassKey)
					err = createPassKeyEnv(createPassKey)
					if err != nil {
						return err
					}

				case "SECRET_KEY_NOT_FOUND":
					fmt.Print("SECRET_KEY não encontrada, digite a SECRET_KEY para criptografar as chaves de API:\n")
					fmt.Scanln(&createPassKey)
					err = updateSecretEnv(createPassKey)
					if err != nil {
						return err
					}

				default:
					return err
				}
			} else {
				return err
			}
		}
	} else {
		secret = mainSecret
		found, err := NameAlreadyExists(".env", "SECRET_KEY")
		if err != nil {
			return fmt.Errorf("Erro ao validar arquivo .env: %s", err)
		}
		if found {
			var confirm string
			fmt.Print("Foi encontrado uma Secret Key no arquivo .env, deseja sobrescrever[S/N]")
			fmt.Scanln(&confirm)

			if confirm == "N" || confirm == "n" {
				secret, err = getSecretFromEnv()
				if err != nil {
					return err
				}
				return nil
			} else {
				err = updateSecretEnv(secret)
				if err != nil {
					return err
				}
			}
		} else {
			createPassKeyEnv(secret)
		}
	}

	for _, config := range filesConfig {
		for name, gateway := range config.Gateways {
			key := gateway.Secrets.Api_Key

			found, err := NameAlreadyExists(".keys", name)
			if err != nil {
				return err
			}
			if found {
				return e.GenerateError(*GatewayNameAlreadyExists, err)
			}

			hashedKey, err := security.EncryptKey(key, secret)
			if err != nil {
				return err
			}

			err = system.CreateDotKeys(name, hashedKey)
			if err != nil {
				return fmt.Errorf("Erro ao registrar key criptografada: %s", err)
			}
		}
	}
	return nil
}

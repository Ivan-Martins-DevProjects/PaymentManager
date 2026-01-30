package init

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"

	e "github.com/Ivan-Martins-DevProjects/PayHub/internal/appErrors"
	"github.com/Ivan-Martins-DevProjects/PayHub/internal/models"
	"github.com/Ivan-Martins-DevProjects/PayHub/internal/repository"
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
			fmt.Println(err)
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
					err = CreateUpdatePassKeyEnv(createPassKey)
					if err != nil {
						return err
					}

				case "SECRET_KEY_NOT_FOUND":
					fmt.Print("SECRET_KEY não encontrada, digite a SECRET_KEY para criptografar as chaves de API:\n")
					fmt.Scanln(&createPassKey)
					err = CreateUpdatePassKeyEnv(createPassKey)
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
			for {
				var confirm string
				fmt.Print("Foi encontrado uma Secret Key no arquivo .env, deseja sobrescrever[S/N]")
				fmt.Scanln(&confirm)

				switch strings.ToLower(confirm) {
				case "n":
					secret, err = getSecretFromEnv()
					if err != nil {
						return err
					}
					err = insertGatewayInfo(filesConfig, secret)
					if err != nil {
						return err
					}
					return nil

				case "s":
					fmt.Println(secret)
					fmt.Println(filesConfig)
					err := insertGatewayInfo(filesConfig, secret)
					if err != nil {
						return err
					}
					return nil

				default:
					fmt.Print("Opção incorreta, por favor digite 's' para 'Sim' e 'n' para 'Não'\n")
				}
			}
		} else {
			err = insertGatewayInfo(filesConfig, secret)
			if err != nil {
				return err
			}

			return nil
		}
	}

	err = repository.InitRepo(filesConfig, secret)
	if err != nil {
		return err
	}
	return nil
}

func insertGatewayInfo(filesConfig []*models.Config, secret string) error {
	err := SetPassKey(secret)
	if err != nil {
		return err
	}

	err = repository.InitRepo(filesConfig, secret)
	if err != nil {
		return err
	}

	return nil
}

func SetPassKey(secret string) error {
	err := CreateUpdatePassKeyEnv(secret)
	if err != nil {
		return err
	}
	return nil
}

func CreateConfigFile(filesConfig []*models.Config, secret string) error {
	filePath := ".keys"

	err := os.Remove(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("Criando arquivo de configuração!")
		} else {
			return e.GenerateError(*system.DeleteKeysError, err)
		}
	}

	fmt.Println("Estruturando arquivo de configuração")
	for _, config := range filesConfig {
		for name, gateway := range config.Gateways {
			key := gateway.Secrets.Api_Key

			found, err := NameAlreadyExists(filePath, name)
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

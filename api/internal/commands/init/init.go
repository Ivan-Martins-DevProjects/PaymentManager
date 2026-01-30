package init

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	e "github.com/Ivan-Martins-DevProjects/PayHub/internal/appErrors"
	"github.com/Ivan-Martins-DevProjects/PayHub/internal/cache"
	"github.com/Ivan-Martins-DevProjects/PayHub/internal/models"
	"github.com/Ivan-Martins-DevProjects/PayHub/internal/repository"
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

type InsertConfig struct {
	redis    *cache.MainCache
	postgres *repository.MainRepo
	ctx      context.Context
}

func (i *InsertConfig) SaveGatewayInfo(filesConfig []*models.Config, secret string) error {
	err := CreateUpdatePassKeyEnv(secret)
	if err != nil {
		return err
	}

	redisInput, err := i.postgres.InsertGatewayInfo(i.ctx, filesConfig, secret)
	if err != nil {
		return err
	}

	err = i.redis.InsertGatewayCacheInfo(i.ctx, redisInput)
	if err != nil {
		return err
	}

	return nil
}

func funcInit(mainSecret string) error {
	MainCache, err := cache.CreateRedisClient()
	if err != nil {
		return err
	}
	ctx := context.Background()

	p := &repository.PostgresDb{}
	_, err = p.GetPool(ctx)
	if err != nil {
		return err
	}

	repo := &repository.MainRepo{
		DB: p,
	}
	i := &InsertConfig{
		redis:    MainCache,
		postgres: repo,
		ctx:      ctx,
	}

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
					err = i.SaveGatewayInfo(filesConfig, secret)
					if err != nil {
						return err
					}
					return nil

				case "s":
					err := i.SaveGatewayInfo(filesConfig, secret)
					if err != nil {
						return err
					}
					return nil

				default:
					fmt.Print("Opção incorreta, por favor digite 's' para 'Sim' e 'n' para 'Não'\n")
				}
			}
		} else {
			err = i.SaveGatewayInfo(filesConfig, secret)
			if err != nil {
				return err
			}

			return nil
		}
	}

	err = i.SaveGatewayInfo(filesConfig, secret)
	if err != nil {
		return err
	}
	return nil
}

package token

import (
	"fmt"

	tokenjwt "github.com/Ivan-Martins-DevProjects/PayHub/internal/security/tokenJWT"
	"github.com/Ivan-Martins-DevProjects/PayHub/internal/system"
	"github.com/spf13/cobra"
)

var (
	flagExpire string

	PayHubToken = &cobra.Command{
		Use:   "token",
		Short: "Gere seu token de autenticação",
		Run: func(cmd *cobra.Command, args []string) {
			tokens, err := GenerateToken(flagExpire)
			if err != nil {
				fmt.Println(err)
			}

			response := make(map[string]string)

			for _, item := range tokens {
				response[item.ID] = item.Token
			}
			fmt.Printf("Lista de tokens Gerados:\n%v", response)
		},
	}
)

func init() {
	PayHubToken.Flags().StringVarP(&flagExpire, "expire", "e", "", "Tempo de expiração do token")
}

func GenerateToken(expire string) ([]*tokenjwt.ResponseTokens, error) {
	filesConfig, err := system.CreateGatewayConfig()
	if err != nil {
		return nil, err
	}

	var inputs []*tokenjwt.JwtInput

	for _, config := range filesConfig {
		for name, item := range config.Gateways {
			inputs = append(
				inputs,
				&tokenjwt.JwtInput{
					ID:      name,
					ApiURL:  item.Info.Api_URL,
					Timeout: int(item.Retries.Timeout),
					Retries: int(item.Retries.Retries),
				},
			)
		}
	}

	secretInputs, err := tokenjwt.SetSecretJwtInput(inputs)
	if err != nil {
		return nil, err
	}

	tokens, err := tokenjwt.CreateJWT(secretInputs)
	if err != nil {
		return nil, err
	}
	return tokens, nil
}

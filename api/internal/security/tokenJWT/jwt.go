package tokenjwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type ResponseTokens struct {
	ID    string
	Token string
}

func CreateJWT(input []*JwtInput) ([]*ResponseTokens, error) {
	var tokens []*ResponseTokens

	for _, item := range input {
		claims := claims{
			ID:      item.ID,
			ApiURL:  item.ApiURL,
			Timeout: item.Timeout,
			Retries: item.Retries,
			RegisteredClaims: jwt.RegisteredClaims{
				IssuedAt:  jwt.NewNumericDate(time.Now()),
				ExpiresAt: jwt.NewNumericDate(item.Expires),
			},
		}

		rawToken := jwt.NewWithClaims(jwt.SigningMethodHS256, &claims)
		token, err := rawToken.SignedString(item.Secret)
		if err != nil {
			return nil, err
		}

		tokens = append(
			tokens,
			&ResponseTokens{
				ID:    item.ID,
				Token: token,
			},
		)
	}

	return tokens, nil
}

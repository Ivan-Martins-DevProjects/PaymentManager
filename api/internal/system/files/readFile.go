package files

import (
	"github.com/joho/godotenv"
)

func ReadFile(filePath, target string) (string, error) {
	envMap, err := godotenv.Read(filePath)
	if err != nil {
		return "", err
	}

	secret := envMap[target]
	if secret == "" {
		return "", err
	}
	return secret, nil
}

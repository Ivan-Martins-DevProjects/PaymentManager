package init

import (
	"errors"
	"fmt"
	"os"
	"strings"

	e "github.com/Ivan-Martins-DevProjects/PayHub/internal/appErrors"
	s "github.com/Ivan-Martins-DevProjects/PayHub/internal/system"
	"github.com/joho/godotenv"
)

func NameAlreadyExists(Filename, GatewayID string) (bool, error) {
	var file *os.File
	file, err := os.Open(Filename)
	if err != nil {
		if os.IsNotExist(err) {
			file, err = os.Create(Filename)
			if err != nil {
				return false, e.GenerateError(*s.CreateKeysError, err)
			}
		} else {
			return false, e.GenerateError(*s.InternalEnvError, err)
		}
	}
	defer file.Close()

	var content string
	buffer := make([]byte, 1024)

	for {
		n, err := file.Read(buffer)
		if err != nil && err.Error() != "EOF" {
			return false, e.GenerateError(*s.InternalEnvError, err)
		}
		if n == 0 {
			break
		}

		content += string(buffer[:n])

		if strings.Contains(content, GatewayID) {
			return true, nil
		}
	}

	return false, nil
}

func CreateUpdatePassKeyEnv(secretKey string) error {
	fmt.Println("Ajustando .env com as novas informações")
	envMap, err := godotenv.Read(".env")
	if err != nil {
		if os.IsNotExist(err) {
			envMap = make(map[string]string)
		} else {
			return e.GenerateError(*s.OpenEnvError, err)
		}
	}

	envMap["SECRET_KEY"] = secretKey

	err = godotenv.Write(envMap, ".env")
	if err != nil {
		return e.GenerateError(*s.UpdateEnvError, err)
	}

	return nil
}

func getSecretFromEnv() (string, error) {
	var envSecret string
	err := godotenv.Load()
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return "", e.GenerateError(*s.EnvNotFound, err)
		}
		return "", e.GenerateError(*s.InternalEnvError, err)
	}

	envSecret = os.Getenv("SECRET_KEY")
	if envSecret == "" {
		return "", e.GenerateError(*SecretMissing, err)
	}

	return envSecret, nil
}

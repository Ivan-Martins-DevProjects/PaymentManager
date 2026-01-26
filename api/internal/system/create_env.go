package system

import (
	"fmt"
	"os"

	e "github.com/Ivan-Martins-DevProjects/PayHub/internal/appErrors"
)

func CreateDotKeys(GatewayID string, EncriptedKey string) error {
	filename := ".keys"

	var file *os.File

	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		if os.IsNotExist(err) {
			file, err = os.Create(filename)
			if err != nil {
				return e.GenerateError(*InternalEnvError, err)
			}
		}
		return err
	}
	defer file.Close()

	line := fmt.Sprintf("%s=%s\n", GatewayID, EncriptedKey)
	_, err = file.WriteString(line)
	if err != nil {
		return e.GenerateError(*InternalEnvError, err)
	}

	return nil
}

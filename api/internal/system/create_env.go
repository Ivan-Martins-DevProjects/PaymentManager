package system

import (
	"fmt"
	"os"
	"strings"
)

func CreateEnv(GatewayID string, EncriptedKey string) error {
	filename := ".env"

	var file *os.File

	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		if os.IsNotExist(err) {
			file, err = os.Create(filename)
			if err != nil {
				return err
			}
		}
		return err
	}
	defer file.Close()

	line := fmt.Sprintf("%s=%s\n", GatewayID, EncriptedKey)
	_, err = file.WriteString(line)
	if err != nil {
		return err
	}

	return nil
}

func NameAlreadyExists(Filename, GatewayID string) (bool, error) {
	file, err := os.Open(Filename)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		} else {
			return false, err
		}
	}
	defer file.Close()

	var content string
	buffer := make([]byte, 1024)

	for {
		n, err := file.Read(buffer)
		if err != nil && err.Error() != "EOF" {
			return false, err
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

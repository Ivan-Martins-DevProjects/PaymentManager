package system

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"os"
	"path/filepath"

	models "github.com/Ivan-Martins-DevProjects/PayHub/internal/models"
	"gopkg.in/yaml.v3"
)

func CreateGatewayConfig() ([]*models.Config, error) {
	dir := "./gateways"

	files, err := filepath.Glob(filepath.Join(dir, "*.yml"))
	if err != nil {
		return nil, fmt.Errorf("Erro ao listar arquivos YAML: %v", err)
	}

	if len(files) == 0 {
		return nil, fmt.Errorf("Nenhum arquivo encontrado")
	}

	var config *models.Config
	var filesUnpack []*models.Config

	for _, file := range files {
		config, err = ExtractGatewayConfig(file)
		if err != nil {
			return nil, err
		}

		filesUnpack = append(filesUnpack, config)
	}

	return filesUnpack, nil
}

func ExtractGatewayConfig(file string) (*models.Config, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, fmt.Errorf("Erro ao abrir arquivo '%s': %v", file, err)
	}
	defer f.Close()

	var config models.Config
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&config)
	if err != nil {
		return nil, fmt.Errorf("Erro ao decodificar arquivo: %v", err)
	}

	validate := validator.New()
	for key, gateway := range config.Gateways {
		err = validate.Struct(gateway)
		if err != nil {
			validationErrors, ok := err.(validator.ValidationErrors)
			if ok {
				for _, fieldError := range validationErrors {
					return nil, fmt.Errorf("Chave obrigatória ausente: %s", fieldError.Field())
				}
			}
			return nil, fmt.Errorf("Erro de validação com o item: %s", key)
		}
	}

	return &config, nil
}

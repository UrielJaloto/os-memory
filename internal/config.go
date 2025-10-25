package internal

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Config struct {
	MemoriaFisicaSize int `json:"memoria_fisica_size"`
	MemoriaLogicaSize int `json:"memoria_logica_size"`
	PageSize          int `json:"page_size"`
}

func NewConfig(path string) (*Config, error) {
	fileContent, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var config Config
	err = json.Unmarshal(fileContent, &config)
	if err != nil {
		return nil, fmt.Errorf("failed to parse config file: %w", err)
	}
	err = validateConfig(&config)
	if err != nil {
		return nil, err
	}
	return &config, nil

}

func validateConfig(config *Config) error {
	if !isPowerOfTwo(config.MemoriaFisicaSize) {
		return errors.New("MemoriaFisicaSize must be a power of two")
	}
	if !isPowerOfTwo(config.MemoriaLogicaSize) {
		return errors.New("MemoriaLogicaSize must be a power of two")
	}
	if !isPowerOfTwo(config.PageSize) {
		return errors.New("PageSize must be a power of two")
	}

	return nil
}

func isPowerOfTwo(n int) bool {
	if n <= 0 {
		return false
	}
	return (n & (n - 1)) == 0
}

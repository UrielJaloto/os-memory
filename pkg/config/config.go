package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	TamanhoMemoriaFisica  int `json:"tamanho_memoria_fisica"`
	TamanhoPagina         int `json:"tamanho_pagina"`
	TamanhoMaximoProcesso int `json:"tamanho_maximo_processo"`
}

func isPowerOfTwo(n int) bool {
	if n <= 0 {
		return false
	}
	return (n & (n - 1)) == 0
}

func LoadConfig(path string) (*Config, error) {
	var cfg Config

	configFile, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("erro ao abrir arquivo de configuração: %w", err)
	}
	defer configFile.Close()

	decoder := json.NewDecoder(configFile)
	err = decoder.Decode(&cfg)
	if err != nil {
		return nil, fmt.Errorf("erro ao decodificar JSON: %w", err)
	}

	if !isPowerOfTwo(cfg.TamanhoMemoriaFisica) {
		return nil, fmt.Errorf("TamanhoMemoriaFisica (%d) não é uma potência de dois", cfg.TamanhoMemoriaFisica)
	}
	if !isPowerOfTwo(cfg.TamanhoPagina) {
		return nil, fmt.Errorf("TamanhoPagina (%d) não é uma potência de dois", cfg.TamanhoPagina)
	}
	if !isPowerOfTwo(cfg.TamanhoMaximoProcesso) {
		return nil, fmt.Errorf("TamanhoMaximoProcesso (%d) não é uma potência de dois", cfg.TamanhoMaximoProcesso)
	}

	if cfg.TamanhoPagina > cfg.TamanhoMemoriaFisica {
		return nil, fmt.Errorf("TamanhoPagina (%d) não pode ser maior que TamanhoMemoriaFisica (%d)", cfg.TamanhoPagina, cfg.TamanhoMemoriaFisica)
	}

	if cfg.TamanhoMaximoProcesso > cfg.TamanhoMemoriaFisica {
		return nil, fmt.Errorf("TamanhoMaximoProcesso (%d) não pode ser maior que TamanhoMemoriaFisica (%d)", cfg.TamanhoMaximoProcesso, cfg.TamanhoMemoriaFisica)
	}

	return &cfg, nil
}
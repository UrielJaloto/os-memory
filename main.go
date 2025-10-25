package main

import (
	"fmt"
	"log"
	"os"

	"os-memory/pkg/cli"
	"os-memory/pkg/config"
	"os-memory/pkg/manager"
)

const configPath = "config.json"

func main() {
	cfg, err := config.LoadConfig(configPath)
	if err != nil {
		fmt.Printf("Erro fatal ao carregar configuração de '%s': %v\n", configPath, err)
		fmt.Println("Por favor, verifique se o arquivo 'config.json' existe e se os valores são potências de dois.")
		os.Exit(1)
	}

	mgr, err := manager.NewManager(cfg)
	if err != nil {
		log.Fatalf("Erro fatal ao inicializar o gerenciador de memória: %v", err)
	}

	app := cli.NewCLI(mgr)
	app.Run()
}
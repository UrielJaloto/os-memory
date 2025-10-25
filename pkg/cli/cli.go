package cli

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"os-memory/pkg/manager"
)

type CLI struct {
	manager *manager.MemoryManager
	reader  *bufio.Reader
}

func NewCLI(m *manager.MemoryManager) *CLI {
	return &CLI{
		manager: m,
		reader:  bufio.NewReader(os.Stdin),
	}
}

func (c *CLI) Run() {
	for {
		c.displayMenu()
		choiceStr, _ := c.reader.ReadString('\n')
		choice, err := strconv.Atoi(strings.TrimSpace(choiceStr))
		if err != nil {
			fmt.Println("\nOpção inválida. Por favor, digite um número.")
			continue
		}

		fmt.Println("---")
		switch choice {
		case 1:
			c.handleViewMemory()
		case 2:
			c.handleCreateProcess()
		case 3:
			c.handleViewPageTable()
		case 0:
			fmt.Println("Encerrando simulador...")
			return
		default:
			fmt.Println("Opção inválida. Tente novamente.")
		}
		fmt.Println("---")
	}
}

func (c *CLI) displayMenu() {
	fmt.Println("\n--- Simulador de Gerenciamento de Memória ---")
	fmt.Println("1. Visualizar Memória")
	fmt.Println("2. Criar Processo")
	fmt.Println("3. Visualizar Tabela de Páginas de um Processo")
	fmt.Println("0. Sair")
	fmt.Print("Escolha uma opção: ")
}

func (c *CLI) handleViewMemory() {
	percentage, frameStatus, frameSamples := c.manager.GetMemoryStatus()

	fmt.Printf("Memória Livre: %.2f%%\n\n", percentage)
	fmt.Println("Estado dos Quadros (Frames):")

	if len(frameStatus) == 0 {
		fmt.Println("Nenhum quadro de memória configurado.")
		return
	}

	for i, status := range frameStatus {
		fmt.Printf("  Quadro %d: %-20s %s\n", i, status, frameSamples[i])
	}
}

func (c *CLI) handleCreateProcess() {
	id, err := c.readInt("Digite o ID do processo: ")
	if err != nil {
		fmt.Printf("Erro ao ler ID: %v\n", err)
		return
	}

	size, err := c.readInt("Digite o tamanho do processo (em bytes): ")
	if err != nil {
		fmt.Printf("Erro ao ler tamanho: %v\n", err)
		return
	}

	err = c.manager.CreateProcess(id, size)
	if err != nil {
		fmt.Printf("\nErro ao criar processo: %v\n", err)
		return
	}

	fmt.Println("\nProcesso criado com sucesso!")
}

func (c *CLI) handleViewPageTable() {
	id, err := c.readInt("Digite o ID do processo: ")
	if err != nil {
		fmt.Printf("Erro ao ler ID: %v\n", err)
		return
	}

	size, table, err := c.manager.GetProcessPageTable(id)
	if err != nil {
		fmt.Printf("\nErro ao visualizar tabela: %v\n", err)
		return
	}

	fmt.Printf("\nTabela de Páginas do Processo %d (Tamanho: %d B):\n", id, size)
	if len(table.Entries) == 0 {
		fmt.Println("  Processo não possui páginas alocadas.")
		return
	}

	pages := make([]int, 0, len(table.Entries))
	for pageIndex := range table.Entries {
		pages = append(pages, pageIndex)
	}
	sort.Ints(pages)

	for _, pageIndex := range pages {
		frameIndex := table.Entries[pageIndex]
		fmt.Printf("  Página %d -> Quadro %d\n", pageIndex, frameIndex)
	}
}

func (c *CLI) readInt(prompt string) (int, error) {
	fmt.Print(prompt)
	inputStr, _ := c.reader.ReadString('\n')
	input, err := strconv.Atoi(strings.TrimSpace(inputStr))
	if err != nil {
		return 0, fmt.Errorf("entrada inválida, esperado um número inteiro")
	}
	return input, nil
}
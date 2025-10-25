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

const (
	colorReset  = "\033[0m"
	colorRed    = "\033[31m"
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
	colorCyan   = "\033[36m"
	colorBold   = "\033[1m"
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
		c.clearScreen()
		c.displayMenu()
		choiceStr, _ := c.reader.ReadString('\n')
		choice, err := strconv.Atoi(strings.TrimSpace(choiceStr))
		if err != nil {
			fmt.Println(colorRed, "\nOpção inválida. Por favor, digite um número.", colorReset)
			c.pause()
			continue
		}

		c.clearScreen()
		var needsPause bool = true

		switch choice {
		case 1:
			c.handleViewMemory()
		case 2:
			c.handleCreateProcess()
		case 3:
			c.handleViewPageTable()
		case 0:
			fmt.Println(colorCyan, "Encerrando simulador...", colorReset)
			needsPause = false
			return
		default:
			fmt.Println(colorRed, "Opção inválida. Tente novamente.", colorReset)
		}

		if needsPause {
			c.pause()
		}
	}
}

func (c *CLI) clearScreen() {
	fmt.Print("\033[H\033[2J")
}

func (c *CLI) pause() {
	fmt.Print(colorYellow, "\n\n(Pressione ENTER para continuar...)", colorReset)
	c.reader.ReadString('\n')
}

func (c *CLI) printHeader(title string) {
	fmt.Printf("%s%s== %s ==%s\n\n", colorBold, colorCyan, title, colorReset)
}

func (c *CLI) displayMenu() {
	fmt.Printf("%s%s--- Simulador de Gerenciamento de Memória ---%s\n", colorBold, colorCyan, colorReset)
	fmt.Println("1. Visualizar Memória")
	fmt.Println("2. Criar Processo")
	fmt.Println("3. Visualizar Tabela de Páginas de um Processo")
	fmt.Printf("%s0. Sair%s\n", colorRed, colorReset)
	fmt.Print(colorYellow, "\nEscolha uma opção: ", colorReset)
}

func (c *CLI) handleViewMemory() {
	c.printHeader("Visualizar Memória")
	percentage, frameStatus, frameSamples := c.manager.GetMemoryStatus()

	fmt.Printf("%sMemória Livre: %.2f%%%s\n\n", colorBold, percentage, colorReset)
	fmt.Printf("%sEstado dos Quadros (Frames):%s\n", colorBold, colorReset)

	if len(frameStatus) == 0 {
		fmt.Println("Nenhum quadro de memória configurado.")
		return
	}

	for i, status := range frameStatus {
		statusStr := status
		if status == "Livre" {
			statusStr = fmt.Sprintf("%s%-20s%s", colorGreen, status, colorReset)
		} else {
			statusStr = fmt.Sprintf("%s%-20s%s", colorYellow, status, colorReset)
		}

		fmt.Printf("  Quadro %d: %s %s\n", i, statusStr, frameSamples[i])
	}
}

func (c *CLI) handleCreateProcess() {
	c.printHeader("Criar Novo Processo")

	id, err := c.readInt("Digite o ID do processo: ")
	if err != nil {
		fmt.Printf("%sErro ao ler ID: %v%s\n", colorRed, err, colorReset)
		return
	}

	size, err := c.readInt("Digite o tamanho do processo (em bytes): ")
	if err != nil {
		fmt.Printf("%sErro ao ler tamanho: %v%s\n", colorRed, err, colorReset)
		return
	}

	err = c.manager.CreateProcess(id, size)
	if err != nil {
		fmt.Printf("\n%sErro ao criar processo: %v%s\n", colorRed, err, colorReset)
		return
	}

	fmt.Printf("\n%sProcesso criado com sucesso!%s\n", colorGreen, colorReset)
}

func (c *CLI) handleViewPageTable() {
	c.printHeader("Visualizar Tabela de Páginas")
	id, err := c.readInt("Digite o ID do processo: ")
	if err != nil {
		fmt.Printf("%sErro ao ler ID: %v%s\n", colorRed, err, colorReset)
		return
	}

	size, table, err := c.manager.GetProcessPageTable(id)
	if err != nil {
		fmt.Printf("\n%sErro ao visualizar tabela: %v%s\n", colorRed, err, colorReset)
		return
	}

	fmt.Printf("\n%sTabela de Páginas do Processo %d (Tamanho: %d B):%s\n", colorBold, id, size, colorReset)
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
	fmt.Print(colorYellow, prompt, colorReset)
	inputStr, _ := c.reader.ReadString('\n')
	input, err := strconv.Atoi(strings.TrimSpace(inputStr))
	if err != nil {
		return 0, fmt.Errorf("entrada inválida, esperado um número inteiro")
	}
	return input, nil
}
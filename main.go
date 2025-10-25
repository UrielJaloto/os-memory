package main

import (
	"fmt"
	"sistop/internal"
)

const PAGE_SIZE = 4096

func main() {
	cfg, err := internal.NewConfig("./config.json")
	if err != nil {
		panic(err)
	}

	fmt.Println("Config loaded:", cfg)

	// tui := &tui{}
	// memFisicaSize := tui.getMemoriaFisicaSize()

	// memFisica := make([]byte, memFisicaSize)
	// memLogicaSize := tui.getMemoriaLogicaSize()

	// memLogica := make([]byte, memLogicaSize)
}

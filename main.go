package main

type terminal interface {
	getMemoriaLogicaSize() int
}

type tui struct {
}

func (t *tui) getMemoriaLogicaSize() int {
	return 200000000
}

func main() {

	memFisica := make([]byte, 500000000)
	tui := &tui{}
	memLogicaSize := tui.getMemoriaLogicaSize()

	memLogica := make([]byte, memLogicaSize)
}

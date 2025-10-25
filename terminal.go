package main

type terminal interface {
	getMemoriaLogicaSize() int
}

type tui struct {
}

func (t *tui) getMemoriaLogicaSize() int {
	return 200000000
}

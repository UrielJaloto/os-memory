package internal

type terminal interface {
	getMemoriaLogicaSize() int
	getMemoriaFisicaSize() int
}

type tui struct {
}

func (t *tui) getMemoriaLogicaSize() int {
	return 200000000
}

func (t *tui) getMemoriaFisicaSize() int {
	return 1000000000
}

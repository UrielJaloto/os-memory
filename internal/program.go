package internal

type program struct {
	id            string
	tamanho       int // em bytes
	numeroPaginas int
}

func NewProgram(id string, tamanho int, pageSize int) *program {
	return &program{
		id:            id,
		tamanho:       tamanho,
		numeroPaginas: tamanho / pageSize,
	}
}

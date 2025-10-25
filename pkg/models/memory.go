package models

type PageTable struct {
	Entries map[int]int
}

func NewPageTable() *PageTable {
	return &PageTable{
		Entries: make(map[int]int),
	}
}

type Process struct {
	ID        int
	Size      int
	PageTable *PageTable
}

type PhysicalMemory struct {
	Data []byte
}

func NewPhysicalMemory(size int) *PhysicalMemory {
	return &PhysicalMemory{
		Data: make([]byte, size),
	}
}
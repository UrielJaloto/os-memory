package manager

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
	"time"

	"os-memory/pkg/config"
	"os-memory/pkg/models"
)

type MemoryManager struct {
	config     *config.Config
	memory     *models.PhysicalMemory
	processes  map[int]*models.Process
	freeFrames []bool
	numFrames  int
}

func NewManager(cfg *config.Config) (*MemoryManager, error) {
	if cfg == nil {
		return nil, fmt.Errorf("configuração não pode ser nula")
	}

	memory := models.NewPhysicalMemory(cfg.TamanhoMemoriaFisica)
	numFrames := cfg.TamanhoMemoriaFisica / cfg.TamanhoPagina

	freeFrames := make([]bool, numFrames)
	for i := 0; i < numFrames; i++ {
		freeFrames[i] = true
	}

	rand.Seed(time.Now().UnixNano())

	return &MemoryManager{
		config:     cfg,
		memory:     memory,
		processes:  make(map[int]*models.Process),
		freeFrames: freeFrames,
		numFrames:  numFrames,
	}, nil
}

func (m *MemoryManager) CreateProcess(id int, size int) error {
	if size <= 0 {
		return fmt.Errorf("tamanho do processo deve ser maior que zero")
	}

	if size > m.config.TamanhoMaximoProcesso {
		return fmt.Errorf("tamanho solicitado (%d B) excede o máximo permitido (%d B)", size, m.config.TamanhoMaximoProcesso)
	}

	if _, exists := m.processes[id]; exists {
		return fmt.Errorf("processo com ID %d já existe", id)
	}

	numPages := int(math.Ceil(float64(size) / float64(m.config.TamanhoPagina)))

	freeFrameCount := 0
	for _, isFree := range m.freeFrames {
		if isFree {
			freeFrameCount++
		}
	}

	if numPages > freeFrameCount {
		return fmt.Errorf("memória física insuficiente. Necessário: %d quadros, Disponível: %d quadros", numPages, freeFrameCount)
	}

	pageTable := models.NewPageTable()
	allocatedFrameCount := 0
	frameSearchIndex := 0

	for pageIndex := 0; pageIndex < numPages; pageIndex++ {

		for ; frameSearchIndex < m.numFrames; frameSearchIndex++ {
			if m.freeFrames[frameSearchIndex] {

				m.freeFrames[frameSearchIndex] = false
				pageTable.Entries[pageIndex] = frameSearchIndex

				m.initializeFrameData(frameSearchIndex)

				allocatedFrameCount++
				frameSearchIndex++
				break
			}
		}
	}

	process := &models.Process{
		ID:        id,
		Size:      size,
		PageTable: pageTable,
	}

	m.processes[id] = process
	return nil
}

func (m *MemoryManager) initializeFrameData(frameIndex int) {
	pageSize := m.config.TamanhoPagina
	startIndex := frameIndex * pageSize
	endIndex := startIndex + pageSize

	if endIndex > len(m.memory.Data) {
		endIndex = len(m.memory.Data)
	}

	frameSlice := m.memory.Data[startIndex:endIndex]
	rand.Read(frameSlice)
}

func (m *MemoryManager) GetMemoryStatus() (float64, []string, []string) {
	frameStatus := make([]string, m.numFrames)
	frameSamples := make([]string, m.numFrames)

	reverseMap := make(map[int]string)
	for procID, process := range m.processes {
		for page, frame := range process.PageTable.Entries {
			reverseMap[frame] = fmt.Sprintf("Proc %d (P%d)", procID, page)
		}
	}

	freeFrameCount := 0
	for frameIndex, isFree := range m.freeFrames {
		if isFree {
			freeFrameCount++
			frameStatus[frameIndex] = "Livre"
			frameSamples[frameIndex] = "[vazio]"
		} else {
			if status, ok := reverseMap[frameIndex]; ok {
				frameStatus[frameIndex] = status
			} else {
				frameStatus[frameIndex] = "Ocupado (Desconhecido)"
			}

			pageSize := m.config.TamanhoPagina
			startIndex := frameIndex * pageSize

			sampleSize := 4
			if sampleSize > pageSize {
				sampleSize = pageSize
			}

			endIndex := startIndex + sampleSize
			if endIndex > len(m.memory.Data) {
				endIndex = len(m.memory.Data)
			}

			sampleBytes := m.memory.Data[startIndex:endIndex]

			var s strings.Builder
			s.WriteString("[")
			for i, b := range sampleBytes {
				s.WriteString(fmt.Sprintf("0x%02x", b))
				if i < len(sampleBytes)-1 {
					s.WriteString(" ")
				}
			}
			
			if pageSize > sampleSize {
				s.WriteString(" ...]")
			} else {
				s.WriteString("]")
			}
			frameSamples[frameIndex] = s.String()
		}
	}

	percentage := (float64(freeFrameCount) / float64(m.numFrames)) * 100.0
	return percentage, frameStatus, frameSamples
}

func (m *MemoryManager) GetProcessPageTable(id int) (int, *models.PageTable, error) {
	process, exists := m.processes[id]
	if !exists {
		return 0, nil, fmt.Errorf("processo com ID %d não encontrado", id)
	}

	return process.Size, process.PageTable, nil
}
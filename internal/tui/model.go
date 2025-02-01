package tui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/evertras/bubble-table/table"
	"github.com/pedrolopesme/books/internal/proclist"
)

const (
	columnKeyPID    = "pid"
	columnKeyName   = "name"
	columnKeyCPU    = "cpu"
	columnKeyMemory = "memory"

	fixedVerticalMargin = 4
)

type Model struct {
	flexTable        table.Model
	totalWidth       int
	totalHeight      int
	horizontalMargin int
	verticalMargin   int
	columnSortKey    string
}

func InitialModel() (Model, error) {
	procs, err := proclist.GetProcesses()
	if err != nil {
		return Model{}, fmt.Errorf("failed to fetch processes: %w", err)
	}

	rows := fetchProcessRows(procs)
	t := newTable(rows)

	return Model{flexTable: t}, nil
}

func (m Model) Init() tea.Cmd {
	return nil
}

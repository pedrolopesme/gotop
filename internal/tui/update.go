package tui

import (
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		return m.handleKeyInput(msg), nil

	case tea.WindowSizeMsg:
		m.updateTableSize(msg)
	}

	m.flexTable, _ = m.flexTable.Update(msg)
	return m, nil
}

func (m Model) handleKeyInput(msg tea.KeyMsg) Model {
	switch msg.String() {
	case "ctrl+c", "esc", "q":
		os.Exit(0)
	case "n":
		m.columnSortKey = columnKeyName
		m.flexTable = m.flexTable.SortByAsc(m.columnSortKey)
	case "c":
		m.columnSortKey = columnKeyCPU
		m.flexTable = m.flexTable.SortByAsc(m.columnSortKey).ThenSortByDesc(m.columnSortKey)
	case "m":
		m.columnSortKey = columnKeyMemory
		m.flexTable = m.flexTable.SortByDesc(m.columnSortKey)
	}
	return m
}

func (m *Model) updateTableSize(msg tea.WindowSizeMsg) {
	m.totalWidth = msg.Width
	m.totalHeight = msg.Height
	m.flexTable = m.flexTable.
		WithTargetWidth(m.calculateWidth()).
		WithMinimumHeight(m.calculateHeight())
}

func (m Model) calculateWidth() int {
	return m.totalWidth - m.horizontalMargin
}

func (m Model) calculateHeight() int {
	return m.totalHeight - m.verticalMargin - fixedVerticalMargin
}

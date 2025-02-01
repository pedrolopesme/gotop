package tui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

func StartTUI() error {
	model, err := InitialModel()
	if err != nil {
		return fmt.Errorf("error initializing TUI: %w", err)
	}

	if _, err := tea.NewProgram(model).Run(); err != nil {
		return fmt.Errorf("error starting TUI: %w", err)
	}
	return nil
}

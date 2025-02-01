package tui

import (
	"fmt"
	"log"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/pedrolopesme/books/internal/proclist"
	"github.com/pedrolopesme/books/pkg/utils"
)

type Model struct {
	table table.Model
}

func InitialModel() Model {
	procs, err := proclist.GetProcesses()
	if err != nil {
		log.Fatal(err)
	}

	rows := make([]table.Row, len(procs))
	for i, p := range procs {
		rows[i] = table.Row{
			fmt.Sprintf("%d", p.Pid),
			p.Name,
			fmt.Sprintf("%.2f%%", p.CPUPerc),
			utils.FormatBytes(p.Memory),
		}
	}

	t := table.New(
		table.WithColumns([]table.Column{
			{Title: "PID", Width: 10},
			{Title: "Name", Width: 40},
			{Title: "CPU Usage", Width: 10},
			{Title: "Memory", Width: 10},
		}),
		table.WithRows(rows),
		table.WithFocused(true),
		// table.WithHeight(7),
	)

	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderBottom(true).
		Bold(false)

	t.SetStyles(s)

	return Model{
		table: t,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}
	m.table, cmd = m.table.Update(msg)
	return m, cmd
}

func (m Model) View() string {
	return m.table.View()
}

func StartTUI() error {
	if _, err := tea.NewProgram(InitialModel()).Run(); err != nil {
		return fmt.Errorf("error starting TUI: %w", err)
	}
	return nil
}

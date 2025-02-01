package tui

import (
	"strings"
)

func (m Model) View() string {
	var body strings.Builder
	body.WriteString("A filtered simple default table\n" +
		"Currently filter by Title and Author, press / + letters to start filtering, and escape to clear filter.\nPress q or ctrl+c to quit\n\n")
	body.WriteString(m.flexTable.View())
	return body.String()
}

package tui

import (
	"fmt"

	"github.com/evertras/bubble-table/table"
	"github.com/pedrolopesme/books/internal/proclist"
	"github.com/pedrolopesme/books/pkg/utils"
)

func fetchProcessRows(procs []proclist.ProcessInfo) []table.Row {
	rows := make([]table.Row, len(procs))
	for i, p := range procs {
		rows[i] = table.NewRow(table.RowData{
			columnKeyPID:    fmt.Sprintf("%d", p.Pid),
			columnKeyName:   p.Name,
			columnKeyCPU:    fmt.Sprintf("%.2f%%", p.CPUPerc),
			columnKeyMemory: utils.FormatBytes(p.Memory),
		})
	}
	return rows
}

func newTable(rows []table.Row) table.Model {
	return table.New([]table.Column{
		table.NewColumn(columnKeyPID, "PID", 10),
		table.NewFlexColumn(columnKeyName, "Name", 1),
		table.NewColumn(columnKeyCPU, "CPU", 10),
		table.NewColumn(columnKeyMemory, "Memory", 10),
	}).
		WithRows(rows).
		Focused(true)
}

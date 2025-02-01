package main

import (
	"fmt"

	"github.com/shirou/gopsutil/v4/process"
)

func main() {

	processes, err := process.Processes()
	if err != nil {
		fmt.Println("Impossible to get processes")
		return
	}

	for _, p := range processes {
		pid := p.Pid

		name, err := p.Name()
		if err != nil {
			continue
		}

		fmt.Printf("PID: %d, Name: %s\n", pid, name)
	}
}

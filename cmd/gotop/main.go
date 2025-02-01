package main

import (
	"fmt"
	"log"

	"github.com/pedrolopesme/books/internal/proclist"
)

func main() {
	procs, err := proclist.GetProcesses()
	if err != nil {
		log.Fatal(err)
	}

	for _, p := range procs {
		fmt.Printf("PID: %d, Name: %s\n", p.Pid, p.Name)
	}
}

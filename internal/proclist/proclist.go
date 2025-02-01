package proclist

import (
	"github.com/shirou/gopsutil/v4/process"
)

type ProcessInfo struct {
	Pid     int32
	Name    string
	CPUPerc float64
	Memory  uint64 // in bytes
}

// GetProcesses retrieves a list of currently running processes and their associated information,
// including PID, name, and CPU usage.
func GetProcesses() ([]ProcessInfo, error) {
	processes, err := process.Processes()
	if err != nil {
		return nil, err
	}

	procs := make([]ProcessInfo, 0, len(processes))
	for _, p := range processes {
		pid := p.Pid

		name, err := p.Name()
		if err != nil {
			continue
		}

		cpuPerc, err := getCPUPercent(p)
		if err != nil {
			continue
		}

		memInfo, err := p.MemoryInfo()
		if err != nil {
			continue
		}

		procs = append(procs, ProcessInfo{
			Pid:     pid,
			Name:    name,
			CPUPerc: cpuPerc,
			Memory:  memInfo.RSS,
		})
	}

	return procs, nil
}

func getCPUPercent(p *process.Process) (float64, error) {
	cpuUsage, err := p.CPUPercent()
	if err != nil {
		return 0, err
	}

	return cpuUsage, nil

}

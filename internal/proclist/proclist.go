package proclist

import "github.com/shirou/gopsutil/v4/process"

type ProcessInfo struct {
	Pid  int32
	Name string
}

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

		procs = append(procs, ProcessInfo{Pid: pid, Name: name})
	}

	return procs, nil
}

package supervisordtypes

type ProcessState struct {
	Pid           int
	Name          string
	Exe           string
	CommandLine   string
	Executable    string
	MemoryPercent int
	CPUPercent    int
	OpenFiles     int
	Connections   int
}

type ProcessStates []ProcessState

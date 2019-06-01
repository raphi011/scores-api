package job

// State represents the states a job can be in
type State int

const (
	// StateStopped is set if a job was explicitly stopped
	StateStopped State = iota
	// StateWaiting is set if a job is in the queue waiting to be run
	StateWaiting
	// StateRunning is set if a job is currently running
	StateRunning
	// StateErrored is set if a job is in an error state
	StateErrored
)

// String returns the State alias
func (s State) String() string {
	switch s {
	case StateStopped:
		return "stopped"
	case StateWaiting:
		return "waiting"
	case StateRunning:
		return "running"
	case StateErrored:
		return "errored"
	default:
		return "unknown"
	}
}

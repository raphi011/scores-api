package job

type State int

const (
	StateStopped State = iota
	StateWaiting
	StateRunning
	StateStopping
	StateErrored
)

func (s State) String() string {
	switch s {
	case StateStopped:
		return "stopped"
	case StateStopping:
		return "stopping"
	case StateWaiting:
		return "waiting"
	case StateRunning:
		return "running"
	case StateErrored:
		return "errored"
	default:
		return "unkown"
	}
}

package fsm

type State string

const (
	StateInitial State = "initial"
	StateEnd     State = "end"
)

// FSM – finite-state machine.
type FSM map[State][]State

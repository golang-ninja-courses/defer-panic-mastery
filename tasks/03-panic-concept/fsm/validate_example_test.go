package fsm

const (
	StateWaitForBet     State = "wait-for-bet"
	StateWaitForSpin    State = "wait-for-spin"
	StateWheelSpinning  State = "wheel-spinning"
	StateNumberReceived State = "number-received"
)

var RouletteFSM = FSM{
	StateInitial:        {StateWaitForBet},
	StateWaitForBet:     {StateWaitForSpin},
	StateWaitForSpin:    {StateWheelSpinning},
	StateWheelSpinning:  {StateNumberReceived},
	StateNumberReceived: {StateEnd},
}

func init() {
	if err := Validate(RouletteFSM); err != nil {
		panic(err)
	}
}

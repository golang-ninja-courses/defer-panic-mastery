package fsm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidate(t *testing.T) {
	const (
		state1 State = "state-1"
		state2 State = "state-2"
		state3 State = "state-3"
		state4 State = "state-4"
		state5 State = "state-5"
		state6 State = "state-6"
	)

	cases := []struct {
		name    string
		f       FSM
		wantErr bool
	}{
		{
			name: "no initial state",
			f: FSM{
				state1: {state2},
				state2: {state3},
				state3: {state4},
				state4: {state5},
				state5: {StateEnd},
			},
			wantErr: true,
		},
		{
			name: "no end state",
			f: FSM{
				StateInitial: {state1},
				state1:       {state2},
				state2:       {state3},
				state3:       {state4},
				state4:       {state5},
			},
			wantErr: true,
		},
		{
			name: "no transition",
			f: FSM{
				StateInitial: {state1},
				state1:       {state2},
				state3:       {state4},
				state4:       {state5},
				state5:       {StateEnd},
			},
			wantErr: true,
		},
		{
			name: "loop without end",
			f: FSM{
				StateInitial: {state1},
				state1:       {state2},
				state2:       {state3},
				state3:       {state4},
				state4:       {state2},
				state5:       {StateEnd},
			},
			wantErr: true,
		},
		{
			name: "many loops without end",
			f: FSM{
				StateInitial: {state1},
				state1:       {state2, state1},
				state2:       {state4, state3},
				state3:       {state1},
				state4:       {state3, state4},
				state5:       {StateEnd},
			},
			wantErr: true,
		},
		{
			name: "orphan vertices 1",
			f: FSM{
				StateInitial: {StateEnd},
				state1:       {state2},
				state2:       {state3},
				state3:       {state4},
				state4:       {state5},
				state5:       {StateEnd},
			},
			wantErr: true,
		},
		{
			name: "orphan vertices 2",
			f: FSM{
				StateInitial: {StateEnd},
				state1:       {state2},
				state2:       {state3},
				state3:       {state4},
				state4:       {state5},
				state5:       {StateEnd},
				StateEnd:     nil,
			},
			wantErr: true,
		},
		{
			name: "orphan vertices 3",
			f: FSM{
				StateInitial: {state1},
				state1:       {state2},
				state2:       {state3, state6},
				state3:       {state4},
				state4:       {state5},
				state5:       {StateEnd},
			},
			wantErr: true,
		},
		{
			name: "two separate loops 1",
			f: FSM{
				StateInitial: {state1},
				state1:       {state2},
				state2:       {state1},
				state3:       {state4},
				state4:       {state5},
				state5:       {state3},
				StateEnd:     nil,
			},
			wantErr: true,
		},
		{
			name: "two separate loops 2",
			f: FSM{
				StateInitial: {state1},
				state1:       {state2},
				state2:       {state1},
				state3:       {state4},
				state4:       {state5},
				state5:       {state3, state6},
				state6:       {StateEnd},
			},
			wantErr: true,
		},
		{
			name: "broken path 1",
			f: FSM{
				StateInitial: {state1},
				state1:       {state2},
				state2:       {state3},
				state3:       {state4, state5},
				state4:       nil,
				state5:       {StateEnd},
			},
			wantErr: true,
		},
		{
			name: "broken path 2",
			f: FSM{
				StateInitial: {state1},
				state1:       {state2},
				state2:       {state3, state4},
				state3:       {StateEnd},
				state4:       {state5},
				state5:       {state6},
			},
			wantErr: true,
		},
		{
			name: "broken path 3",
			f: FSM{
				StateInitial: {state3},
				state1:       {state2},
				state2:       {state3},
				state3:       {state4},
				state4:       {state5},
				state5:       {StateEnd},
			},
			wantErr: true,
		},
		{
			name: "broken path 4",
			f: FSM{
				StateInitial: {state1},
				state1:       {state2},
				state2:       {state3, state4},
				state3:       {state2},
				state4:       {state5},
				state5:       nil,
				StateEnd:     nil,
			},
			wantErr: true,
		},
		{
			name: "broken path 5",
			f: FSM{
				state1:       {StateInitial},
				StateInitial: {state2},
				state2:       {state3, state4},
				state3:       {state2},
				state4:       {state5},
				state5:       {StateEnd},
			},
			wantErr: true,
		},
		{
			name: "broken path 6",
			f: FSM{
				state1:       {state2},
				state2:       {StateInitial},
				StateInitial: {state3, state4},
				state3:       {state2},
				state4:       {state5},
				state5:       {StateEnd},
			},
			wantErr: true,
		},

		// Positive cases.
		{
			name: "simple working fsm",
			f: FSM{
				StateInitial: {state1},
				state1:       {state2},
				state2:       {state3},
				state3:       {state4},
				state4:       {state5},
				state5:       {StateEnd},
			},
			wantErr: false,
		},
		{
			name: "complex fsm but no loops",
			f: FSM{
				StateInitial: {state1, state5},
				state1:       {state2, state4, state5},
				state2:       {state3},
				state3:       {state4, state5},
				state4:       {state5},
				state5:       {StateEnd},
			},
			wantErr: false,
		},
		{
			name: "several paths to end",
			f: FSM{
				StateInitial: {state1, state3},
				state1:       {state2},
				state2:       {StateEnd},
				state3:       {state4},
				state4:       {state5},
				state5:       {StateEnd},
			},
			wantErr: false,
		},
		{
			name: "several paths to end and loop",
			f: FSM{
				StateInitial: {state1, state3},
				state1:       {state2},
				state2:       {StateEnd},
				state3:       {state1, state4},
				state4:       {state5},
				state5:       {state3, StateEnd},
			},
			wantErr: false,
		},
		{
			name: "two separate paths to end",
			f: FSM{
				StateInitial: {state1, state2},

				state1: {state3},
				state3: {state5},
				state5: {StateEnd},

				state2: {state4},
				state4: {state6},
				state6: {StateEnd},
			},
			wantErr: false,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			err := Validate(tt.f)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

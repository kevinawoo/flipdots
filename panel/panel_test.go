package panel

import (
	"testing"
	"reflect"
	"github.com/stretchr/testify/assert"
)

func TestNewPanel(t *testing.T) {
	t.SkipNow()
	tests := []struct {
		name string

		givenWidth, givenHeight int
		port                    string
		baud                    int

		expectedState State
	}{
		{"Create 2x4 panel", 2, 4, "/dev/tty", 9600,
			State{
				{false, false},
				{false, false},
				{false, false},
				{false, false},
			},
		},

		{"Create 8x4 panel", 8, 4, "/dev/tty", 9600,
			State{
				{false, false, false, false, false, false, false, false},
				{false, false, false, false, false, false, false, false},
				{false, false, false, false, false, false, false, false},
				{false, false, false, false, false, false, false, false},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			p := NewPanel(test.givenWidth, test.givenHeight, test.port, test.baud)
			defer p.Close()

			assert.Truef(t, reflect.DeepEqual(p.State, test.expectedState), "Expected state to look like: \n%s\n\nGot state looking like: \n%s\n", test.expectedState, p.State)
		})
	}
}

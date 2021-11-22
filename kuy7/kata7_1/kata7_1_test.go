package kata7_1

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

type equal struct {
	vrb1       float64
	vrb2, vrb3 int
	exp        int
}

var equalExampleForMean = []equal{
	{vrb1: 10,
		vrb2: 10,
		vrb3: 10,
		exp:  22},
	{vrb1: 10,
		vrb2: 10,
		vrb3: 5,
		exp:  29},
	{vrb1: 100,
		vrb2: 5,
		vrb3: 5,
		exp:  59},
}

func TestMyCodeBaby(t *testing.T) {
	var act int

	for _, eq := range equalExampleForMean {
		act = Evaporator(eq.vrb1, eq.vrb2, eq.vrb3)
		assert.Equal(t, eq.exp, act)
	}

}

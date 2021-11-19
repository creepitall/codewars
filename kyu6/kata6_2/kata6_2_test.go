package kata6_2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type equal struct {
	vrb string
	exp string
}

var equalExample = []equal{
	{vrb: "01|15|59, 1|47|16, 01|17|20, 1|32|34, 2|17|17",
		exp: "Range: 01|01|18 Average: 01|38|05 Median: 01|32|34"},
	{vrb: "02|15|59, 2|47|16, 02|17|20, 2|32|34, 2|17|17, 2|22|00, 2|31|41",
		exp: "Range: 00|31|17 Average: 02|26|18 Median: 02|22|00"},
}

func TestMyCodeBaby(t *testing.T) {
	var act string

	for _, eq := range equalExample {
		act = Stati(eq.vrb)
		assert.Equal(t, eq.exp, act)
	}
}

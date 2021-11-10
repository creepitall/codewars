package kata4_1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type equal struct {
	vrb []int
	exp string
}

var equalExample = []equal{
	{vrb: []int{15, 30, -45}, exp: "(2 30)(3 0)(5 0)"},
	{vrb: []int{12, 15}, exp: "(2 12)(3 27)(5 15)"},
	{vrb: []int{15, 21, 24, 30, 45}, exp: "(2 54)(3 135)(5 90)(7 21)"},
	{vrb: []int{107, 158, 204, 100, 118, 123, 126, 110, 116, 100},
		exp: "(2 1032)(3 453)(5 310)(7 126)(11 110)(17 204)(29 116)(41 123)(59 118)(79 158)(107 107)"},
	{vrb: []int{3526, 1842, -1277, 1224, -111, 702, 2348, 1062, 534, 1002, 3306, 2316, 784, -4483, -3099},
		exp: "(2 18646)(3 8778)(7 784)(13 702)(17 1224)(19 3306)(29 3306)(37 -111)(41 3526)(43 3526)(59 1062)(89 534)(167 1002)(193 2316)(307 1842)(587 2348)(1033 -3099)(1277 -1277)(4483 -4483)"},
}

func TestMyCodeBaby(t *testing.T) {
	var act string

	for _, eq := range equalExample {
		act = SumOfDivided(eq.vrb)
		assert.Equal(t, eq.exp, act)
	}
}

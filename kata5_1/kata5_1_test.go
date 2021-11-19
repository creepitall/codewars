package kata5_1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type equal struct {
	vrb [][]int
	exp string
}

var equalExample = []equal{
	{vrb: [][]int{{1, 2}, {1, 3}, {1, 4}},
		exp: "(6,12)(4,12)(3,12)"},
	{vrb: [][]int{{69, 130}, {87, 1310}, {30, 40}},
		exp: "(18078,34060)(2262,34060)(25545,34060)"},
	{vrb: [][]int{{2, 5}, {2, 5}, {4, 7}, {1, 2}, {3, 4}},
		exp: "(56,140)(56,140)(80,140)(70,140)(105,140)"},
	{vrb: [][]int{{4, 10}, {2, 5}, {4, 7}, {2, 4}, {6, 8}},
		exp: "(56,140)(56,140)(80,140)(70,140)(105,140)"},
	{vrb: [][]int{{1, 3}, {1, 3}, {3, 5}, {1, 2}, {1, 2}},
		exp: "(10,30)(10,30)(18,30)(15,30)(15,30)"},
	{vrb: [][]int{{6, 18}, {3, 9}, {6, 10}, {3, 6}, {9, 18}},
		exp: "(10,30)(10,30)(18,30)(15,30)(15,30)"},
}

func TestMyCodeBaby(t *testing.T) {
	var act string

	for qeID, eq := range equalExample {
		act = ConvertFracts(eq.vrb)
		assert.Equalf(t, eq.exp, act, "eq id: %v", qeID+1)
	}
}

func BenchmarkMyCodeBaby(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, t := range equalExample {
			ConvertFracts(t.vrb)
		}
	}
}

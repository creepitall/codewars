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
	//{vrb:[][]int{[]int{1, 2}, []int{1, 3}, []int{1, 4}},
	//	exp: "(6,12)(4,12)(3,12)"},
	//{vrb: [][]int{[]int{69, 130}, []int{87, 1310}, []int{30, 40}},
	//	exp: "(18078,34060)(2262,34060)(25545,34060)"},
	//{vrb: [][]int{[]int{2, 5}, []int{2, 5}, []int{4, 7}, []int{1, 2}, []int{3, 4}},
	//	exp: "(56,140)(56,140)(80,140)(70,140)(105,140)"},
	{vrb: [][]int{[]int{4, 10}, []int{2, 5}, []int{4, 7}, []int{2, 4}, []int{6, 8}},
		exp: "(56,140)(56,140)(80,140)(70,140)(105,140)"},
	//{vrb: [][]int{[]int{1, 3}, []int{1, 3}, []int{3, 5}, []int{1, 2}, []int{1, 2}},
	//	exp: "(10,30)(10,30)(18,30)(15,30)(15,30)"},
	//{vrb: [][]int{[]int{6, 18}, []int{3, 9}, []int{6, 10}, []int{3, 6}, []int{9, 18}},
	//	exp: "(10,30)(10,30)(18,30)(15,30)(15,30)"},
}

func TestMyCodeBaby(t *testing.T) {
	var act string

	for qeID, eq := range equalExample {
		act = ConvertFracts(eq.vrb)
		assert.Equalf(t, eq.exp, act, "eq id: %v", qeID+1)
	}
}

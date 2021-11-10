package kata4_1

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMyCodeBaby(t *testing.T) {
	var ans string

	lst1 := []int{12, 15}
	exp1 := "(2 12)(3 27)(5 15)"

	ans = SumOfDivided(lst1)
	assert.Equal(t, exp1, ans)

	lst2 := []int{15,21,24,30,45}
	exp2 := "(2 54)(3 135)(5 90)(7 21)"

	ans = SumOfDivided(lst2)
	assert.Equal(t, exp2, ans)

	lst3 := []int{107, 158, 204, 100, 118, 123, 126, 110, 116, 100}
	exp3 := "(2 1032)(3 453)(5 310)(7 126)(11 110)(17 204)(29 116)(41 123)(59 118)(79 158)(107 107)"

	ans = SumOfDivided(lst3)
	assert.Equal(t, exp3, ans)

	lst4 := []int{3526, 1842, -1277, 1224, -111, 702, 2348, 1062, 534, 1002, 3306, 2316, 784, -4483, -3099}
	exp4 := "(2 18646)(3 8778)(7 784)(13 702)(17 1224)(19 3306)(29 3306)(37 -111)(41 3526)(43 3526)(59 1062)(89 534)(167 1002)(193 2316)(307 1842)(587 2348)(1033 -3099)(1277 -1277)(4483 -4483)"

	ans = SumOfDivided(lst4)
	assert.Equal(t, exp4, ans)

	lst5 := []int{15, 30, -45}
	exp5 := "(2 30)(3 0)(5 0)"

	ans = SumOfDivided(lst5)
	assert.Equal(t, exp5, ans)
}
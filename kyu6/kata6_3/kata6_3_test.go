package kata6_3

import (
	"fmt"
	"math"
	"testing"

	"github.com/go-playground/assert/v2"
)

type equal struct {
	vrb string
	exp float64
}

var equalExampleForMean = []equal{
	{vrb: "London",
		exp: 51.199999999999996},
	{vrb: "Beijing",
		exp: 52.416666666666664},
}

var equalExampleForVariance = []equal{
	{vrb: "London",
		exp: 57.42833333333374},
	{vrb: "Beijing",
		exp: 4808.37138888889},
}

func TestMyCodeBaby(t *testing.T) {
	var act bool

	for _, eq := range equalExampleForMean {
		act = assertFuzzyMean(eq.vrb, RainData, eq.exp)
		assert.Equal(t, true, act)
	}

	for _, eq := range equalExampleForVariance {
		act = assertFuzzyVar(eq.vrb, RainData, eq.exp)
		assert.Equal(t, true, act)
	}
}

func assertFuzzyMean(town string, strng string, expect float64) bool {
	var inrange bool
	var merr float64 = 1e-2
	var actual = Mean(town, strng)
	inrange = (math.Abs(actual-expect) <= merr)
	if inrange == false {
		fmt.Printf("Expected should be near: %0.2e , got: %0.2e\n", expect, actual)
	}
	return inrange
}

func assertFuzzyVar(town string, strng string, expect float64) bool {
	var inrange bool
	var merr float64 = 1e-2
	var actual = Variance(town, strng)
	inrange = (math.Abs(actual-expect) <= merr)
	if inrange == false {
		fmt.Printf("Expected should be near: %0.2e , got: %0.2e\n", expect, actual)
	}
	return inrange
}

package kata6_1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type equal struct {
	vrb string
	exp string
}

var equalExample = []equal{
	{vrb: "din", exp: "((("},
	{vrb: "recede", exp: "()()()"},
	{vrb: "(( @", exp: "))(("},
	{vrb: "Success", exp: ")())())"},
}

func TestMyCodeBaby(t *testing.T) {
	var act string

	for _, eq := range equalExample {
		act = DuplicateEncode(eq.vrb)
		assert.Equal(t, eq.exp, act)
	}
}

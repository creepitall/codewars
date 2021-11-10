package kata6_1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type equal struct {
	word string
	exp  string
}

var equalExample = []equal{
	{word: "din", exp: "((("},
	{word: "recede", exp: "()()()"},
	{word: "(( @", exp: "))(("},
	{word: "Success", exp: ")())())"},
}

func TestMyCodeBaby(t *testing.T) {
	var act string

	for _, eq := range equalExample {
		act = DuplicateEncode(eq.word)
		assert.Equal(t, eq.exp, act)
	}
}

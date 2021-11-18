package kata5_1

import (
	"fmt"
)

func ConvertFracts(a [][]int) (str string) {
	for _, dopel := range a {
		var a, b int
		a = dopel[0]
		b = dopel[1]
		for a != 0 && b != 0 {
			if a > b {
				a = a % b
			} else {
				b = b % a
			}
		}

		del := b
		if a != 0 {
			del = a
		}

		dopel[0] = dopel[0]/del
		dopel[1] = dopel[1]/del
	}

	minimap := make(map[int]int)
	var i int = 1
	var nok int
	var exit bool
	for !exit {
		for _, dopel := range a {
			value := dopel[1]
			minimap[i*value] ++

			if minimap[i*value] == len(a){
				exit = true
				nok = i*value
				break
			}
		}
		i++
	}

	for _, vl := range a {
		str += fmt.Sprintf("(%v,%v)", nok/vl[1]*vl[0], nok)
	}

	return str
}

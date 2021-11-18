package kata5_1

import "fmt"

func ConvertFracts(a [][]int) (str string) {
	minimap := make(map[int]int)
	var i int = 1
	var nok int
	var exit bool

	newA := make([][]int, 0)
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

		newA = append(newA, []int{dopel[0]/del, dopel[1]/del})
	}

	for !exit {
			for _, dopel := range newA {
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

	for _, vl := range newA {
		str += fmt.Sprintf("(%v,%v)", nok/vl[1]*vl[0], nok)
	}

	return str
}
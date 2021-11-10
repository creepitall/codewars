package kata4_1

import (
	"fmt"
	"sort"
)

func SumOfDivided(lst []int) (str string) {
	var a,b int
	mainmap := make(map[int]int)
	for _, value := range lst {
		b = value
		if b < 0 {
			b = b * -1
		}
		searchMap := make(map[int]int)
		for i:=2; i<=b; i++ {
			if b % i == 0 {
				a, b = i, b / i
				if _,ok := searchMap[a]; !ok {
					mainmap[a] = mainmap[a] + value
					searchMap[a]++
				}
				i=1
			} else if b / i == 1 {
				if _,ok := searchMap[b]; !ok {
					mainmap[b] = mainmap[b] + value
					searchMap[b]++
				}
				break
			}
		}
	}

	keys := make([]int, 0, len(mainmap))
	for k := range mainmap {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	for _, k := range keys {
		str = str + fmt.Sprintf("(%v %v)", k, mainmap[k])
	}

	return str
}

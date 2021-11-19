package kata6_2

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func Stati(strg string) string {
	var seconds []int
	var all, r, av, ma int

	arrayStr := strings.Split(strg, ", ")
	for _, timeStr := range arrayStr {
		timePart := strings.Split(timeStr, "|")
		var sec int
		for id, value := range timePart {
			vlInt, err := strconv.Atoi(value)
			if err != nil {
				return ""
			}
			switch id {
			case 0:
				sec += vlInt * 60 * 60
			case 1:
				sec += vlInt * 60
			case 2:
				sec += vlInt
			}
		}
		all += sec
		seconds = append(seconds, sec)
	}
	sort.Ints(seconds)

	r = seconds[len(seconds)-1] - seconds[0]
	av = all / len(arrayStr)
	if len(seconds)%2 == 0 {
		coor := (len(seconds) / 2) - 1
		ma = (seconds[coor] + seconds[coor+1]) / 2
	} else {
		coor := (len(seconds) / 2)
		ma = seconds[coor]
	}

	return fmt.Sprintf("Range: %v Average: %v Median: %v", returnTime(r), returnTime(av), returnTime(ma))
}

func returnTime(x int) string {
	m := make(map[string]int)
	for x != 0 {
		if x >= 3600 {
			m["h"] += 1
			x -= 3600
		} else if x >= 60 {
			m["m"] += 1
			x -= 60
		} else {
			m["s"] += x
			x -= x
		}
	}

	return fmt.Sprintf("%02d|%02d|%02d", m["h"], m["m"], m["s"])
}

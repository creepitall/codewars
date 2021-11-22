package kata7_1

func Evaporator(content float64, evapPerDay int, threshold int) (i int) {
	deadline := (float64(threshold) / 100.00) * content
	for content > deadline {
		i++
		percent := 1.00 - (float64(evapPerDay) / 100.00)
		content = content * percent
	}
	return i
}

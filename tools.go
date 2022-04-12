package tools

impoert "strconv"

func atof(str string) float64 {
	f, _ := strconv.ParseFloat(str, 64)
	return f
}

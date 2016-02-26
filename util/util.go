package util

import (
	"fmt"
	"strconv"
)

func LimitFloat(f float64) float64 {
	str := fmt.Sprintf("%.6f", f)
	new_f, _ := strconv.ParseFloat(str, 64)
	return new_f
}

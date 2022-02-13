package mathUtil

import (
	"math"
	"strconv"
)

// RoundFloat64Unsafe float小数点保留位数，四舍五入，不能保证一定准确
// @param f
// @param scale 保留几位小数点后数字
// @return float64
func RoundFloat64Unsafe(f float64, scale int) float64 {
	result, _ := strconv.ParseFloat(strconv.FormatFloat(f, 'f', scale+1, 64), 64)
	pow := math.Pow(10, float64(scale))
	return math.Round(result*pow) / pow
}

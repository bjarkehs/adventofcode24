package helpers

import (
	"math"
	"strconv"
)

func Map[T, V any](ts []T, fn func(T) V) []V {
	result := make([]V, len(ts))
	for i, t := range ts {
		result[i] = fn(t)
	}
	return result
}

func MapToInts(stringValues []string) []int {
	return Map(stringValues, func(t string) int {
		value, err := strconv.Atoi(t)
		Check(err)
		return value
	})
}

func AbsInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func PowInt(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func RemoveIndex(s []int, index int) []int {
	ret := make([]int, 0)
	ret = append(ret, s[:index]...)
	return append(ret, s[index+1:]...)
}

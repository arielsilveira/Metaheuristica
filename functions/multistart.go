package functions

import (
	"fmt"
	"math"
)

func MultiStart(n int, s []int, distance [][]float64) (s_star []int, fo_star float64) {
	var fo float64
	iter := 0
	iterMax := 1000

	fo = math.MaxFloat64
	fo_star = math.MaxFloat64

	for iter < iterMax {
		s = RandomConstruction(n)
		fo, s = DescentRandomImprovement(n, s, distance, 500)

		if fo < fo_star {
			s_star = s
			fo_star = fo
			iter = 0
			fmt.Println("Função objetivo: ", fo_star)
		}
		iter++
	}

	return s_star, fo_star
}

package functions

import (
	"math"
)

type ConstructionMethod func(n int, distance [][]float64, alpha float64) (solution []int)

func Grasp(n int, solution []int, distance [][]float64, alpha float64, grasp_max int, Method ConstructionMethod) ([]int, float64) {
	var s []int
	var fo float64
	fo_star := math.MaxFloat32
	for iter_max := 0; iter_max < grasp_max; iter_max++ {
		s = Method(n, distance, alpha)
		fo, s = DescentBestImprovement(n, s, distance)

		if fo < fo_star {
			fo_star = fo
			solution = s

		}
	}

	return solution, fo_star
}

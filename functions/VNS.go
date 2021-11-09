package functions

import (
	"Metaheuristica/src"
	"fmt"
	"math/rand"
)

func VNS(n int, solution []int, distance [][]float64, vns_max int, r int) (fo float64, s_star []int) {
	k := 0
	fo = src.CalculateOF(solution, distance)
	var fo_neighbor float64

	for i := 0; i < vns_max; i++ {
		k = 1

		for k <= r {
			s_star = solution
			_, s_star = AnyNeighbor(n, s_star, distance, fo, k)
			fo_neighbor, s_star = DescentFirstImprovement(n, s_star, distance)
			if fo_neighbor < fo {
				solution = s_star
				fo = fo_neighbor
				k = 1
				fmt.Println(fo)
				fmt.Println(fo_neighbor)
				fmt.Println()
			} else {
				k++
			}
		}
	}

	return fo, solution
}

func AnyNeighbor(n int, solution []int, distance [][]float64, fo float64, k int) (fo_neighbor float64, s []int) {
	switch k {
	case 1:
		fo_neighbor, s = ReplaceNeighbor(n, solution, distance, fo)
	case 2:
		fo_neighbor, s = ReinsertionAnyNeighbor_1(n, solution, distance, fo)
	case 3:
		fo_neighbor, s = ReinsertionAnyNeighbor_2(n, solution, distance, fo)
	}

	return fo_neighbor, s
}

func ReplaceNeighbor(n int, solution []int, distance [][]float64, fo float64) (float64, []int) {
	fo_neighbor := fo
	var i, j int
	var delta1, delta2 float64

	j = rand.Intn(n)
	i = rand.Intn(n)
	for i == j {
		i = rand.Intn(n)
	}

	delta1 = DeltaCalculated(n, distance, solution, i, j)

	solution[i], solution[j] = solution[j], solution[i]

	delta2 = DeltaCalculated(n, distance, solution, i, j)

	fo_neighbor = fo - delta1 + delta2

	return fo_neighbor, solution
}

func ReinsertionAnyNeighbor_1(n int, solution []int, distance [][]float64, fo float64) (float64, []int) {
	var fo_neighbor float64

	var i, j, aux int

	i = rand.Intn(n-1) + 1
	j = rand.Intn(n)

	for i <= j {
		j = rand.Intn(n)
	}

	aux = solution[i]
	for i > j {
		solution[i] = solution[i-1]
		i--
	}

	solution[j] = aux

	fo_neighbor = src.CalculateOF(solution, distance)

	return fo_neighbor, solution
}

func ReinsertionAnyNeighbor_2(n int, solution []int, distance [][]float64, fo float64) (float64, []int) {
	var fo_neighbor float64

	var i, j, aux int

	i = rand.Intn(n - 1)
	j = rand.Intn(n)

	for i >= j {
		j = rand.Intn(n)
	}

	aux = solution[i]
	for i < j {
		solution[i] = solution[i+1]
		i++
	}

	solution[j] = aux

	fo_neighbor = src.CalculateOF(solution, distance)

	return fo_neighbor, solution
}

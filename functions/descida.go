package functions

import (
	"Metaheuristica/src"
	"math/rand"
)

func DeltaCalculated(n int, distance [][]float64, solution []int, i int, j int) (delta float64) {

	iBefore := i - 1
	iAfter := i + 1
	jBefore := j - 1
	jAfter := j + 1

	if iBefore < 0 {
		iBefore = n - 1
	} else if iAfter > n-1 {
		iAfter = 0
	}

	if jBefore < 0 {
		jBefore = n - 1
	} else if jAfter > n-1 {
		jAfter = 0
	}

	delta = distance[solution[iBefore]][solution[i]] + distance[solution[i]][solution[iAfter]] +
		distance[solution[jBefore]][solution[j]] + distance[solution[j]][solution[jAfter]]

	return delta
}

func BestNeighbor(n int, solution []int, distance [][]float64, fo float64, bestI int, bestJ int) (float64, int, int) {
	bestNeighbor := fo

	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			delta1 := DeltaCalculated(n, distance, solution, i, j)
			solution[i], solution[j] = solution[j], solution[i]

			delta2 := DeltaCalculated(n, distance, solution, i, j)

			neighborFO := fo - delta1 + delta2

			if neighborFO < bestNeighbor {
				bestI = i
				bestJ = j
				bestNeighbor = neighborFO
			}

			solution[i], solution[j] = solution[j], solution[i]
		}
	}
	return bestNeighbor, bestI, bestJ
}

func RandomNeighbor(n int, solution []int, distance [][]float64, fo float64, best_i int, best_j int) (float64, int, int) {
	var delta1, delta2 float64

	j := rand.Intn(n)
	i := rand.Intn(n)
	for i == j {
		i = rand.Intn(n)
	}

	delta1 = DeltaCalculated(n, distance, solution, i, j)
	solution[i], solution[j] = solution[j], solution[i]
	delta2 = DeltaCalculated(n, distance, solution, i, j)
	solution[i], solution[j] = solution[j], solution[i]

	best_i = i
	best_j = j

	return fo - delta1 + delta2, best_i, best_j
}

func DescentBestImprovement(n int, solution []int, distance [][]float64) (float64, []int) {
	var ibest, jbest int
	var fo_viz, fo float64
	flag := true
	iter := 0
	fo = src.CalculateOF(solution, distance)
	fo_viz = fo

	for flag {
		flag = false
		fo_viz, ibest, jbest = BestNeighbor(n, solution, distance, fo, ibest, jbest)

		if fo_viz < fo {
			iter++
			solution[ibest], solution[jbest] = solution[jbest], solution[ibest]
			fo = fo_viz
			flag = true
		}
	}

	return fo, solution

}

func DescentRandomImprovement(n int, solution []int, distance [][]float64, iterMax int) (float64, []int) {
	var ibest, jbest int
	var fo_viz, fo float64
	iter := 0
	fo = src.CalculateOF(solution, distance)
	fo_viz = fo

	for iter < iterMax {
		iter++
		fo_viz, ibest, jbest = RandomNeighbor(n, solution, distance, fo, ibest, jbest)

		if fo_viz < fo {
			iter = 0
			fo = fo_viz
			solution[ibest], solution[jbest] = solution[jbest], solution[ibest]
		}
	}

	return fo, solution

}

func NeighborFirstImprovement(n int, s []int, distance [][]float64, fo float64, best_i int, best_j int) (float64, []int, int, int) {
	best_fo := fo
	best := false

	vet := RandomConstruction(n)

	for i := 0; i < n && best; i++ {
		for j := 0; j < n && best; j++ {
			aux_i := vet[i]
			aux_j := vet[j]

			d1 := DeltaCalculated(n, distance, s, aux_i, aux_j)

			s[aux_i], s[aux_j] = s[aux_j], s[aux_i]

			d2 := DeltaCalculated(n, distance, s, aux_i, aux_j)

			s[aux_i], s[aux_j] = s[aux_j], s[aux_i]

			fo = fo - d1 + d2

			if fo < best_fo {
				best_fo = fo
				best_i = aux_i
				best_j = aux_j
				best = false
			}

		}
	}

	return fo, s, best_i, best_j
}

func DescentFirstImprovement(n int, solution []int, distance [][]float64) (float64, []int) {
	var best bool = false
	var iter int = 0
	var best_i int = 0
	var best_j int = 0

	fo := src.CalculateOF(solution, distance)
	fo_neighbor := fo

	for !best && iter < 500 {
		fo_neighbor, solution, best_i, best_j = NeighborFirstImprovement(n, solution, distance, fo, best_i, best_j)

		if fo_neighbor < fo {
			solution[best_i], solution[best_j] = solution[best_j], solution[best_i]
			fo = fo_neighbor
			best = true
		}

		iter++
	}

	return fo, solution
}

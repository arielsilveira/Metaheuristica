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

func randomNeighbor(n int, solution []int, distance [][]float64, fo float64, bestI int, bestJ int) (float64, int, int) {
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

	bestI = i
	bestJ = j

	return fo - delta1 + delta2, bestI, bestJ
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
		fo_viz, ibest, jbest = randomNeighbor(n, solution, distance, fo, ibest, jbest)

		if fo_viz < fo {
			iter = 0
			fo = fo_viz
			solution[ibest], solution[jbest] = solution[jbest], solution[ibest]
		}
	}

	return fo, solution

}

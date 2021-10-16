package functions

import (
	"Metaheuristica/src"
	"math/rand"
)

func perturbation(n int, solution []int, distance [][]float64, fo float64, nivel int) (float64, []int) {
	n_swap_max := nivel + 1
	var i, j int

	for n_swap := 0; n_swap < n_swap_max; n_swap++ {
		i = rand.Intn(n)
		j = rand.Intn(n)
		for i == j {
			j = rand.Intn(n)
		}

		delta1 := DeltaCalculated(n, distance, solution, i, j)
		solution[i], solution[j] = solution[j], solution[i]

		delta2 := DeltaCalculated(n, distance, solution, i, j)
		fo = fo - delta1 + delta2
	}
	return fo, solution
}

func ILS(n int, solution []int, distance [][]float64, level int, ils_max int) (float64, []int) {
	fo := src.CalculateOF(solution, distance)
	best_iter := 0

	for iter := 0; iter-best_iter < ils_max; iter++ {
		_, s := perturbation(n, solution, distance, fo, level)
		fo_star, s_star := DescentBestImprovement(n, s, distance)

		if fo_star < fo {
			solution = s_star
			fo = fo_star
			best_iter = iter
			level = 1
		} else {
			level++
		}
	}

	return fo, solution
}

func SmartILS(n int, solution []int, distance [][]float64, level int, level_occurrence int, ils_max int) (float64, []int) {
	fo := src.CalculateOF(solution, distance)
	max_occurrence := 1
	best_iter := 0

	for iter := 0; iter-best_iter < ils_max; iter++ {
		_, s := perturbation(n, solution, distance, fo, level)
		fo_star, s_star := DescentBestImprovement(n, s, distance)

		if fo_star < fo {
			solution = s_star
			fo = fo_star
			best_iter = iter
			level = 1
			level_occurrence = 1
		} else {
			if level_occurrence >= max_occurrence {
				level++
				level_occurrence = 1
			} else {
				level_occurrence++
			}
		}
	}
	return fo, solution
}

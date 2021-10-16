package functions

import (
	"Metaheuristica/src"
	"math/rand"
)

func perturbacao(n int, solution []int, distance [][]float64, fo float64, nivel int) (float64, []int) {
	ntrocasMax := nivel + 1
	var i, j int

	for ntrocas := 0; ntrocas < ntrocasMax; ntrocas++ {
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

func ILS(n int, solution []int, distance [][]float64, nivel int, ILSmax int) (float64, []int) {
	var fo float64
	var best_iter int
	fo = src.CalculateOF(solution, distance)

	for iter := 0; iter-best_iter < ILSmax; iter++ {
		_, s := perturbacao(n, solution, distance, fo, nivel)
		fo_star, s_star := DescentBestImprovement(n, s, distance)

		if fo_star < fo {
			solution = s_star
			fo = fo_star
			best_iter = iter
			nivel = 1
		} else {
			nivel++
		}
	}

	return fo, solution
}

func SmartILS(n int, solution []int, distance [][]float64, nivel int, nVezesNivel int, ILSMax int) (float64, []int) {
	var fo float64
	var best_iter int
	fo = src.CalculateOF(solution, distance)
	vezesMax := 1
	for iter := 0; iter-best_iter < ILSMax; iter++ {
		_, s := perturbacao(n, solution, distance, fo, nivel)
		fo_star, s_star := DescentBestImprovement(n, s, distance)

		if fo_star < fo {
			solution = s_star
			fo = fo_star
			best_iter = iter
			nivel = 1
			nVezesNivel = 1
		} else {
			if nVezesNivel >= vezesMax {
				nivel++
				nVezesNivel = 1
			} else {
				nVezesNivel++
			}
		}
	}
	return fo, solution
}

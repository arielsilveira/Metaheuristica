package functions

import (
	"Metaheuristica/src"
	"math"
	"math/rand"
)

func SA(n int, solution []int, distance [][]float64, alpha float64, sa_max int, temp float64, final_temp float64) (s_star []int, fo_star float64) {
	var fo, fo_viz float64
	var best_i, best_j int
	var iterT int = 0

	s_star = solution
	fo_viz = src.CalculateOF(solution, distance)
	fo = fo_viz
	fo_star = fo_viz

	for temp > final_temp {

		for iterT < sa_max {

			iterT++
			fo_viz, best_i, best_j = RandomNeighbor(n, solution, distance, fo, best_i, best_j)
			delta := fo_viz - fo

			if delta < 0 {
				fo = fo_viz
				solution[best_i], solution[best_j] = solution[best_j], solution[best_i]

				if fo_viz < fo_star {
					s_star = solution
					fo_star = fo_viz
				}

			} else {
				x := rand.Float64()
				if x < math.Exp(-delta/temp) {
					solution[best_i], solution[best_j] = solution[best_j], solution[best_i]
					fo = fo_viz
				}
			}
		}
		temp = alpha * temp
		iterT = 0
	}

	return s_star, fo_star
}

func InitialTemp(n int, solution []int, distance [][]float64, beta float64, gamma float64, sa_max int, temp float64) float64 {
	var best_i, best_j int
	var fo_viz, fo float64

	continua := true
	fo_viz = src.CalculateOF(solution, distance)
	fo = fo_viz

	for continua {
		aceitos := 0
		for i := 1; i <= sa_max; i++ {
			fo_viz, best_i, best_j = RandomNeighbor(n, solution, distance, fo, best_i, best_j)
			delta := fo_viz - fo
			if delta < 0 {
				aceitos++
			} else {
				x := rand.Float64()
				if x < math.Exp(-delta/temp) {
					aceitos++
				}
			}
		}

		if float64(aceitos) >= gamma*float64(sa_max) {
			continua = false
		} else {
			temp = beta * temp
		}

	}

	return temp
}

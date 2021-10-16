package functions

import (
	"Metaheuristica/src"
)

func BT(n int, solution []int, distance [][]float64, list_size int, bt_max int) ([]int, float64) {
	var best_i, best_j int
	var s_star []int
	var fo, fo_star float64
	tabu_list := src.CreateMatrix(n)

	s_star = solution
	fo = src.CalculateOF(solution, distance)
	fo_star = fo
	for iter_bt := 0; iter_bt < bt_max; iter_bt++ {
		fo_vizinho, best_i, best_j := BestNeighborBT(n, solution, distance, fo, best_i, best_j, fo_star, iter_bt, tabu_list)

		tabu_list[best_i][best_j] = float64(iter_bt + list_size)

		solution[best_i], solution[best_j] = solution[best_j], solution[best_i]

		fo = fo_vizinho

		if fo < fo_star {
			s_star = solution
			fo_star = fo
		}
	}

	return s_star, fo_star
}

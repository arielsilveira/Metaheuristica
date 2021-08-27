package functions

import (
	"math"
	"math/rand"
	"sort"

	"Metaheuristica/src"
)

func GreedySolution(n int, distance [][]float64) (solution []int) {
	var next int
	unvisited := src.Unvisited(n)

	solution = append(solution, 0)

	for i := 1; i < n; i++ {
		dist := math.MaxFloat64
		var pos = 0

		for j := 0; j < len(unvisited); j++ {
			if distance[solution[i-1]][unvisited[j]] < dist {
				dist = distance[solution[i-1]][unvisited[j]]
				next = unvisited[j]
				pos = j
			}
		}

		solution = append(solution, next)
		unvisited = src.RemoveElement(unvisited, pos)
	}

	return solution
}

type order struct {
	d         [][]float64
	index     int
	unvisited []int
}

func GreedySolutionNeighborNearby(n int, distance [][]float64, alpha float64) (solution []int) {
	var length int
	unvisited := src.Unvisited(n)

	var o order
	o.d = distance
	o.index = 0

	solution = append(solution, 0)
	j := 1
	var choiceCity int

	for j < n {
		length = len(unvisited)
		lengthLRC := math.Max(1, alpha*float64(length))

		pos := rand.Intn(int(lengthLRC))

		choiceCity = unvisited[pos]

		solution = append(solution, choiceCity)

		src.RemoveElement(unvisited, pos)

		sort.Ints(unvisited)
		// ordena
		o.index = choiceCity
		o.unvisited = unvisited

		// ordem.index = cidade_escolhida; //atualiza ultima cidade inserida
		// sort(nao_visitadas.begin(), nao_visitadas.end());  //ordena pelos indices (para manter a estabilidade)
		// stable_sort(nao_visitadas.begin(), nao_visitadas.end(), ordem); //ordena pelas distancias

		unvisited = o.unvisited

		j++
	}

	return solution
}

func CheaperGreedySolution(n int, distance [][]float64) (solution []int) {
	unvisited := src.Unvisited(n)

	solution = append(solution, 0)

	var next int

	var city_i, city_j, city_k, pos int
	var best_k int
	var sij float64
	i := 1

	for i < 3 {
		dist := math.MaxFloat64
		pos_j := 0
		for j := 0; j < len(unvisited); j++ {
			if distance[solution[i-1]][unvisited[j]] < dist {
				dist = distance[solution[i-1]][unvisited[j]]
				next = unvisited[j]
				pos_j = j
			}
		}
		solution = append(solution, next)
		unvisited = src.RemoveElement(unvisited, pos_j)
		i++
	}
	j := i
	for j < n {
		best_sij := math.MaxFloat64
		pos_k := 0

		for k := 0; k < len(unvisited); k++ {
			city_k = unvisited[k]

			for i := 0; i < j; i++ {
				city_i = solution[i]

				if (i + 1) != j {
					city_j = solution[i+1]
				} else {
					city_j = 0
				}

				sij = distance[city_i][city_k] + distance[city_k][city_j] - distance[city_i][city_j]

				if sij < best_sij {
					best_k = city_k
					best_sij = sij
					pos = i + 1
					pos_k = k
				}
			}
		}

		solution = src.InsertPos(solution, pos, best_k)
		unvisited = src.RemoveElement(unvisited, pos_k)

		j++
	}

	return solution
}

package functions

import (
	"math"
	"math/rand"
	"sort"
	"time"

	"Metaheuristica/src"
)

// constroi_solucao_gulosa_vizinho_mais_proximo
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

// constroi_solucao_parcialmente_gulosa_vizinho_mais_proximo
func GreedySolutionNeighborNearby(n int, distance [][]float64, alpha float64) (solution []int) {
	var length int
	unvisited := src.Unvisited(n)

	solution = append(solution, 0)

	var choiceCity int = 0

	sort.SliceStable(unvisited, func(i, j int) bool {
		return distance[choiceCity][unvisited[i]] < distance[choiceCity][unvisited[j]]
	})

	for j := 1; j < n; j++ {
		length = len(unvisited)
		lengthLRC := math.Max(1, alpha*float64(length))
		pos := rand.Intn(int(lengthLRC))

		choiceCity = unvisited[pos]

		solution = append(solution, choiceCity)

		unvisited = src.RemoveElement(unvisited, pos)

		sort.Ints(unvisited)

		sort.SliceStable(unvisited, func(i, j int) bool {
			return distance[choiceCity][unvisited[i]] < distance[choiceCity][unvisited[j]]
		})
	}

	return solution
}

// constroi_solucao_gulosa_insercao_mais_barata
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

// constroi_solucao_aleatoria
func RandomConstruction(n int) (solution []int) {
	for i := 0; i < n; i++ {
		solution = append(solution, i)
	}

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(n, func(i, j int) {
		solution[i], solution[j] = solution[j], solution[i]
	})

	return solution
}

// constroi_solucao_parcialmente_gulosa_insercao_mais_barata
func PartiallyGreedyWithCheaperInsert(n int, distance [][]float64, alpha float64) (solution []int) {
	var gMin, gMax float64
	unvisited := src.Unvisited(n)

	solution = append(solution, 0)

	var choiceCity int = 0
	for i := 1; i < n; i++ {
		var lrc []int
		sort.SliceStable(unvisited, func(i, j int) bool {
			return distance[choiceCity][unvisited[i]] < distance[choiceCity][unvisited[j]]
		})
		gMin = distance[choiceCity][unvisited[0]]
		gMax = distance[choiceCity][unvisited[len(unvisited)-1]]

		value := gMin + alpha*(gMax-gMin)
		for j := 0; j < len(unvisited); j++ {
			if distance[choiceCity][unvisited[j]] > value {
				break
			}
			lrc = append(lrc, unvisited[j])
		}

		pos := rand.Intn(len(lrc))

		choiceCity = unvisited[pos]

		solution = append(solution, choiceCity)

		unvisited = src.RemoveElement(unvisited, pos)
	}

	return solution
}

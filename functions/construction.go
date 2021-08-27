package functions

import (
	"math"
	"math/rand"

	"Metaheuristica/src"
)

func GreedySolution(n int, distance [][]float64) (solution []int) {
	var next int
	var unvisited []int

	for i := 1; i < n; i++ {
		unvisited = append(unvisited, i)
	}

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
	d     [][]float64
	index int
}

func GreedySolutionNeighborNearby(n int, distance [][]float64, alpha float64) (solution []int) {
	var length int
	var unvisited []int

	var o order
	o.d = distance
	o.index = 0

	// Sort by age, keeping original order or equal elements.
	// sort.SliceStable(o, func(i, j int) bool {
	// 	return o.d[o.index][j] < o.d[o.index][j]
	// })

	for i := 1; i < n; i++ {
		unvisited = append(unvisited, i)
	}

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

		// ordena
		j++
	}

	return solution
}

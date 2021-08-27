package menus

import (
	"fmt"

	"Metaheuristica/functions"
	"Metaheuristica/src"
)

func PrincipalMenu(distance [][]float64, n int) {
	var s []int
	options := []interface{}{
		"Sair",
		"Gere solucao inicial",
		"Descida com Best Improvement",
		"Descida randomica",
		"Descida com Primeiro de Melhora (First Improvement)",
		"Multi-Start",
		"Simulated Annealing",
		"Busca Tabu",
		"ILS",
		"GRASP",
		"VND",
		"VNS",
		"Algoritmos Geneticos",
		"LAHC",
		"Algoritmos Memeticos",
		"Colonia de Formigas",
	}

	for {
		choice := src.AskOption(src.QUESTION, options)

		switch choice {
		case 0:
			fmt.Println("BYE BYE!")
			return
		case 1:
			InitSolution(distance, n, s)
		case 2:
			fmt.Println("Not Implemented")
		default:
			fmt.Println("Not Implemented")
		}
	}
}

func InitSolution(distance [][]float64, n int, s []int) {
	var alpha float64
	options := []interface{}{
		"Voltar",
		"Gulosa (Vizinho mais proximo)",
		"Parcialmente gulosa (Vizinho mais proximo)",
		"Gulosa (Insercao Mais Barata)",
		"Parcialmente gulosa (Insercao Mais Barata)",
		"Aleatoria",
	}

	for {
		choice := src.AskOption(src.SOLUTION, options)

		switch choice {
		case 0:
			return
		case 1:
			s = functions.GreedySolution(n, distance)
			src.PrintInfos(s, distance)
		case 2:
			s := functions.GreedySolutionNeighborNearby(n, distance, alpha)
			src.PrintInfos(s, distance)
		case 3:
			fmt.Println("Not Implemented")
		case 4:
			fmt.Println("Not Implemented")
		case 5:
			fmt.Println("Not Implemented")
		}
	}
}
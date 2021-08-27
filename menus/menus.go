package menus

import (
	"fmt"

	"Metaheuristica/functions"
	"Metaheuristica/src"
)

func PrincipalMenu(distance [][]float64, n int, best_fit_lit float64) {
	var solution []int
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
			solution = InitSolution(distance, n)
		case 2:
			if len(solution) == 0 {
				fmt.Print("\n\n#### É necessario gerar a solução inicial. ####\n\n")
				break
			}
			BestImprovement(n, distance, solution)
		default:
			fmt.Println("Not Implemented")
		}
	}
}

func InitSolution(distance [][]float64, n int) (s []int) {
	var alpha float64 = 0.1
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
			return s
		case 1:
			s = functions.GreedySolution(n, distance)
			src.PrintInfos(s, distance)
		case 2:
			// fmt.Println("Not Implemented")
			s := functions.GreedySolutionNeighborNearby(n, distance, alpha)
			src.PrintInfos(s, distance)
		case 3:
			s = functions.CheaperGreedySolution(n, distance)
			src.PrintInfos(s, distance)
		case 4:
			fmt.Println("Not Implemented")
		case 5:
			fmt.Println("Not Implemented")
		}
	}
}

func BestImprovement(n int, distance [][]float64, solution []int) {
	options := []interface{}{
		"Voltar",
		"Aleatório",
		"Melhor vizinho",
	}

	for {
		choice := src.AskOption(src.IMPROVEMENT, options)
		switch choice {
		case 0:
			return
		case 1:
			fo, solution := functions.DescentBestImprovement(n, solution, distance, functions.RandomNeighbor)
			src.PrintInfos(solution, distance)
			fmt.Println("Função Objetivo = ", fo)
		case 2:
			fo, solution := functions.DescentBestImprovement(n, solution, distance, functions.BestNeighbor)
			src.PrintInfos(solution, distance)
			fmt.Println("Função Objetivo = ", fo)
		}
	}
}

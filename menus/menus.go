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
		"Método de Descida",
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
		case 3:
			s, fo := functions.MultiStart(n, solution, distance)
			fmt.Println(s)
			fmt.Println(fo)
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
			s := functions.GreedySolutionNeighborNearby(n, distance, alpha)
			src.PrintInfos(s, distance)
		case 3:
			s = functions.CheaperGreedySolution(n, distance)
			src.PrintInfos(s, distance)
		case 4:
			fmt.Println("Not Implemented")
		case 5:
			s = functions.RandomConstruction(n)
			src.PrintInfos(s, distance)
		}
	}
}

func BestImprovement(n int, distance [][]float64, solution []int) {
	var iterMax int = 1000
	options := []interface{}{
		"Voltar",
		"Descida com Best Improvement",
		"Descida randomica",
		"Descida com Primeiro de Melhora (First Improvement)",
	}

	for {
		choice := src.AskOption(src.IMPROVEMENT, options)
		switch choice {
		case 0:
			return
		case 1:
			start := src.GetTime()
			fo, solution := functions.DescentBestImprovement(n, solution, distance)
			end := src.GetTime()
			src.PrintInfos(solution, distance)
			fmt.Println("Função Objetivo = ", fo)
			fmt.Println("Tempo de execução = ", src.CalculatedTime(start, end))
		case 2:
			start := src.GetTime()
			fo, solution := functions.DescentRandomImprovement(n, solution, distance, iterMax)
			end := src.GetTime()
			src.PrintInfos(solution, distance)
			fmt.Println("Função Objetivo = ", fo)
			fmt.Println("Tempo de execução = ", src.CalculatedTime(start, end))
		case 3:
			fmt.Println("Not Implemented")
		}

	}
}

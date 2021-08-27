package main

import (
	"fmt"

	"Metaheuristica/menus"
	"Metaheuristica/src"
)

func main() {
	// var iterMax int
	var n int
	// var nind int

	// var fo float64
	var best_fit_lit float64
	// var max_desvio float64
	// var prob_crossover float64
	// var prob_mutacao float64
	// var temp float64

	n, best_fit_lit = src.ReadFile()

	fmt.Println(best_fit_lit)
	distance := src.InitializeMatrix(n)
	fmt.Println(len(distance))

	menus.PrincipalMenu(distance, n)

}

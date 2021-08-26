package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

func getTime() time.Time {
	return time.Now()
}

func calculatedTime(start time.Time, end time.Time) time.Duration {
	return end.Sub(start)
}

func stringToFloat(s string) float64 {
	value, _ := strconv.ParseFloat(s, 64)
	return value
}

func stringToInt(s string) int {
	value := stringToFloat(s)
	return int(value)
}

func readFile() (int, float64) {

	file, err := os.Open("teste.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	var text []string

	for scanner.Scan() {
		text = append(text, scanner.Text())
	}

	return stringToInt(text[0]), stringToFloat(text[1])
}

func createMatrix(n int) [][]float64 {
	matrix := make([][]float64, n)

	for i := 0; i < n; i++ {
		matrix[i] = make([]float64, n)
	}

	return matrix
}

func splitStringToTriple(s string) (int, int, int) {
	splitString := strings.Fields(s)
	return stringToInt(splitString[0]), stringToInt(splitString[1]), stringToInt(splitString[2])
}

func initializeMatrix(n int) [][]float64 {
	matrix := createMatrix(n)
	file, err := os.Open("c50.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	vetX := make([]int, n)
	vetY := make([]int, n)

	for scanner.Scan() {
		city, x, y := splitStringToTriple(scanner.Text())
		vetX[city] = x
		vetY[city] = y
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			valueX := float64(vetX[i] - vetX[j])
			valueY := float64(vetX[i] - vetX[j])
			matrix[i][j] = math.Sqrt(math.Pow(valueX, 2) + math.Pow(valueY, 2))
		}
	}

	return matrix
}

func main() {
	// var iterMax int64
	var n int
	// var nind int64

	// var s []int64

	// var alpha float64
	// var fo float64
	var best_fit_lit float64
	// var max_desvio float64
	// var prob_crossover float64
	// var prob_mutacao float64
	// var temp float64

	n, best_fit_lit = readFile()

	fmt.Println(n)
	fmt.Println(best_fit_lit)
	// fmt.Println(distance)
	distance := src.initializeMatrix(n)
	fmt.Println(len(distance))
}

/*

  //srand(1000); // pega o numero 1000 como semente de numeros aleatorios
  srand((unsigned) time(NULL)); // pega a hora do relogio como semente

  le_arq_matriz((char*)"data/c50.txt", n, d);

*/

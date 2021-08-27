package src

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"sync"
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

func ReadFile() (int, float64) {

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

	for i := range matrix {
		matrix[i] = make([]float64, n)
	}

	return matrix
}

func splitStringToTriple(s string) (int, int, int) {
	splitString := strings.Fields(s)
	return stringToInt(splitString[0]), stringToInt(splitString[1]), stringToInt(splitString[2])
}

func InitializeMatrix(n int) [][]float64 {
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

	var wg sync.WaitGroup

	wg.Add((n - 1) * n / 2)
	for i := 0; i < n-1; i++ {
		matrix[i][i] = 0
		for j := i + 1; j < n; j++ {
			go func(w *sync.WaitGroup, matrix [][]float64, i int, j int) {

				valueX := math.Pow(float64(vetX[i]-vetX[j]), 2)
				valueY := math.Pow(float64(vetY[i]-vetY[j]), 2)
				matrix[i][j] = math.Sqrt(valueX + valueY)
				matrix[j][i] = matrix[i][j]
				w.Done()
			}(&wg, matrix, i, j)
		}
	}
	wg.Wait()

	return matrix
}

func AskOption(question string, options []interface{}) (answer int) {

	fmt.Println(":::::::::: " + strings.ToUpper(question) + " ::::::::::")
	fmt.Println()
	for optKey, optVal := range options {
		fmt.Printf("[%d] %s\n", optKey, optVal)
	}
	answer = -1
	for ok := true; ok; ok = (answer < 0 || answer > len(options)) {
		fmt.Printf("\n" + "Escolha: ")
		questionInputScanner := bufio.NewScanner(os.Stdin)
		for questionInputScanner.Scan() {
			if answerChosen, err := strconv.Atoi(questionInputScanner.Text()); err == nil {
				if answerChosen >= 0 && answerChosen < len(options) {
					answer = answerChosen
					break
				}
			}
			fmt.Println("Oops. Answer out of range. Try again.")
			break
		}
	}
	return answer
}

func RemoveElement(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}

func printRoute(s []int) {
	for i := 0; i < len(s); i++ {
		fmt.Print(s[i], " -> ")
	}
	fmt.Println(s[0])
}

func PrintInfos(s []int, distance [][]float64) {
	fo := CalculateOF(s, distance)

	fmt.Println("Solucao obtida usando a estrategia Best Improvement do Metodo da Descida:")
	printRoute(s)
	fmt.Println("Função objetivo = ", fo)
}

func CalculateOF(s []int, distance [][]float64) float64 {
	var route float64 = distance[s[len(s)-1]][s[0]]
	for i := 0; i < len(s)-1; i++ {
		route += distance[s[i]][s[i+1]]
	}

	return route
}

func Unvisited(n int) (unvisited []int) {

	for i := 1; i < n; i++ {
		unvisited = append(unvisited, i)
	}

	return unvisited
}

func InsertPos(v []int, pos int, value int) []int {
	return append(v[:pos], append([]int{value}, v[pos:]...)...)
}

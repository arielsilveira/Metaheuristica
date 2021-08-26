package src

import (
	"bufio"
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

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("Enter your weight in kg: ")
	weight := input()

	fmt.Print("Enter your height in m: ")
	height := input()

	bmi := weight / math.Pow(height, 2)

	if bmi < 18.5 {
		fmt.Printf("You have a bmi of %.2f, you are underweight\n", bmi)
	} else if bmi < 25 {
		fmt.Printf("You have a bmi of %.2f, you are a normal weight\n", bmi)
	} else if bmi < 30 {
		fmt.Printf("You have a bmi of %.2f, you are overweight\n", bmi)
	} else if bmi < 35 {
		fmt.Printf("You have a bmi of %.2f, you are obese\n", bmi)
	} else {
		fmt.Printf("You have a bmi of %.2f, you are clinically obese\n", bmi)
	}
}

// Get input from the terminal
func input() float64 {
	reader := bufio.NewReader(os.Stdin)
	stringValue, _ := reader.ReadString('\n')
	stringValue = strings.TrimSpace(stringValue)
	input, _ := strconv.ParseFloat(stringValue, 64)

	return input
}

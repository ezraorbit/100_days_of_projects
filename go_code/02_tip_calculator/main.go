// Tip calculator
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to the tip calculator")
	fmt.Print("What was the total bill? $")
	totalBill, _ := strconv.ParseFloat(input(), 64)

	fmt.Print("What percentage tip would you like to give? 10, 12, or 15? ")
	tipPercentage, _ := strconv.ParseFloat(input(), 64)
	tipPercentage /= 100

	fmt.Print("How many people to split the bill? ")
	totalPeople, _ := strconv.ParseFloat(input(), 64)

	tipTotal := totalBill * tipPercentage
	totalWithTip := totalBill + tipTotal
	totalToPay := totalWithTip / totalPeople

	fmt.Printf("Each person should pay: $%.2f\n", totalToPay)

}

func input() string {
	reader := bufio.NewReader(os.Stdin)
	input, error := reader.ReadString('\n')
	if error != nil {
		log.Fatal(error)
	}

	input = strings.TrimSpace(input)
	return input
}

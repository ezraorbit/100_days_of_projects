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
	fmt.Println("Which year do you want to check?")
	year := input()

	if year%4 == 0 && year%100 != 0 || year%100 == 0 && year%400 == 0 {
		fmt.Println("The year is a leap year🦘")
	} else {
		fmt.Println("The year is not a leap year🙅‍♂️")
	}

}

func input() int64 {
	reader := bufio.NewReader(os.Stdin)
	newString, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	newString = strings.TrimSpace(newString)
	input, _ := strconv.ParseInt(newString, 10, 64)
	return input
}

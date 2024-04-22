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
	var counterLove int64
	var counterTrue int64
	countTrue := map[string]int64{
		"t": 0,
		"r": 0,
		"u": 0,
		"e": 0,
	}
	countLove := map[string]int64{
		"l": 0,
		"o": 0,
		"v": 0,
		"e": 0,
	}

	fmt.Println("What is your name?")
	name1 := input()

	fmt.Println("What is their name?")
	name2 := input()

	for _, letter := range name1 {
		if _, ok := countLove[string(letter)]; ok {
			countLove[string(letter)]++
		}
		if _, ok := countTrue[string(letter)]; ok {
			countTrue[string(letter)]++
		}
	}
	for _, letter2 := range name2 {
		if _, ok := countLove[string(letter2)]; ok {
			countLove[string(letter2)]++
		}
		if _, ok := countTrue[string(letter2)]; ok {
			countTrue[string(letter2)]++
		}
	}

	for _, value := range countLove {
		counterLove += value
	}
	for _, value := range countTrue {
		counterTrue += value
	}

	trueLoveCount, _ := strconv.ParseInt(fmt.Sprintf("%v%v", counterTrue, counterLove), 10, 64)

	if trueLoveCount < 10 || trueLoveCount > 90 {
		fmt.Printf("Your score is %v, you go together like coke and mentos\n", trueLoveCount)
	} else if trueLoveCount >= 40 && trueLoveCount <= 50 {
		fmt.Printf("Your score is %v, you are alright together\n", trueLoveCount)
	} else {
		fmt.Printf("Your score is %v\n", trueLoveCount)
	}

	// fmt.Printf("This is the map: %v\nThis is the map: %v\n", countLove, countTrue)
	// fmt.Printf("Your score is %v%v\n", counterTrue, counterLove)
}

func input() string {

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	input = strings.TrimSpace(input)
	input = strings.ToLower(input)

	return input
}

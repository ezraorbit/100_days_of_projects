package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func main() {
	secretRow, secretCol := rand.Int63n(3), rand.Int63n(3)

	gameOn := true

	for gameOn {
		treaureHunt(&gameOn, secretRow, secretCol)
	}

}

func treaureHunt(gameOn *bool, secretRow, secretCol int64) {
	row1 := []string{"⬜", "⬜", "⬜"}
	row2 := []string{"⬜", "⬜", "⬜"}
	row3 := []string{"⬜", "⬜", "⬜"}

	gameMap := [][]string{row1, row2, row3}

	fmt.Printf("%v\n%v\n%v\n", row1, row2, row3)

	fmt.Println("Where is the treasure👀\nPlease enter your guess between 0 - 2, e.g '0, 1'")
	row, col := mapInput()

	if row == secretRow && col == secretCol {
		fmt.Println("You win. WOW WOW🥳")
		gameMap[secretRow][secretCol] = "✅"
		fmt.Printf("%v\n%v\n%v\n", row1, row2, row3)
		*gameOn = playAgain()
	} else {
		fmt.Println("Sorry, try again next time🥲")
		gameMap[secretRow][secretCol] = "🪙"
		gameMap[row][col] = "❌"
		fmt.Printf("%v\n%v\n%v\n", row1, row2, row3)
		*gameOn = playAgain()
	}
}

func mapInput() (int64, int64) {
	var input1, input2 int64
	reader := bufio.NewReader(os.Stdin)
	str, _ := reader.ReadString('\n')
	str = strings.TrimSpace(str)
	strSlice := strings.Split(str, ", ")
	for index, value := range strSlice {
		if index == 0 {
			input1, _ = strconv.ParseInt(value, 10, 64)
		} else if index == 1 {
			input2, _ = strconv.ParseInt(value, 10, 64)
		}
	}

	return input1, input2
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

func playAgain() bool {
	fmt.Print("Do you want to play again? type \"y\" or \"n\": ")
	choice := input()
	if choice == "y" {
		return true
	} else if choice == "n" {
		return false
	} else {
		fmt.Println("Please input your answer again🤲")
		return playAgain()
	}
}

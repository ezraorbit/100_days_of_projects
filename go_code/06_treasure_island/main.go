package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Welcome to Treasure Island.\nYour mission is to find the treasure")
	fmt.Println("You're at a cross road. Where do you want to go? Type \"left\" or \"right\"")
	choice := input()
	switch choice {
	case "right":
		fmt.Println("Fell into a hole! Game Over😭")
	case "left":
		fmt.Println("You come to a lake. There is an island in the middle of the lake. Type \"wait\" to wait for a boat. Type \"swim\" to swim across")
		choice = input()
		switch choice {
		case "swim":
			fmt.Println("You were attacked by a trout! Game Over😭")
		case "wait":
			fmt.Println("You arrive at the island unharmed😎. There is a house with 3 doors. One red, one yellow and one blue. Which color do you choose?")
			choice = input()
			switch choice {
			case "red":
				fmt.Println("Burned by fire🔥. Game Over😭")
			case "blue":
				fmt.Println("You were eaten by beasts🦁. Game Over😭")
			case "yellow":
				fmt.Println("You Win🥳🍾")
			default:
				fmt.Println("Something isn't right. Please try again🥲")
			}
		default:
			fmt.Println("Something isn't right. Please try again🥲")
		}
	default:
		fmt.Println("Something isn't right. Please try again🥲")

	}

}

func input() string {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	input = strings.ToLower(input)
	return input
}

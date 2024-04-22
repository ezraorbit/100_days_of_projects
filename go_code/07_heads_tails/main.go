package main

import (
	"fmt"
	"math/rand"
)

func main() {
	coinFlip := rand.Intn(2)

	if coinFlip == 1 {
		fmt.Println("Heads")
	} else {
		fmt.Println("Tails")
	}
}

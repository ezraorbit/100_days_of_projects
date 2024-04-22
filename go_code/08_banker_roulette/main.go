package main

import (
	"fmt"
	"math/rand"
)

func main() {
	names := []string{"Timothy", "Greggory", "Dawson", "Philippe", "Kai", "Martinelli"}
	randomChoice := rand.Intn(len(names))
	fmt.Printf("Hey %s, you are paying for everyone's meal today.\n", names[randomChoice])

}

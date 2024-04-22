// A project that generates a bnadname based of a city name and pet name
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	fmt.Println("What's the name of the city you grew up in?")
	cityName := takeInput()

	fmt.Println("What's your pet's name?")
	petName := takeInput()

	fmt.Println(cityName, petName)
}

func takeInput() string {
	reader := bufio.NewReader(os.Stdin)
	input, error := reader.ReadString('\n')
	if error != nil {
		log.Fatal(error)
	}

	input = strings.TrimSpace(input)
	return input
}

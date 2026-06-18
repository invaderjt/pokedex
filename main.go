package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		if !scanner.Scan() {
			fmt.Printf("Invalid input: %v\n", scanner.Err())
			continue
		}
		input := cleanInput(scanner.Text())[0]
		if _, ok := commands[input]; ok {
			commands[input].callback()
			continue
		} else {
			fmt.Println("Unknown command")
		}
	}
}

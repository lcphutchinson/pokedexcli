package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	cli := bufio.NewScanner(os.Stdin)
	commandMap := getCommandMap()
	fmt.Print("Pokedex > ")
	for cli.Scan() {
		command, ok := commandMap[cli.Text()]
		if ok {
			command.callback()
		} else {
			fmt.Printf("Error: invalid command \"%v\"\n", cli.Text())
		}
		fmt.Print("Pokedex > ")
	}
}

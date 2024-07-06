package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"time"
	"github.com/lcphutchinson/pokecache"
)

const cacheInterval time.Duration = 30 * time.Second

func main() {
	cli := bufio.NewScanner(os.Stdin)
	commandMap := getCommandMap()
	cache := pokecache.NewCache(cacheInterval)
	config := Config{
		cache:	&cache,
		nxtMap:	"https://pokeapi.co/api/v2/location-area/",
		prvMap:	"",
	}
	fmt.Print("Pokedex > ")
	for cli.Scan() {
		inputs := strings.Split(cli.Text(), " ")
		command, ok := commandMap[inputs[0]]
		if !ok {
			fmt.Printf("Error: invalid command \"%v\"\n", inputs[0])
			continue
		}
		err := command.callback(&config, inputs)
		if err != nil {
			fmt.Printf("Error in %v: %v\n", command.name, err)
		}
		fmt.Print("Pokedex > ")
	}
}

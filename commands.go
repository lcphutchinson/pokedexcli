package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"github.com/lcphutchinson/pokecache"
	"github.com/lcphutchinson/pokecaller"
	"github.com/lcphutchinson/pokejson"
	"github.com/lcphutchinson/pokedex"
)

type cliCommand struct {
	name		string
	description	string
	callback	func(c *Config, params []string) error
}

type Config struct {
	cache	*pokecache.Cache
	dex	*pokedex.Pokedex
	nxtMap	string
	prvMap	string
}

func getCommandMap() map[string]cliCommand {
	return map[string]cliCommand {
		"exit": {
			name:		"exit",
			description:	"Exits the session",
			callback:	func(c *Config, params []string) error {
				fmt.Println("Thank you for using PokedexCLI!")
				os.Exit(0)
				return nil
			},
		},
		"help": {
			name:		"help",
			description:	"Displays a help message",
			callback:	func(c *Config, params []string) error {
				fmt.Println("No help subcommand structure yet")
				return nil
			},
		},
		"map": {
			name:		"map",
			description:	"fetches a new page of location areas from PokeAPI",
			callback:	func(c *Config, params []string) error {
				if c.nxtMap == "" {
					return fmt.Errorf("No more areas to list")
				}
				fmt.Println("Fetching some locations...")
				lst := pokejson.AreaQuery{}
				res, ok := c.cache.Get(c.nxtMap)
				var err error = nil
				if !ok {
					res, err = pokecaller.Call(c.nxtMap)
					if err != nil {
						return err
					}
					c.cache.Add(c.nxtMap, res)
				}
				err = json.Unmarshal(res, &lst)
				if err != nil {
					return err
				}
				for _, area := range lst.Results {
					fmt.Println(area.Name)
				}
				c.prvMap = c.nxtMap
				c.nxtMap = lst.Next
				return nil
			},
		},
		"mapb": {
			name:		"mapb",
			description:	"fetches the previous page of location areas",
			callback:	func(c *Config, params []string) error {
				if c.prvMap == "" {
					return fmt.Errorf("No more areas to list")
				}
				fmt.Println("Fetching some locations...")
				lst := pokejson.AreaQuery{}
				res, ok := c.cache.Get(c.prvMap)
				var err error = nil
				if !ok {
					res, err = pokecaller.Call(c.prvMap)
					if err != nil {
						return err
					}
					c.cache.Add(c.prvMap, res)
				}
				err = json.Unmarshal(res, &lst)
				if err != nil {
					return err
				}
				for _, area := range lst.Results {
					fmt.Println(area.Name)
				}
				c.nxtMap = c.prvMap
				c.prvMap = lst.Previous
				return nil
			},
		},
		"explore": {
			name:		"explore",
			description:	"explores a named map area and returns pokemon found there",
			callback:	func (c *Config, params []string) error {
				if len(params) < 2 {
					return fmt.Errorf("Command requires a location argument")
				}
				query := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%v", params[1])
				area := pokejson.Area{}
				res, ok := c.cache.Get(query)
				var err error = nil
				if !ok {
					res, err = pokecaller.Call(query)
					if err != nil {
						return err
					}
					c.cache.Add(query, res)
				}
				err = json.Unmarshal(res, &area)
				fmt.Printf("Exploring %v...\n", area.Name)
				for _, encounter := range area.Encounters {
					fmt.Println(encounter.Pokemon.Name)
				}
				return nil
			},
		},
		"catch": {
			name:		"catch",
			description:	"attempts to catch a named pokemon",
			callback:	func (c *Config, params []string) error {
				if len(params) < 2 {
					return fmt.Errorf("Command requires a pokemon argument")
				}
				query := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%v", params[1])
				pokemon := pokejson.Pokemon{}
				res, ok := c.cache.Get(query)
				var err error = nil
				if !ok {
					res, err = pokecaller.Call(query)
					if err != nil {
						return err
					}
					c.cache.Add(query, res)
				}
				err = json.Unmarshal(res, &pokemon)
				fmt.Println("You used Pokeball!")
				chance := pokemon.BaseEXP
				if chance > 350 {
					chance = 350
				}
				for i := 0; i < 3; i++ {
					if rand.Intn(750) < chance {
						fmt.Printf("%v broke free!\n", pokemon.Name)
						return nil
					}
					fmt.Println("...")
				}
				if rand.Intn(750) < pokemon.BaseEXP {
					fmt.Printf("%v broke free!\n", pokemon.Name)
					return nil
				}
				fmt.Printf("%v was caught!\n", pokemon.Name)
				ok = c.dex.Add(pokemon)
				if !ok {
					return nil
				}
				fmt.Printf("added %v's data to the Pokedex\n", pokemon.Name)
				return nil
			},
		},
		"inspect": {
			name:		"inspect",
			description:	"view a pokemon's data in the pokedex",
			callback:	func (c *Config, params []string) error {
				if len(params) < 2 {
					return fmt.Errorf("Command requires a pokemon argument")
				}
				data, ok := c.dex.Get(params[1])
				if !ok {
					fmt.Println("you have not caught that pokemon")
					return nil
				}
				fmt.Printf("Name: %v\n", data.Name)
				fmt.Printf("Height: %v\n", data.Height)
				fmt.Printf("Weight: %v\n", data.Weight)
				fmt.Println("Stats:")
				for _, stat := range data.Stats {
					fmt.Printf("\t-%v: %v\n", stat.Label.Name, stat.Value)
				}
				fmt.Println("Types:")
				for _, pType := range data.Types {
					fmt.Printf("\t-%v\n", pType.Label.Name)
				}
				return nil
			},
		},
		"pokedex": {
			name:		"pokedex",
			description:	"displays the user's caught pokemon",
			callback:	func (c *Config, params []string) error {
				fmt.Println("Your Pokedex: ")
				for _, mon := range c.dex.List() {
					fmt.Printf(" - %v\n", mon)
				}
				return nil
			},
		},
	}
}



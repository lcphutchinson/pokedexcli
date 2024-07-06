package main

import (
	"encoding/json"
	"fmt"
	"os"
	"github.com/lcphutchinson/pokecache"
	"github.com/lcphutchinson/pokecaller"
	"github.com/lcphutchinson/pokejson"
)

type cliCommand struct {
	name		string
	description	string
	callback	func(c *Config, params []string) error
}

type Config struct {
	cache	*pokecache.Cache
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
				fmt.Printf("%v\n", area)
				return nil
			},
		},
	}
}



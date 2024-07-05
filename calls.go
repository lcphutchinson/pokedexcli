package main

import (
	"fmt"
	"io"
	"net/http"
	"encoding/json"
)

type result struct {
	Name	string
}

type batch struct {
	Results	[]result
	Next	*string	`json:"next"`
	Prev	*string	`json:"irevious"`
}

func callForMaps() error {
	res, err := http.Get("https://pokeapi.co/api/v2/location-area/")
	if err != nil {
		return err
	}
	if res.StatusCode > 299 {
		return fmt.Errorf("Get Request failed with Status '%v' \n", res.Status)
	}

	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return err
	}

	data := batch{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return err
	}

	for i := 0; i < 20; i++ {
		fmt.Println(data.Results[i])
	}
	return nil	
}

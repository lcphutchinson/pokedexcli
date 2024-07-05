package pokecaller

import (
	"fmt"
	"io"
	"net/http"
)

func Call(url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if res.StatusCode > 299 {
		return nil, fmt.Errorf("Get Request failed with Status '%v' \n", res.Status)
	}

	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return nil, err
	}
	return body, nil
}

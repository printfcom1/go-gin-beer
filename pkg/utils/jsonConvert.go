package utils

import "encoding/json"

type BeerData struct {
	BeerName        string `json:"beer_name"`
	BeerType        string `json:"beer_type"`
	BeerDescription string `json:"beer_description"`
}

func JsonConvert(data string) (*BeerData, error) {
	var beerData BeerData
	err := json.Unmarshal([]byte(data), &beerData)
	if err != nil {
		return nil, err
	}
	return &beerData, nil
}

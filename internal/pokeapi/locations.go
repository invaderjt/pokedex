package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

type Locations struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	} `json:"results"`
}

func (c *Client) LocationsList(pageURL *string) (Locations, error) {
	url := baseURL + "/location-area"

	if pageURL != nil {
		url = *pageURL
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Locations{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Locations{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return Locations{}, err
	}

	locations := Locations{}
	err = json.Unmarshal(data, &locations)
	if err != nil {
		return Locations{}, err
	}

	return locations, nil
}

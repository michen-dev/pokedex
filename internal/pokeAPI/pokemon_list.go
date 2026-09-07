package pokeAPI

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)




func (c *Client) Get_pokemon_list(location_area string) (Pokemon_API_Response, error) {
	url := baseURL + "/location-area" + "/" + location_area + "/"

	if entry, ok := c.cache.Get(url); ok {
		var data Pokemon_API_Response
		err := json.Unmarshal(entry, &data)
		if err != nil {
			return Pokemon_API_Response{}, err
		}
		return data, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Pokemon_API_Response{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon_API_Response{}, err
	}
	defer res.Body.Close()

	if res.StatusCode == http.StatusNotFound {
		return Pokemon_API_Response{}, fmt.Errorf("Location area does not exist")
	}

	dat, err := io.ReadAll(res.Body)
	if err != nil {
		return Pokemon_API_Response{}, err
	}
	c.cache.Add(url, dat)

	var pokemon_list Pokemon_API_Response
	if err := json.Unmarshal(dat, &pokemon_list); err != nil {
		return Pokemon_API_Response{}, nil
	}

	c.cache.Add(url, dat)
	return pokemon_list, nil
}
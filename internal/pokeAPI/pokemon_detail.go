package pokeAPI

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)


func (c *Client) Get_pokemon_detail(name string) (Poke_Detail, error) {
	url := baseURL + "/pokemon" + "/" + name + "/"

	if entry, ok := c.cache.Get(url); ok {
		var poke_detail Poke_Detail
		err := json.Unmarshal(entry, &poke_detail)
		if err != nil {
			return Poke_Detail{}, err
		}
		return poke_detail, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Poke_Detail{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return Poke_Detail{}, err
	}
	defer res.Body.Close()
	
	if res.StatusCode != http.StatusOK {
		return Poke_Detail{}, fmt.Errorf("404 Request failed")
	}

	dat, err := io.ReadAll(res.Body)
	if err != nil {
		return Poke_Detail{}, err
	}

	var poke_detail Poke_Detail
	if err := json.Unmarshal(dat, &poke_detail); err != nil {
		return Poke_Detail{}, err
	}

	c.cache.Add(url, dat)
	return poke_detail, nil
}
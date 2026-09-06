package pokeAPI

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) Get_location_areas(page_url *string) (LA_API_Response, error) {
	url := baseURL + "/location-area"
	if page_url != nil {
		url = *page_url
	}

	if entry, ok := c.cache.Get(url); ok {
		var data LA_API_Response
		err := json.Unmarshal(entry, &data)
		if err != nil {
			return LA_API_Response{}, err
		}
		return data, nil
	}

	
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {		
		return LA_API_Response{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return LA_API_Response{}, err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return LA_API_Response{}, fmt.Errorf("Request failed with status code: %s", res.Status)
	}

	dat, err := io.ReadAll(res.Body)
	if err != nil {
		return LA_API_Response{}, err
	}

	var data LA_API_Response
	err = json.Unmarshal(dat, &data)
	if err != nil {
		return LA_API_Response{}, err
	}

	c.cache.Add(url, dat)
	return data, nil
}
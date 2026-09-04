package pokeAPI

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (c *Client) Get_location_areas(page_url *string) (LA_API_Response, error) {
	url := baseURL + "/location-area"
	if page_url != nil {
		url = *page_url
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

	var data LA_API_Response
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&data); err != nil {
		return data, err
	}
	return data, nil
}
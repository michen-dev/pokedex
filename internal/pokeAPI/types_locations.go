package pokeAPI

type Location_Area struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type LA_API_Response struct {
	Count    int             `json:"count"`
	Next     *string         `json:"next"`
	Previous *string         `json:"previous"`
	Results  []Location_Area `json:"results"`
}

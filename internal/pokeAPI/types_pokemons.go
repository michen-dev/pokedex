package pokeAPI


type Pokemon struct {
	Name string	`json:"name"`
	Url string `json:"url"`
}

type Encounter struct {
	Pokemon Pokemon	`json:"pokemon"`
}

type Pokemon_API_Response struct {
	Pokemon_encounters []Encounter	`json:"pokemon_encounters"`
}
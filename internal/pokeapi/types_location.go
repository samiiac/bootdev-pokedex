package pokeapi

type location struct{
	Name string `"json:name"`
	URL  string `"json:url"`
}

type LocationAreaResponse struct{
	Count int  `"json:count"`
	Next  string  `"json:next"`
	Previous string  `"json:previous"`
	Results  []location `"json:results"`
}


type ExploreLocationAreaResponse struct {
	Id              int        `json:"id"`
	Name            string     `json:"name"`
	GameIndex       string      `"json:"game_index"`
	PokemonEnc      []struct {
		            Pokemon struct {
		Name     string `json:"name"`
		URL      string  `json:"url"`
		} `json:"pokemon"`
	}  `json:"pokemon_encounters"`
	 
}



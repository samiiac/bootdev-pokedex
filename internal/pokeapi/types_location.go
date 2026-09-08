package pokeapi

type location struct{
	Name string `"json:name"`
	URL  string `"json:url"`
}

type locationAreaResponse struct{
	Count int  `"json:count"`
	Next  string  `"json:next"`
	Previous string  `"json:previous"`
	Results  []location `"json:results"`
}
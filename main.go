package main 
import ("github.com/samiiac/bootdev-pokedex/internal/pokeapi")

type config struct{
	commandRegistry map[string]command
	previous        string
	next            string
	caughtPokemons map[string]pokeapi.Pokemon 
}



func main(){
	c :=  config{
	commandRegistry : getCommands(),
	previous : "",
	next : "",
	caughtPokemons:make(map[string]pokeapi.Pokemon),

 }

	startRepl(&c)
}
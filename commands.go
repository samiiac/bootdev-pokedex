package  main

import ("fmt" 
"os"
)


type command struct {
	name string
	description string
	callback func(c *config,args string) error
}



func getCommands()  map[string]command {
	return map[string]command{
	"exit":{
		name:"exit",
		description:"Exit the pokedex",
		callback:exit,
},
   "help":{
		name:"help",
		description:"Displays a help message.",
		callback:help,
},
  "map":{
	    name:"map",
		description:"Displays the names of 20 location areas in the Pokemon world.",
		callback:listLocation,
  },
  "mapb":{
        name:"mapb",
		description:"Displays the previous names of 20 location areas in the Pokemon world.",
		callback:listPrevLocation,
  },
  "explore":{
        name:"explore <location-name>",
		description:"Displays the section of areas for the location name typed.",
		callback:exploreLocation,
  },
  "catch":{
	    name:"catch <pokemon_name>",
		description:"Displays whether the pokemon was caught or not.",
		callback:catch,
  },
  "inspect":{
	    name:"inspect <pokemon_name>",
		description:"Displays the detail of the pokemon you  only if, you've caught it.",
		callback:inspect,
  },
   "pokedex":{
	    name:"pokedex",
		description:"Displays all the pokemons you've caught ;))",
		callback:listCaughtPokemons,
  },
	}
}

func exit(c *config,args string) error {
 fmt.Println("Closing the Pokedex... Goodbye!")
 os.Exit(0) 
return nil

} 

func help(c *config,args string) error {

 fmt.Println("Welcome to the Pokedex!\nUsage:")
 for _,desc := range c.commandRegistry {
   fmt.Printf("%v : %v \n",desc.name,desc.description)
 }
 return nil
 

}







package  main

import ("fmt" 
"os"
"github.com/samiiac/bootdev-pokedex/internal/pokeapi"
)


type command struct {
	name string
	description string
	callback func(c *config) error
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
		description:"Displays a help message",
		callback:help,
},
  "map":{
	    name:"map",
		description:"Displays the names of 20 location areas in the Pokemon world",
		callback:listLocation,
  },
  "mapb":{
        name:"mapb",
		description:"Displays the previous names of 20 location areas in the Pokemon world",
		callback:listPrevLocation,
  },
	}
}

func exit(c *config) error {
 fmt.Println("Closing the Pokedex... Goodbye!")
 os.Exit(0) 
return nil

} 

func help(c *config) error {

 fmt.Println("Welcome to the Pokedex!\nUsage:")
 for command,desc := range c.commandRegistry {
   fmt.Printf("%v : %v \n",command,desc.description)
 }
 return nil
 

}


func listPrevLocation(c *config) error{
	url := c.previous
	if url == ""{
		return fmt.Errorf("No previous location areas found")
	}
	
   res,err := pokeapi.fetchLocation(url)

   if err != nil {
    return err
   }


   c.next = res.Next;
   c.previous = res.Previous;

  
   for _,l := range response.Results {
	fmt.Println(l.Name)
   }

   return nil
}

func listLocation(c *config) error{
	url := "https://pokeapi.co/api/v2/location-area"
	if c.next != "" {
		url  = c.next
	}
   res,err := pokeapi.fetchLocation(url)

   if err != nil {
    return err
   }

   c.next = res.Next;
   c.previous = res.Previous;

  
   for _,l := range res.Results {
	fmt.Println(l.Name)
   }

   return nil
}



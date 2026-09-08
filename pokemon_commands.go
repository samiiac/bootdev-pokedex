package main
import ("fmt" 
"github.com/samiiac/bootdev-pokedex/internal/pokeapi"
"math/rand"
)

func catch(c *config,args string) error {
	if args == "" {
		return fmt.Errorf("Name of the pokemon is required")
	}
	url:= "https://pokeapi.co/api/v2/pokemon/"+args+"/"

	fmt.Printf("Throwing a Pokeball at %v...\n",args)
    res,err := pokeapi.Fetch[pokeapi.Pokemon](url)
    if err != nil {
    return err
    }
	threshold := 75
	chance := rand.Intn(1+res.BaseExperience)
	if chance > threshold {
		fmt.Printf("%v was caught!\n",args)
		c.caughtPokemons[args] = res
	} else{
		fmt.Printf("%v escaped!\n",args)
	}

	return nil

  
}

func inspect(c *config,args string) error{
  if args == "" {
		return fmt.Errorf("Name of the pokemon is required")
	}
    pokemon,ok := c.caughtPokemons[args]
	if !ok{
		return fmt.Errorf("You have not caught "+ args + " yet")
	}

	fmt.Printf("Name: %v\nHeight: %v\nWeight: %v\nStats:",pokemon.Name,pokemon.Height,pokemon.Weight)

    for _,pstat := range pokemon.Stats{
		fmt.Printf("\n\t-%v: %v",pstat.Stat.Name,pstat.BaseStat)
	}
    fmt.Printf("\nTypes:")
	for _,ptype := range pokemon.Types{
		fmt.Printf("\n\t- %v",ptype.Type.Name)
	}
	fmt.Printf("\n")

	return nil
}


func listCaughtPokemons(c *config,args string)error {
   if len(c.caughtPokemons) == 0 {
	return fmt.Errorf("No pokemons caught!")
   }
   fmt.Println("Your Pokedex: ")
   for _,pokemon := range c.caughtPokemons{
	fmt.Printf("- %v\n",pokemon.Name)
   }

   return nil
}

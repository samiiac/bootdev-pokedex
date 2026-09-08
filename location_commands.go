package main
import ("fmt" 

"github.com/samiiac/bootdev-pokedex/internal/pokeapi"

)

func listPrevLocation(c *config, args string) error{
	url := c.previous
	if url == ""{
		return fmt.Errorf("No previous location areas found")
	}
	
   res,err := pokeapi.Fetch[pokeapi.LocationAreaResponse](url)

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

func listLocation(c *config,args string) error{
	url := "https://pokeapi.co/api/v2/location-area"
	if c.next != "" {
		url  = c.next
	}
   res,err := pokeapi.Fetch[pokeapi.LocationAreaResponse](url)

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

func exploreLocation(c *config,args string) error{
	
	if args == "" || len(args) <= 0 {
		return fmt.Errorf("location name is required")
	}

	url := "https://pokeapi.co/api/v2/location-area/"+args

	res,err := pokeapi.Fetch[pokeapi.ExploreLocationAreaResponse](url)
 
   if err != nil {
    return err
   }
  
   for _,encounter := range res.PokemonEnc {
	fmt.Println(encounter.Pokemon.Name)
   }

   return nil

}
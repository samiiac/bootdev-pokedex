package pokeapi 
import (
	"github.com/samiiac/bootdev-pokedex/internal/pokecache"
	"time"
)


var cache = pokecache.NewCache(5 * time.Second)




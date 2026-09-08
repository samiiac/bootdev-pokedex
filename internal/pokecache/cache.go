package pokecache
import ("time"
"sync")

type Cache struct {
  cacheMap  map[string]CacheEntry
  mu        *sync.Mutex
}


type CacheEntry struct {
	createAt  time.Time
	val        []byte
}

func NewCache(interval time.Duration) {
	newCache := Cache{
		cacheMap : make(map[string]CacheEntry),
		mu : &sync.Mutex{},
	}
	newCache.reapLoop(interval * time.Second)
}

func (cache Cache) Add(key string,val []byte) {
    cache.mu.Lock()
	defer cache.mu.Unlock()
    entry := CacheEntry{
		createAt:time.Now(),
		val,
	}
	cache.cacheMap[key] = entry
}

func (cache Cache) Get(key string) ([]byte,bool) {
    cache.mu.Lock()
	defer cache.mu.Unlock()
	if entry,ok := cache.cacheMap[key]; !ok{
		return []byte{},false
	}else{
      return entry.val,true
	}
	
}

func (cache Cache) reapLoop(interval time.Duration) {
	 
     ticker := time.NewTicker(interval)
	 go for currTime := <-ticker.C {
		cache.mu.Lock()
	    defer cache.mu.Unlock() //confused
		for key,entry := range cache.cacheMap {
            if entry.createdAt < currTime {
			 delete(cache.cacheMap,key)
			}
		}
	 }
}



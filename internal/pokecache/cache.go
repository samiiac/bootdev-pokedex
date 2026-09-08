package pokecache
import ("time"
"sync")

type Cache struct {
  cacheMap  map[string]CacheEntry
  mu        *sync.Mutex
}


type CacheEntry struct {
	createdAt  time.Time
	val        []byte
}

func NewCache(interval time.Duration) Cache{
	newCache := Cache{
		cacheMap : make(map[string]CacheEntry),
		mu : &sync.Mutex{},
	}
	newCache.reapLoop(interval)
	return newCache
}

func (cache *Cache) Add(key string,val []byte) {
    cache.mu.Lock()
	defer cache.mu.Unlock()
    entry := CacheEntry{
		createdAt:time.Now(),
		val:val,
	}
	cache.cacheMap[key] = entry
}

func (cache *Cache) Get(key string) ([]byte,bool) {
    cache.mu.Lock()
	defer cache.mu.Unlock()
	if entry,ok := cache.cacheMap[key]; !ok{
		return []byte{},false
	}else{
      return entry.val,true
	}
	
}

func (cache *Cache) reapLoop(interval time.Duration) {
	 
     ticker := time.NewTicker(interval)
	 go func(){
	  for currTime := range ticker.C {
		cache.mu.Lock()
	     
		for key,entry := range cache.cacheMap {
            if currTime.Sub(entry.createdAt) > interval {
			 delete(cache.cacheMap,key)
			}
		}
		cache.mu.Unlock() 
	 }}()
}



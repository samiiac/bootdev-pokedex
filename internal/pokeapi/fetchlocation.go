package pokeapi
import (
"net/http"
"encoding/json"
"io"

)


func Fetch[T any](url string) (T,error) {
	bytes := []byte{}
    var response T
	if dataFromCache, ok := cache.Get(url);ok {
		bytes = dataFromCache
	}else {
	res,err := http.Get(url)
   
	if err != nil {
		return response,err
	}

	defer res.Body.Close()


	data,err := io.ReadAll(res.Body)

		if err != nil {
		return response,err
	    }
   cache.Add(url,data)
   bytes = data

 }
   
   if err := json.Unmarshal(bytes,&response); err != nil{
	return response,err
   }
 
   return response,nil

}



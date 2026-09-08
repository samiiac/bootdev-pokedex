package pokeapi
import (
"net/http"
"encoding/json"
"io"
)


func fetchLocation(url string) locationAreaResponse{},error {
	bytes := []byte{}

	if dataFromCache, ok := cache.Get(url);ok {
		bytes = dataFromCache
	}else {
	res,err := http.Get(url)

	if err != nil {
		return nil,err
	}

	defer res.Body.Close()


	data,err := io.ReadAll(res.Body)

		if err != nil {
		return nil,err
	    }
   cache.Add(url,data)
   bytes = data

 }
   response :=  locationAreaResponse{}
   if err := json.Unmarshal(data,&response); err != nil{
	return nil,err
   }

   return response,nil

}

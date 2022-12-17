package main

import (
	"strings"
	"net/http"
	"os"
	"io"
	"encoding/json"
)

func shorten(inputUrl string)string{
	const myUrl = "https://reduced.to/api/v1/shortener"

	var request string = "{ \"originalUrl\" : \"" + inputUrl + "\" }";

	requestBody := strings.NewReader(request)


	response, err := http.Post(myUrl,"application/json",requestBody)


	if !(err==nil) {
		Box(RedText+"Invalid URL"+NormalText)
		os.Exit(1)
	}

	defer response.Body.Close()

	content , _ :=io.ReadAll(response.Body)

	data := []byte(content)

	var dat map[string]interface{}

	if err := json.Unmarshal(data, &dat); err != nil {
        panic(err)
    }
	url := "https://reduced.to/"  + dat["newUrl"].(string)
	
	Box(url)

	return url
}

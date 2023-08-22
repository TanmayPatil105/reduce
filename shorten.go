package main

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"regexp"
	"strings"
)

func shorten(inputUrl string, result chan string) {

	const myUrl = "https://reduced.to/api/v1/shortener"

	var request string = "{ \"originalUrl\" : \"" + inputUrl + "\" }"

	requestBody := strings.NewReader(request)

	response, err := http.Post(myUrl, "application/json", requestBody)

	if !(err == nil) {
		Box(RedText+"Invalid URL"+NormalText, "❌")
		os.Exit(1)
	}

	defer response.Body.Close()

	content, _ := io.ReadAll(response.Body)

	data := []byte(content)

	var dat map[string]interface{}

	if err := json.Unmarshal(data, &dat); err != nil {
		Box(RedText+"Couldn't Parse JSON"+NormalText, "✂️")
	}

	short := dat["newUrl"]

	if short == nil {
		Box(RedText+"Invalid URL"+NormalText, "❌")
		os.Exit(1)
	}

	url := "https://reduced.to/" + dat["newUrl"].(string)

	Box(url, "URL")

	result <- url

}

func checkURL(inputUrl string) {

	if regx, _ := regexp.MatchString("https://reduced.to", inputUrl); regx == true {
		Box(YellowText+"URL is already shortened"+NormalText, "🙌")
		os.Exit(1)
	}

	res, err := http.Get(inputUrl)

	if err != nil {

		_, netError := http.Get("https://www.google.com")

		if netError != nil {
			Box("No Internet", "🦖")
			os.Exit(1)
		}

		Box(RedText+"Invalid URL"+NormalText, "❌")
		os.Exit(1)

	}

	if res.StatusCode != 200 {
		Box(RedText+"Not reachable"+NormalText, "🚫")
		os.Exit(1)
	}

}

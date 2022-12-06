package main

import (
	"fmt"
	"os"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
	"github.com/Delta456/box-cli-maker/v2"
	"github.com/atotto/clipboard"
	// "github.com/pborman/options"
)


const (
	NormalText       = "\033[0m" 
	BlackText        = "\033[30m"
	RedText          = "\033[31m"
	GreenText        = "\033[32m"
	YellowText       = "\033[33m"
	BlueText         = "\033[34m"
	MagentaText      = "\033[35m"
	CyanText         = "\033[36m"
	WhiteText        = "\033[37m"
	DefaultColorText = "\033[39m" 
	BoldText         = "\033[1m"
	BoldBlackText    = "\033[1;30m"
	BoldRedText      = "\033[1;31m"
	BoldGreenText    = "\033[1;32m"
	BoldYellowText   = "\033[1;33m"
	BoldBlueText     = "\033[1;34m"
	BoldMagentaText  = "\033[1;35m"
	BoldCyanText     = "\033[1;36m"
	FaintText        = "\033[2m"
	FaintBlackText   = "\033[2;30m"
	FaintRedText     = "\033[2;31m"
	FaintGreenText   = "\033[2;32m"
	FaintYellowText  = "\033[2;33m"
	FaintBlueText    = "\033[2;34m"
	FaintMagentaText = "\033[2;35m"
	FaintCyanText    = "\033[2;36m"
	FaintWhiteText   = "\033[2;37m"
	DefaultText      = "\033[22;39m" 
)

func main(){
	printAscii()

	var inputUrl string
	if(len(os.Args)==1){
		fmt.Printf("\nEnter the URL : " + GreenText) 
		fmt.Scanln(&inputUrl)
		fmt.Printf("\n") 
	} else if(os.Args[1] == "--clip" || os.Args[1] == "-c" ){
		inputUrl,_ = clipboard.ReadAll()
	}
	postReq(inputUrl)

}

func Box(url string){
	Box := box.New(box.Config{Px: 5, Py: 2, Type: "Round",TitlePos:"Top", TitleColor:"Green",Color: "Cyan"})
 	Box.Print("URL",url)
}

func printAscii(){
	fmt.Println(BoldGreenText)
	fmt.Println("                    __                        __     ")
	fmt.Println("   _____ ___   ____/ /__  __ _____ ___   ____/ /     ")
	fmt.Println("  / ___// _ \\ / __  // / / // ___// _ \\ / __  /    ")
	fmt.Println(" / /   /  __// /_/ // /_/ // /__ /  __// /_/ /     ")
	fmt.Println("/_/    \\___/ \\__ _/ \\__ _/ \\___/ \\___/ \\__ _/  ")
	fmt.Println(NormalText)
}

func checkUrl(url string)bool{
    resp,err := http.Get(url)
    if err != nil {
        return false
    }
	fmt.Println(resp.StatusCode)
    return resp.StatusCode == 200
}

func postReq(inputUrl string){
	const myUrl = "https://reduced.to/api/v1/shortener"

	var request string = "{ \"originalUrl\" : \"" + inputUrl + "\" }";

	requestBody := strings.NewReader(request)


	response, err := http.Post(myUrl,"application/json",requestBody)


	if !(err==nil) {
		// Box(RedText+"Invalid URL"+NormalText)
		return
	}

	defer response.Body.Close()

	content , _ :=ioutil.ReadAll(response.Body)

	data := []byte(content)

	var dat map[string]interface{}

	if err := json.Unmarshal(data, &dat); err != nil {
        panic(err)
    }
	url := "https://reduced.to/"  + dat["newUrl"].(string)

	clipboard.WriteAll(url)
	
	Box(url)


}

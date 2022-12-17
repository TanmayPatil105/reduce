package main

import (
	"fmt"
	"github.com/Delta456/box-cli-maker/v2"
)

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


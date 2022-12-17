package main

import (
	"fmt"
	"flag"
	"os"
	"github.com/atotto/clipboard"
)

const CurrentVersion = "0.0.3"

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

	clip := flag.Bool("b", false, "copy from clipboard")
	copy := flag.Bool("c", false, "copy to clipboard")
	ascii := flag.Bool("d", false,"display ascii art")
	showVersion := flag.Bool("v",false,"show version")
	showHelp := flag.Bool("h",false,"show help")
	flag.Parse()

	var inputUrl string

	if *showVersion{
		fmt.Printf("reduced version %s\n", CurrentVersion)
		os.Exit(0)
	}

	if *showHelp {
		help()
		os.Exit(0)
	}

	if ! *ascii {
		printAscii()
	}

	if *clip {
		inputUrl,_ = clipboard.ReadAll()
	} else {
		fmt.Printf("\nEnter the URL : " + GreenText) 
		fmt.Scanln(&inputUrl)
		fmt.Printf("\n")
	}

	url := shorten(inputUrl)

	if *copy {
		clipboard.WriteAll(url)
	}
}


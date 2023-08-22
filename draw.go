package main

import (
	"fmt"
	"github.com/Delta456/box-cli-maker/v2"
	"math/rand"
	"time"
)

func Box(url string, title string) {
	Box := box.New(box.Config{Px: 5, Py: 2, Type: "Round", TitlePos: "Top", TitleColor: "Green", Color: "Cyan"})
	Box.Print(title, url)
}

func printAscii() {
	fmt.Println(selectRandom())
	fmt.Println("                    __                        __     ")
	fmt.Println("   _____ ___   ____/ /__  __ _____ ___   ____/ /     ")
	fmt.Println("  / ___// _ \\ / __  // / / // ___// _ \\ / __  /    ")
	fmt.Println(" / /   /  __// /_/ // /_/ // /__ /  __// /_/ /     ")
	fmt.Println("/_/    \\___/ \\__ _/ \\__ _/ \\___/ \\___/ \\__ _/  ")
	fmt.Println(NormalText)
}

func selectRandom() string {
	colors := make([]string, 0)
	colors = append(colors,
		NormalText,
		BlackText,
		RedText,
		GreenText,
		YellowText,
		BlueText,
		MagentaText,
		CyanText,
		WhiteText,
		DefaultColorText,
		BoldText,
		BoldBlackText,
		BoldRedText,
		BoldGreenText,
		BoldYellowText,
		BoldBlueText,
		BoldMagentaText,
		BoldCyanText,
		FaintText,
		FaintBlackText,
		FaintRedText,
		FaintGreenText,
		FaintYellowText,
		FaintBlueText,
		FaintMagentaText,
		FaintCyanText,
		FaintWhiteText,
		DefaultText)

	rand.Seed(time.Now().Unix())
	return colors[rand.Intn(len(colors))]

}

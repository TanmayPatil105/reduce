package main

import "fmt"

func help(){
	fmt.Printf("reduce is a command line URL shortener tool\n\n")
	fmt.Printf("Usage: \n")
	fmt.Printf("\t\treduce [arguments]")
	fmt.Printf("\nThe commands are:\n\n")
	fmt.Printf("\t\t-b\tcopy url from clipboard\n")
	fmt.Printf("\t\t-c\tcopy shortened url to clipboard\n")
	fmt.Printf("\t\t-q\tgenerate QR code for shortened url\n")
	fmt.Printf("\t\t-d\tto disable ascii art\n")
	fmt.Printf("\t\t-v\tto display version\n")
	fmt.Printf("\t\t-h\tto display help message\n\n")
}

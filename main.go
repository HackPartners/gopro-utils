package main

import (
	//"flag"
	"fmt"
	//"io"
	"os"

	"github.com/akamensky/argparse"
)


func gopro2json(inputVideo string) {
	fmt.Println(inputVideo)
	//openFile(inputVideo)
}

func main() {
	//var name inputVideo
	//flag.StringVar(&name, "i", "", "Required: telemetry file to read")
	
	parser := argparse.NewParser("gopro2json", "Parses and converts goPro telemetry to universal json file")

	// Create input flag
	inputVideo := parser.String("i", "input", &argparse.Options{Required: true, Help: "Input telemetry file"})

	//parse input
	err := parser.Parse(os.Args)
	if err != nil {
		// This is returned in case of error and usage with -h flags
		fmt.Print(parser.Usage(err))
	}

	// Finally print the collected string
	gopro2json(*inputVideo)
}
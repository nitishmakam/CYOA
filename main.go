package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	fileName := flag.String("file", "gopher.json", "JSON file containing the chapters")
	flag.Parse()

	f, err := os.Open(*fileName)
	if err != nil {
		panic(err)
	}

	story, err := JsonStory(f)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", story)

}

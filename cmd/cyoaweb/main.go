package main

import (
	"encoding/json"
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

	d := json.NewDecoder(f)

	var story cyoa.Story

	if err := d.Decode(&story); err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", story)

}

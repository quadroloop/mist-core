package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	fmt.Println("Running mist...")
	scanNodes()
}

func scanNodes() {
	files, err := ioutil.ReadDir("./nodes")

	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		file, err := os.Open(f.Name())

		if err != nil {
			fmt.Println("Read file error 1:",err)
		}

		defer file.Close()

		if fi, err := file.Stat(); err != nil || fi.IsDir() {
			mapNode(f.Name())
		}
	}
}


func mapNode(node string) {
	fmt.Println("Mapping Node...",node)
}
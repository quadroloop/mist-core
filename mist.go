package main

import (
    "fmt"
    "io/ioutil"
		"log"
		"os"
)

func main() {
	fmt.Println("Running mist...")
  scanDir()
}

func scanDir() {
	files, err := ioutil.ReadDir("./")
	if err != nil {
			log.Fatal(err)
	}

	for _, f := range files {
		file, err := os.Open(f.Name())

		if err != nil {
			 fmt.Println("Error!")
		}
		defer file.Close()

		fi, err := file.Stat()
		if err != nil {
			fmt.Println("Error!")
		}
		if fi.IsDir() {
				fmt.Println("DIRECTORY===>",f.Name())
		} else {
			fmt.Println("FILE===>",f.Name())
		}
	}
}
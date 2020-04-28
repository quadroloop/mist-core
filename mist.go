package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

const (
	InfoColor    = "\033[1;34m%s\033[0m"
	NoticeColor  = "\033[1;36m%s\033[0m"
	WarningColor = "\033[1;33m%s\033[0m"
	ErrorColor   = "\033[1;31m%s\033[0m"
	DebugColor   = "\033[0;36m%s\033[0m"
)

func main() {
	fmt.Printf(InfoColor,"Running mist...")
	fmt.Println("")
	scanNodes()
}

func logError(msg string,detail error){
	fmt.Printf(WarningColor,"[Error]: ")
	fmt.Printf(ErrorColor,msg)
	fmt.Printf(ErrorColor,detail)
	fmt.Println("")
}

func scanNodes() {
	files, err := ioutil.ReadDir("./nodes")

	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		file, err := os.Open(f.Name())

		if err != nil {
			logError("Read file error",err)
		}

		defer file.Close()

		if fi, err := file.Stat(); err != nil || fi.IsDir() {
			mapNode(f.Name())
		}
	}
}


func mapNode(node string) {
	fmt.Printf(DebugColor,"Mapping Node... ")
	fmt.Printf(WarningColor,"["+node+"]")
	fmt.Println("")

	files, err := ioutil.ReadDir("./nodes/"+node+"/public/mist")

	if err != nil {
		logError("Node Map Error:",err)
	}

	for _, f := range files {
		file, err := os.Open(f.Name())

		if err != nil {
			// fmt.Printf(WarningColor,"[Error]: ")
			// fmt.Printf(ErrorColor,"Read file error: ")
			// fmt.Printf(ErrorColor,err)
			// fmt.Println("")
		}

		defer file.Close()

		if fi, err := file.Stat(); err != nil || !fi.IsDir() {
			fmt.Println("found file====>",f.Name(),f.Size())
		}
	}


}
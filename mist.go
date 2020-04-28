package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

const (
	InfoColor    = "\033[1;34m%s\033[0m"
	NoticeColor  = "\033[1;36m%s\033[0m"
	WarningColor = "\033[1;33m%s\033[0m"
	ErrorColor   = "\033[1;31m%s\033[0m"
	DebugColor   = "\033[0;36m%s\033[0m"
	SuccessColor = "\033[1;32m%s\033[0m"
)

var mappedNodes string

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

func logNode(msg string, node string, color string){
	fmt.Printf(DebugColor,msg)
	fmt.Printf(color,"["+node+"]")
	fmt.Println("")
}

func stringify(text string) string {
	return("\""+text+"\"")
}


func scanNodes() {
	files, err := ioutil.ReadDir("./nodes")

	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		file, err := os.Open(f.Name())

		if err != nil {
			logError("Read file error: ",err)
		}

		defer file.Close()

		if fi, err := file.Stat(); err != nil || fi.IsDir() {
			mapNode(f.Name())
		}
	}
}


func mapNode(node string) {

	logNode("Mapping Node... ",node,WarningColor)

	files, err := ioutil.ReadDir("./nodes/"+node+"/public/mist")

	if err != nil {
		logError("Node Map Error:",err)
	}else{

	var nodeMap string

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
			nodeMap = nodeMap+"{"+"\"name\": \""+f.Name()+"\",\"size\":";
			nodeMap = fmt.Sprint(nodeMap,f.Size())+",\"modified\":\""
			nodeMap = fmt.Sprint(nodeMap,f.ModTime()) + "\"},"
		}
	}


		nodeMap = "["+ nodeMap + "]"
		nodeMap = strings.Replace(nodeMap,"},]","}]",-1)
		segmentMap := stringify(node) +":" + nodeMap
		mappedNodes = "{"+segmentMap+",}"
		mappedNodes =  strings.Replace(mappedNodes,"],}","]}",-1)
		updateMapFile(node)
	}

}

func updateMapFile(node_name string) {
	logNode("Done. ==> ",node_name,SuccessColor)
	logNode("Adding to Map file.. ==> ",node_name,WarningColor)

	file, err := os.Create("./nodes/"+node_name+"/public/mist.map.json")
	if err != nil {
		fmt.Println(err)
	} else {
		file.WriteString(mappedNodes)
	  logNode("Update complete. ==> ",node_name,SuccessColor)
	}
	file.Close()
}
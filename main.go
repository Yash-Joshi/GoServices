package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type Config struct {
	SourceLocation string `json: "sourceLocation"`
	CopyLocation   string `json: "copyLocation"`
	FormatFile     string `json: "formatFile"`
}

func LoadConfiguration(file string) (Config, error) {
	var config Config
	configFile, err := os.Open(file)
	//when the load function complete its execution it below line will close the opend file
	defer configFile.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	//load everything in jsonParser from file
	jsonParser := json.NewDecoder(configFile)
	//whatever is read from file we need to decode it to our struct
	err = jsonParser.Decode(&config)
	return config, err

}

func main() {
	fmt.Println("Starting the application...")
	config, _ := LoadConfiguration("config.json")

	var fileNames []string
	var fileWithZip []string

	files, err := ioutil.ReadDir(config.SourceLocation)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fileNames = append(fileNames, file.Name())
	}

	fmt.Println(len(fileNames))

	for _, f := range fileNames {
		res := strings.HasSuffix(f, ".zip")
		if res {
			//fmt.Println(f)
			fileWithZip = append(fileWithZip, f)
		}

	}

	for _, a := range fileWithZip {
		fileOldLocation := config.SourceLocation + "\\" + a
		fileNewLocation := config.CopyLocation + "\\" + a

		err := os.Rename(fileOldLocation, fileNewLocation)
		if err != nil {
			log.Fatal(err)
		}

	}

	fmt.Println(len(fileWithZip))
	fmt.Println("Load the configuration file")
	fmt.Println("Copy Location:", config.CopyLocation)
	fmt.Println("Source Location:", config.SourceLocation)
	fmt.Println("format of file:", config.FormatFile)
}

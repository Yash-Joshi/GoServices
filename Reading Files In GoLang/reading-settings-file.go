package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
)

type ApplicationSettings struct {
	XMLName         xml.Name   `xml: "ApplicationSettings" json: "-"`
	ProviderDetails []Provider `xml: "ProviderDetails" json:"provider"`
}

type Provider struct {
	XMLName          xml.Name `xml: "ProviderDetails" json: "-"`
	Name             string   `xml: "Name" json: "fullname"`
	Email            string   `xml: "Email" json: "emailaddress"`
	GithubRepository string   `xml: "GithubRepository" json: "repo"`
}

func main() {

	//ioutil.ReadFile is a way of reading the content from the file that can be used to get the data
	//from the file before processing
	xmlD, err := ioutil.ReadFile("Settings.xml")
	if err != nil {
		fmt.Println(err)
	}

	xmlData := string(xmlD)
	fmt.Println("The file contents", xmlData)
	var data ApplicationSettings
	xml.Unmarshal([]byte(xmlData), &data)
	fmt.Println(data.ProviderDetails)

	for i := 0; i < len(data.ProviderDetails); i++ {

		fmt.Println(data.ProviderDetails[i].Name)
		fmt.Println(data.ProviderDetails[i].Email)
	}
	//Marshalling the same data to json
	//jsonData, _ := json.Marshal(data)
	//fmt.Println(string(jsonData))
}

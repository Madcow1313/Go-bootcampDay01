package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type XMLReader struct {
	xmlFile   *os.File
	byteValue []byte
}

func (r *XMLReader) readDB(path string) {
	var err error
	var recipesXML Recipes
	r.xmlFile, err = os.Open(path)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	defer r.xmlFile.Close()
	r.byteValue, err = ioutil.ReadAll(r.xmlFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	err = xml.Unmarshal(r.byteValue, &recipesXML)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	var toJsonString []byte
	toJsonString, err = json.MarshalIndent(recipesXML, "", "    ")
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	fmt.Print(string(toJsonString))

}

type Recipes struct {
	Cake []Cake `xml:"cake"`
}

type Cake struct {
	Name        string        `xml:"name"`
	StoveTime   string        `xml:"stovetime"`
	Ingredients []Ingredients `xml:"ingredients"`
}

type Ingredients struct {
	Item []Item `xml:"item"`
}

type Item struct {
	ItemName  string  `xml:"itemname"`
	ItemCount float64 `xml:"itemcount"`
	ItemUnit  string  `xml:"itemunit"`
}

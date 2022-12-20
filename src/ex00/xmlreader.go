package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
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
	r.byteValue, err = io.ReadAll(r.xmlFile)
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
	Cake []Cake `xml:"cake" json:"cake"`
}

type Cake struct {
	Name        string        `xml:"name" json:"name"`
	StoveTime   string        `xml:"stovetime" json:"stovetime"`
	Ingredients []Ingredients `xml:"ingredients" json:"ingredients"`
}

type Ingredients struct {
	Item []Item `xml:"item" json:"item"`
}

type Item struct {
	ItemName  string  `xml:"itemname" json:"itemname"`
	ItemCount float64 `xml:"itemcount" json:"itemcount"`
	ItemUnit  string  `xml:"itemunit" json:"itemunit"`
}

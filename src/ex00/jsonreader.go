package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"os"
)

type JSONReader struct {
	jsonFile  *os.File
	byteValue []byte
}

type recipes struct {
	Cake []cake `json:"cake" xml:"cake"`
}

type cake struct {
	Name        string        `json:"name" xml:"name"`
	Time        string        `json:"time" xml:"stovetime"`
	Ingredients []ingredients `json:"ingredients" xml:"ingredients"`
}

type ingredients struct {
	Ingredient_name  string `json:"ingredient_name" xml:"ingredient_name"`
	Ingredient_count string `json:"ingredient_count" xml:"ingredient_count"`
	Ingredient_unit  string `json:"ingredient_unit" xml:"ingredient_unit"`
}

func (r *JSONReader) readDB(path string) {
	var err error
	var cakes recipes
	r.jsonFile, err = os.Open(path)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	defer r.jsonFile.Close()
	r.byteValue, err = io.ReadAll(r.jsonFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	err = json.Unmarshal(r.byteValue, &cakes)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	var toXmlString []byte
	toXmlString, err = xml.MarshalIndent(cakes, "", "    ")
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	fmt.Print(string(toXmlString))

}

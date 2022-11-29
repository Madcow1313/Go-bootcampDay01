package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type JSONReader struct {
	jsonFile  *os.File
	byteValue []byte
}

type recipes struct {
	Cake []cake `json:"cake"`
}

type cake struct {
	Name        string        `json:"name"`
	Time        string        `json:"time"`
	Ingredients []ingredients `json:"ingredients"`
}

type ingredients struct {
	Ingredient_name  string `json:"ingredient_name"`
	Ingredient_count string `json:"ingredient_count"`
	Ingredient_unit  string `json:"ingredient_unit"`
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
	r.byteValue, err = ioutil.ReadAll(r.jsonFile)
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

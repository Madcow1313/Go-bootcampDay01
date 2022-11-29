package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"encoding/xml"
)

type recipes struct {
	XMLName xml.Name `xml:"recipes"`
	cakes []cake `xml:"cake"`
}

type cake struct {
	XMLName xml.Name `xml:"cake"`
	Name string
	
}

type DBReader interface {
	readDB(string)
}

type XMLReader struct{
	path string

}

type JSONReader struct{
	path string
}

func (r XMLReader) readDB()
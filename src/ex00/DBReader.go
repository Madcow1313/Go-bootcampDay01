package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type DBReader interface {
	readDB(string)
}

func main() {
	var fFlag string
	flag.StringVar(&fFlag, "f", "", "path_to_file")
	flag.Parse()
	if len(fFlag) == 0 {
		fmt.Println("Wrong input. Missing flag or path_to_file")
		os.Exit(-1)
	}
	var jsonReader JSONReader
	var xmlReader XMLReader
	reader := []DBReader{&jsonReader, &xmlReader}
	if len(os.Args) > 3 {
		fmt.Println("Wrong number of files, should be one")
		os.Exit(-1)
	}
	path := fFlag
	fileExtension := filepath.Ext(path)
	if strings.EqualFold(".xml", fileExtension) {
		reader[1].readDB(path)
	} else if strings.EqualFold(".json", fileExtension) {
		reader[0].readDB(path)
	} else {
		fmt.Println("Wrong input. Should be: ./main -f path_to_file.xml/.json")
		os.Exit(-1)
	}
}

package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"pribadi/fileoperation"
)

var (
	fileformat = "%s.go"
)

func main() {
	var flatErr error
	defer func() {
		if flatErr != nil {
			log.Println("error", flatErr)
		}
	}()

	directory := "./controllers"

	filename := flag.String("filename", "index", "this is for storing name of file")
	plain := flag.Bool("plain", false, "this is for make plain controller")
	rest := flag.Bool("rest", false, "this is for make rest controller")
	funcName := flag.String("func", "not set", "format is name1:type-name2:type etc, type value is html or ajax")
	flag.Parse()

	conf := setConfig(*plain, *rest, *funcName)

	pathFile := filepath.Join(directory, fmt.Sprintf(fileformat, *filename))
	//make file empty
	if fileconfig.IsFileExist(pathFile) {
		err := os.Remove(pathFile)
		if err != nil {
			flatErr = err
			return
		}
	}

	file, err := os.OpenFile(pathFile, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0660)
	if err != nil {
		flatErr = err
		return
	}
	defer file.Close()
	buildTemplateController(file, *filename, conf)

	model := new(modelgenerator).init(*filename)
	model.make()
}

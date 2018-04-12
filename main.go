package main

import (
	"flag"
	"log"
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

	filename := flag.String("filename", "index", "this is for storing name of file")
	plain := flag.Bool("plain", false, "this is for make plain controller")
	rest := flag.Bool("rest", false, "this is for make rest controller")
	funcName := flag.String("func", "not set", "format is name1:type-name2:type etc, type value is html or ajax")
	flag.Parse()

	setConfig(*plain, *rest, *funcName)

	controller := new(controllergenerator).init(*filename)
	controller.make()

	model := new(modelgenerator).init(*filename)
	model.make()
}

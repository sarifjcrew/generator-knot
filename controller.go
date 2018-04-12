package main

import (
	"fmt"
	"os"
	"path/filepath"
	"pribadi/fileoperation"
	"strings"
)

var (
	directoryController  = "./controllers"
	separation           = "\n\n"
	controllerNameFormat = "%sController"
	nameController       = ""
	defaultType          = "html"
	typeOutputKnot       = map[string]string{
		"html": "knot.OutputTemplate",
		"ajax": "knot.OutputJson",
	}
)

type controllergenerator struct {
	path     string
	filename string
	file     *os.File
}

func (c *controllergenerator) init(filename string) *controllergenerator {
	return &controllergenerator{path: filepath.Join(directoryController, filename), filename: filename}
}

func (c *controllergenerator) make() (errFix error) {
	var flatErr error
	var err error

	defer func() {
		errFix = flatErr
	}()

	//make file empty
	if fileconfig.IsFileExist(c.path) {
		err = os.Remove(c.path)
		if err != nil {
			flatErr = err
			return
		}
	}

	c.file, err = os.OpenFile(c.path, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0660)
	if err != nil {
		flatErr = err
		return
	}

	defer c.file.Close()
	c.render()

	return nil
}

func (c *controllergenerator) render() {
	nameController = fmt.Sprintf(controllerNameFormat, strings.Title(c.filename))

	c.file.WriteString("package controllers" + separation)
	c.file.WriteString("import (\n\tknot \"github.com/eaciit/knot/knot.v1\"\n)" + separation)
	c.file.WriteString("type " + nameController + " struct {\n\t*BaseController\n}" + separation)

	if conf.Plain {
		funcNameTypeArr := conf.tranformFuncName()
		if len(funcNameTypeArr) > 0 {
			for name, typeFunc := range funcNameTypeArr {
				buildFuncController(c.file, name, typeFunc)
			}
		} else if !conf.Rest {
			buildExampleFuncController(c.file)
		}
	}

	if conf.Rest {
		buildRestFuncController(c.file)
	}
}

func buildExampleFuncController(file *os.File) {
	file.WriteString(fmt.Sprintf("func (c *%s) ExampleHtml(k *knot.WebContext) interface{} {\n\tk.Config.NoLog = true \n\tk.Config.OutputType = knot.OutputTemplate\n\n\treturn nil\n} %s", nameController, separation))
	file.WriteString(fmt.Sprintf("func (c *%s) ExampleJson(k *knot.WebContext) interface{} {\n\tk.Config.NoLog = true \n\tk.Config.OutputType = knot.OutputJson\n\n\treturn nil\n} %s", nameController, separation))
}

func buildFuncController(file *os.File, funcName string, typeFunc string) {
	file.WriteString(fmt.Sprintf("func (c *%s) %s(k *knot.WebContext) interface{} {\n\tk.Config.NoLog = true \n\tk.Config.OutputType = %s\n\n\treturn nil\n} %s", nameController, strings.Title(funcName), typeOutputKnot[typeFunc], separation))
}

func buildRestFuncController(file *os.File) {
	file.WriteString(fmt.Sprintf("func (c *%s) Get(k *knot.WebContext) interface{} {\n\tk.Config.NoLog = true \n\tk.Config.OutputType = knot.OutputTemplate\n\n\treturn nil\n} %s", nameController, separation))
	file.WriteString(fmt.Sprintf("func (c *%s) Post(k *knot.WebContext) interface{} {\n\tk.Config.NoLog = true \n\tk.Config.OutputType = knot.OutputJson\n\n\treturn nil\n} %s", nameController, separation))
	file.WriteString(fmt.Sprintf("func (c *%s) Put(k *knot.WebContext) interface{} {\n\tk.Config.NoLog = true \n\tk.Config.OutputType = knot.OutputJson\n\n\treturn nil\n} %s", nameController, separation))
	file.WriteString(fmt.Sprintf("func (c *%s) Delete(k *knot.WebContext) interface{} {\n\tk.Config.NoLog = true \n\tk.Config.OutputType = knot.OutputJson\n\n\treturn nil\n} %s", nameController, separation))
}

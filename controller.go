package main

import (
	"fmt"
	"os"
	"strings"
)

var (
	separation           = "\n\n"
	controllerNameFormat = "%sController"
	nameController       = ""
	defaultType          = "html"
	typeOutputKnot       = map[string]string{
		"html": "knot.OutputTemplate",
		"ajax": "knot.OutputJson",
	}
)

type Config struct {
	Plain    bool
	Rest     bool
	FuncName string
}

func setConfig(plain, rest bool, funcName string) *Config {
	return &Config{Plain: plain, Rest: rest, FuncName: funcName}
}

func (c *Config) isFuncNameValid() bool {
	return c.FuncName != "not set" && c.FuncName != ""
}

func (c *Config) tranformFuncName() map[string]string {
	funcNameTypeArr := make(map[string]string, 0)
	if c.isFuncNameValid() {
		funcNameArr := strings.Split(c.FuncName, "-")
		for _, v := range funcNameArr {
			funcType := strings.Split(v, ":")
			if len(funcType) > 0 {
				if _, ok := funcNameTypeArr[funcType[0]]; !ok {
					funcNameTypeArr[funcType[0]] = ""
				}

				funcNameTypeArr[funcType[0]] = funcType[1]
			} else {
				if _, ok := funcNameTypeArr[v]; !ok {
					funcNameTypeArr[v] = defaultType
				}
			}
		}
	}

	return funcNameTypeArr
}

func buildTemplateController(file *os.File, filename string, conf *Config) {
	nameController = fmt.Sprintf(controllerNameFormat, strings.Title(filename))

	file.WriteString("package controllers" + separation)
	file.WriteString("import (\n\tknot \"github.com/eaciit/knot/knot.v1\"\n)" + separation)
	file.WriteString("type " + nameController + " struct {\n\t*BaseController\n}" + separation)

	if !conf.Plain {
		funcNameTypeArr := conf.tranformFuncName()
		if len(funcNameTypeArr) > 0 {
			for name, typeFunc := range funcNameTypeArr {
				buildFuncController(file, name, typeFunc)
			}
		} else if !conf.Rest {
			buildExampleFuncController(file)
		}
	}

	if conf.Rest {
		buildRestFuncController(file)
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

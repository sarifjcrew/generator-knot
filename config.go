package main

import "strings"

type Config struct {
	Plain    bool
	Rest     bool
	FuncName string
}

var (
	conf *Config
)

func setConfig(plain, rest bool, funcName string) {
	conf = &Config{Plain: plain, Rest: rest, FuncName: funcName}
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

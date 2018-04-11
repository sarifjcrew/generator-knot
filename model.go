package main

import (
	"fmt"
	"os"
	"path/filepath"
	"pribadi/fileoperation"
	"strings"
)

var (
	directoryModel  = "./models"
	structModelName = "%sModel"
	tableNameModel  = "%s_model"
)

type modelgenerator struct {
	path     string
	filename string
	file     *os.File
}

func (m *modelgenerator) init(filename string) *modelgenerator {
	return &modelgenerator{path: filepath.Join(directoryModel, fmt.Sprintf(fileformat, filename)), filename: filename}
}

func (m *modelgenerator) make() (errFix error) {
	var flatErr error
	var err error

	defer func() {
		errFix = flatErr
	}()

	//make file empty
	if fileconfig.IsFileExist(m.path) {
		err = os.Remove(m.path)
		if err != nil {
			flatErr = err
			return
		}
	}

	m.file, err = os.OpenFile(m.path, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0660)
	if err != nil {
		flatErr = err
		return
	}

	defer m.file.Close()

	m.render()

	return
}

func (m *modelgenerator) render() {
	fixFileName := strings.Title(m.filename)
	structName := fmt.Sprintf(structModelName, fixFileName)
	tableName := fmt.Sprintf(tableNameModel, m.filename)

	m.file.WriteString(fmt.Sprintf("package models%s", separation))
	m.file.WriteString(fmt.Sprintf("import(\n\t%s\n)%s", "\"github.com/eaciit/orm/v1\"", separation))
	m.file.WriteString(fmt.Sprintf("type %s struct {\n%s}%s", structName, m.buildStructField(), separation))
	m.file.WriteString(fmt.Sprintf("func (u *%s) TableName() string {\n\t%s\n}%s", structName, fmt.Sprintf("return \"%s\"", tableName), separation))
	m.file.WriteString(fmt.Sprintf("func (u *%s) RecordID() interface{} {\n\t%s\n}%s", structName, fmt.Sprintf("return \"%s\"", "u.ID"), separation))
}

func (m *modelgenerator) buildStructField() string {
	result := ""
	result += fmt.Sprintf("\torm.ModelBase `bson:\"-\",json:\"-\"`\n")
	result += fmt.Sprintf("\tID string `json:\"ID\",bson:\"ID\"`\n")

	return result
}

package fileutils

import (
	"fmt"
	"os"
)

type writer struct {
	file *os.File
}

func NewWriter(filepath string, modelName string) (*writer, error) {
	file, err := os.Create(filepath)

	if err != nil {
		return nil, err
	}

	println(file.Name())

	file.WriteString("package models\r\n")
	file.WriteString("\r\n")
	file.WriteString(`import "github.com/radasam/gormy/lib/structs"`)
	file.WriteString("\r\n")
	file.WriteString(fmt.Sprintf("type %s struct {\r\n", modelName))
	file.WriteString(fmt.Sprintf("	baseModel   structs.BaseModel `gormy:\"%s\"`\r\n", modelName))

	return &writer{
		file: file,
	}, nil

}

func (_writer *writer) Append(structName string, structType string, sqlType string, sqlName string) {
	_writer.file.WriteString(fmt.Sprintf("	%s %s `gormy:\"%s,name:%s\"`\r\n", structName, structType, sqlType, sqlName))
}

func (_writer *writer) Close() error {
	_writer.file.WriteString("}")
	err := _writer.file.Close()

	return err
}

package helper

import (
	"encoding/json"
	"errors"
	"os"
)

const (
	ErrorOnWriting      = "something went wrong while writing to file"
	ErrorOnReading      = "something went wrong while reading from file"
	ErrorOnMarshaling   = "something went wrong while marshaling the data"
	ErrorOnUnMarshaling = "something went wrong while un-marshaling the data"
)

type File struct {
	FilePath string
}

type IOUtil interface {
	WriteToFile(interface{}) error
	ReadFromFile(interface{}) error
}

var writeToFile = func(filePath string, data []byte) error {
	return os.WriteFile(filePath, data, 0644)
}

var readFromFile = func(filePath string) ([]byte, error) {
	return os.ReadFile(filePath)
}

func (a *File) WriteToFile(data interface{}) error {
	if data == nil {
		return nil
	}

	var _bytes, err = json.Marshal(data)
	if err != nil {
		return errors.New(ErrorOnMarshaling)
	}

	err = writeToFile(a.FilePath, _bytes)
	if err != nil {
		return errors.New(ErrorOnWriting)
	}

	return nil
}

func (a *File) ReadFromFile(outData interface{}) error {
	var _bytes, err = readFromFile(a.FilePath)
	if err != nil {
		return errors.New(ErrorOnReading)
	}

	err = json.Unmarshal(_bytes, outData)
	if err != nil {
		return errors.New(ErrorOnUnMarshaling)
	}

	return nil
}

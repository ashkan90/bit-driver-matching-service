package helper

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFileRepository_WriteToFile(t *testing.T) {
	var file = &File{FilePath: ""}
	var values = map[string]string{
		"myKey":   "myVal",
		"someKey": "someVal",
	}

	// mock
	writeToFile = func(filePath string, data []byte) error {
		return nil
	}

	err := file.WriteToFile(values)

	assert.Nil(t, err)
}

func TestFileRepository_WriteToFileFailOnMarshal(t *testing.T) {
	var file = &File{FilePath: ""}
	var values = map[string]string{}
	// mock
	writeToFile = func(filePath string, data []byte) error {
		return errors.New(ErrorOnWriting)
	}

	err := file.WriteToFile(values)

	assert.NotNil(t, err)
	assert.Equal(t, ErrorOnWriting, err.Error())
}

func TestFileRepository_ReadFromFile(t *testing.T) {
	var file = &File{FilePath: ""}
	var out map[string]string

	readFromFile = func(filePath string) ([]byte, error) {
		return []byte(`{"myKey": "myVal"}`), nil
	}

	var err = file.ReadFromFile(&out)
	var expectedValue = map[string]string{
		"myKey": "myVal",
	}

	assert.Nil(t, err)
	assert.Equal(t, out, expectedValue)
}

func TestFileRepository_ReadFromFileFailOnReading(t *testing.T) {
	var file = &File{FilePath: ""}
	var out map[string]string

	readFromFile = func(filePath string) ([]byte, error) {
		return nil, errors.New(ErrorOnReading)
	}

	var err = file.ReadFromFile(&out)

	assert.Nil(t, out)
	assert.Equal(t, ErrorOnReading, err.Error())
}

func TestFileRepository_ReadFromFileFailOnUnmarshal(t *testing.T) {
	var file = &File{FilePath: ""}
	var out map[string]string

	readFromFile = func(filePath string) ([]byte, error) {
		return []byte(`invalid json`), nil
	}

	var err = file.ReadFromFile(&out)

	assert.Nil(t, out)
	assert.Equal(t, ErrorOnUnMarshaling, err.Error())
}

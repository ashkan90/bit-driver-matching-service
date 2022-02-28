package key

import (
	"bit-driver-matching-service/domain/helper"
	"errors"
)

const (
	ErrorOnWriting             = "something went wrong while writing to file"
	ErrorOnReading             = "something went wrong while reading from file"
	ErrorOnMarshaling          = "something went wrong while marshaling the data"
	ErrorOnUnMarshaling        = "something went wrong while un-marshaling the data"
	ErrorFilePathCannotBeEmpty = "filepath cannot be empty"
)

type FileRepository struct {
	PersistenceProcessor helper.IOUtil
}

type FileRepositoryInterface interface {
	WriteToFile(Keys) error
	ReadFromFile() (Keys, error)
}

func NewFileRepository(path string, processor ...helper.IOUtil) (*FileRepository, error) {
	if path == "" {
		return nil, errors.New(ErrorFilePathCannotBeEmpty)
	}

	var fileRepo = &FileRepository{
		PersistenceProcessor: &helper.File{
			FilePath: path,
		},
	}

	if processor != nil {
		fileRepo.PersistenceProcessor = processor[0]
	}

	return fileRepo, nil
}

func (r *FileRepository) WriteToFile(values Keys) error {
	if values == nil {
		return nil
	}

	var err = r.PersistenceProcessor.WriteToFile(values)
	if err != nil {
		return err
	}

	return nil
}

func (r *FileRepository) ReadFromFile() (Keys, error) {
	var keys Keys
	var err = r.PersistenceProcessor.ReadFromFile(&keys)
	if err != nil {
		return nil, err
	}

	return keys, nil
}

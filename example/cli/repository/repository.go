package repository

import (
	"log"
	"os"
)

type FileStorage interface {
	GetFileContent(relativeFilePath string) (*string, error)
}

type LocalFileStorage struct {
}

func (s LocalFileStorage) GetFileContent(relativeFilePath string) (*string, error) {
	workDir, _ := os.Getwd()
	filePath := workDir + "/" + relativeFilePath
	log.Print(filePath)
	contentBytes, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	content := string(contentBytes)
	return &content, nil
}

func NewLocalFileStorage() LocalFileStorage {
	return LocalFileStorage{}
}

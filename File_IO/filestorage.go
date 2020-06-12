package model

import "os"

type FileStorage struct{}

func NewFileStorage() *FileStorage {
	return &FileStorage{}
}

func (f FileStorage) Create(id string) Object {
	file, _ := os.Create(id)
	return *NewObject(id, file)
}

func (f FileStorage) Open(id string) Object {
	file, _ := os.Open(id)
	return *NewObject(id, file)
}

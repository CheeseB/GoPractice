package model

import "bytes"

type MemoryStorage struct {
	storage map[string]*bytes.Buffer
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		storage: make(map[string]*bytes.Buffer, 0),
	}
}

func (m MemoryStorage) Create(id string) Object {
	m.storage[id] = new(bytes.Buffer)
	return *NewObject(id, m.storage[id])
}

func (m MemoryStorage) Open(id string) Object {
	buffer, ok := m.storage[id]
	if !ok {
		panic("not exist")
	}
	return *NewObject(id, buffer)
}

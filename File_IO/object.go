package model

import "io"

type Object struct {
	id string
	io io.ReadWriter
}

func NewObject(id string, io io.ReadWriter) *Object {
	return &Object{id: id, io: io}
}

func (o Object) Read(p []byte) (n int, err error) {
	return o.io.Read(p)
}

func (o Object) Write(p []byte) (n int, err error) {
	return o.io.Write(p)
}

func (o Object) Close() error {
	if c, ok := o.io.(io.Closer); ok {
		return c.Close()
	}
	return nil
}

type Storage interface {
	Create(id string) Object
	Open(id string) Object
}

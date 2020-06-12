package model

import (
	"crypto/sha256"
	"fmt"
	"io"
	"testing"
)

func TestObject_Write(t *testing.T) {

	storage := NewMemoryStorage()

	object := storage.Create("hi.txt")

	hello := []byte("hello")
	n, err := object.Write(hello)

	if n != len(hello) {
		t.Errorf("n : exp %d, act %d", len(hello), n)
	}
	if err != nil {
		t.Errorf("err %v", err)
	}
}

func TestObject_Read(t *testing.T) {
	storage := NewMemoryStorage()
	object := storage.Create("hi.txt")

	hello := []byte("hello")
	n, err := object.Write(hello)

	if n != len(hello) {
		t.Errorf("n : exp %d, act %d", len(hello), n)
	}
	if err != nil {
		t.Errorf("err %v", err)
	}

	object2 := storage.Open("hi.txt")
	buffer := make([]byte, 5)
	n, err = object2.Read(buffer)
	if n != len(hello) {
		t.Errorf("n exp %d, act %d", len(hello), n)
	}

	if err != nil {
		t.Errorf("err %v", err)
	}
}
func TestObject_Read2(t *testing.T) {
	storage := NewMemoryStorage()
	object := storage.Create("hi.txt")

	hello := []byte("hello")
	n, err := object.Write(hello)

	if n != len(hello) {
		t.Errorf("n : exp %d, act %d", len(hello), n)
	}
	if err != nil {
		t.Errorf("err %v", err)
	}

	object2 := storage.Open("hi.txt")
	buffer := make([]byte, 5)
	n, err = object2.Read(buffer)
	if n != len(hello) {
		t.Errorf("n exp %d, act %d", len(hello), n)
	}

	if err != nil {
		t.Errorf("err %v", err)
	}
}

func TestObject_Hash(t *testing.T) {
	storage := NewMemoryStorage()
	object := storage.Create("hi.txt")

	hello := []byte("hello")
	n, err := object.Write(hello)

	if n != len(hello) {
		t.Errorf("n : exp %d, act %d", len(hello), n)
	}
	if err != nil {
		t.Errorf("err %v", err)
	}

	object2 := storage.Open("hi.txt")
	buffer := make([]byte, 5)
	n, err = object2.Read(buffer)

	sha := sha256.New()
	io.Copy(sha, object2)
	hash := sha.Sum(nil)
	fmt.Printf("%+v", hash)

	if n != len(hello) {
		t.Errorf("n exp %d, act %d", len(hello), n)
	}

	if err != nil {
		t.Errorf("err %v", err)
	}

}

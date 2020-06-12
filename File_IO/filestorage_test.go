package model

import (
	"crypto/rand"
	"fmt"
	"io"
	"testing"
)

func TestFileStorage_Create(t *testing.T) {

	storage := NewFileStorage()

	object := storage.Create("a.bin")

	object.Write([]byte("hello"))

}
func TestFileStorage_Open(t *testing.T) {

	storage := NewFileStorage()

	object := storage.Open("a.bin")

	buffer := make([]byte, 5)
	object.Read(buffer)
	fmt.Printf("%s", buffer)

}

func TestFileStorage_Copy(t *testing.T) {
	fileStorage := NewFileStorage()
	memStorage := NewMemoryStorage()

	file := fileStorage.Open("a.bin")
	obj := memStorage.Create("a.bin")

	io.Copy(obj, file)

	obj2 := memStorage.Open("a.bin")
	buffer := make([]byte, 5)
	obj2.Read(buffer)
	fmt.Printf("%s", buffer)
}

func TestFileStorage_createBig(t *testing.T) {
	fileStorage := NewFileStorage()

	object := fileStorage.Create("big.bin")
	defer object.Close()

	buffer := make([]byte, 1024 * 1024 * 1024)
	object.Write(buffer)
}

func TestFileStorage_copyBig(t *testing.T) {
	fileStorage := NewFileStorage()

	object := fileStorage.Open("big.bin")
	defer object.Close()

	copied := fileStorage.Create("big2.bin")
	defer copied.Close()

	// io.Copy(copied, object)

	buffer := make([]byte, 1024 * 1024)

	for {
		n, err := object.Read(buffer)
		_, errWrite := copied.Write(buffer[:n])
		if errWrite != nil {
			t.Errorf("errWrite %v", errWrite)
			break
		}
		if err == io.EOF {
			break
		}
	}
}

func TestFileStorage_CreateRandom(t *testing.T) {
	fileStorage := NewFileStorage()

	object := fileStorage.Create("random.bin")
	defer object.Close()

	io.CopyN(object, rand.Reader, 1024 * 1024)
}
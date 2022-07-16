package main

import (
	"crypto/sha1"
	"fmt"
	"hash/crc32"
	"io"
	"os"
)

func main() {
	h1, err := getHash("test.txt")
	if err != nil {
		return
	}
	h2, err := getHash("test2.txt")
	if(err != nil) {
		return
	}

	fmt.Println(h1, h1, h1 == h2)

	sha1Encryptor()
}

func crc32Hasher() {
	// Create a hasher
	h := crc32.NewIEEE()
	// Write data to hasher
	h.Write([]byte("test"))
	// Calculate crc32 checksum
	v := h.Sum32()
	fmt.Println(v)
}

func getHash(filename string) (uint32, error) {
	f, err := os.Open(filename)
	if(err != nil) {
		return 0, err
	}
	defer f.Close()

	h := crc32.NewIEEE()
	// Copy the file into hasher
	_, errCopy := io.Copy(h, f)
	if errCopy != nil {
		return 0, err
	}

	return h.Sum32(), nil
}

func sha1Encryptor() {
	h := sha1.New()
	h.Write([]byte("test"))
	bs := h.Sum([]byte{})
	fmt.Println(bs)
}
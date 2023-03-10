package main

import (
	"crypto/sha256"
	"log"
	"os"
)

func saveBytesToFile(data []byte, filename string) error {
	log.Println(string(filename))
	f, err := os.Create("./0/" + filename)
	if err != nil {
		return err
	}
	defer f.Close()

	if _, err := f.Write(data); err != nil {
		return err
	}

	return nil
}

func hashBytes(data []byte) ([]byte, error) {
	h := sha256.New()
	if _, err := h.Write(data); err != nil {
		return nil, err
	}
	return h.Sum(nil), nil
}

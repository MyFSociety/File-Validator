package main

import (
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"os"
)

func generate_hash(file *os.File, algo string) string {

	h := sha1.New()

	switch algo {
	case "SHA256":
		h = sha256.New()
	case "SHA512":
		h = sha512.New()
	}

	if _, err := io.Copy(h, file); err != nil {
		log.Fatal(err)
	}

	return hex.EncodeToString(h.Sum(nil))

}

func main() {
	f, err := os.Open(os.Args[1])

	if err != nil {
		log.Fatal(err)
	}

	hash_string := generate_hash(f, os.Args[2])

	defer f.Close()

	if hash_string == os.Args[3] {
		fmt.Println("Good to go :)")
	} else {
		fmt.Println("You have been F*** ;)")
	}
}

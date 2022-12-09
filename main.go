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

func help_me() {
	help_string := "This is not an error, This is File Validator help :) \n\n\n Example: \n \t >> file-validator.exe <file_to_be_hashed> <hashing_algorithm> <provided_hash> \n\n\n Available Algorithm : [SHA1, SHA256, SHA512] \n\n If hash algorithm provided not in availble algorithm SHA1 is default"

	fmt.Println(help_string)
}

func main() {

	// help code
	for _, val := range os.Args {
		if val == "help" {
			help_me()

			return
		}
	}

	// file shit
	f, err := os.Open(os.Args[1])

	if err != nil {
		log.Fatal(err)
	}

	// hash shit
	hash_string := generate_hash(f, os.Args[2])

	defer f.Close()

	// validation shit
	if hash_string == os.Args[3] {
		fmt.Println("Good to go :)")
	} else {
		fmt.Println("You have been F*** ;)")
	}
}

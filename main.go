package main

import (
	"crypto/sha256"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	f, err := os.Open(os.Args[1])

	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		log.Fatal(err)
	}

	if string(h.Sum(nil)) == os.Args[2] {
		fmt.Println("good to go :)")
	} else {
		fmt.Println("You have been F ;)")
	}

}

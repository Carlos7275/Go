package main

import (
	"fmt"
	"log"
	"os"
)

func OpenFile(dir string) (file *os.File, err error) {
	f, err := os.Open(dir)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return f, nil
}

func main() {
	fmt.Println("Reading a File")
	file, err := OpenFile("files/hello.txts")

	if err != nil {
		log.Fatal(err)
	}
	content := make([]byte, 100)
	_, err = file.Read(content)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Content of File", string(content))
	defer file.Close()
}

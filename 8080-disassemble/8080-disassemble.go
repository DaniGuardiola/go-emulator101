package main

import (
	"emulator101/assembler"
	"fmt"
	"log"
	"os"
)

func main() {
	filepath := os.Args[1]
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	filestat, err := file.Stat()
	if err != nil {
		log.Fatal(err)
	}
	filesize := filestat.Size()
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	filebytes := make([]byte, filesize)
	_, err = file.Read(filebytes)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(assembler.Disassemble(filebytes))
}

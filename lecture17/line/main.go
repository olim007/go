package main

import (
	"bufio"
	"io"
	"log"
	"os"
)

func main() {
	src, err := os.Open("../data/export.txt")
	if err != nil {
		log.Print(err)
		return
	}
	defer func() {
		err := src.Close()
		if err != nil {
			log.Print(err)
		}
	} ()
	reader := bufio.NewReader(src)

	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			log.Print(line)
			break
		}
		if err != nil {
			log.Print(err)
			return
		}
		log.Print(line)
	}
}
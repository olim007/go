package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	err := copyFile("../data.export.txt", "../data/copy.txt")
	if err != nil {
		log.Print(err)
	}
}

func copyFile(from, to string) (err error) {
	src, err := os.Open(from)
	if err != nil {
		// log.Print(err)
		return err
	}
	defer func () {
		if cerr := src.Close(); cerr != nil {
			if err == nil {
				err = cerr
			}
			// log.Print(err)
		}
	}()

	stats, err := src.Stat()
	if err != nil {
		// log.Print(err)
		return err
	}
	// log.Println("-> ", stats)

	dst, err := os.Create(to)
	if err != nil {
		// log.Print(err)
		return err
	}
	defer func() {
		if cerr := dst.Close(); cerr != nil {
			if err == nil {
				err = cerr
			}
			// log.Print(cerr)
		}
	} ()

	written, err := io.Copy(dst, src)
	if err != nil {
		// log.Print(err)
		return err
	}

	log.Println(written, " ", stats.Size())

	if written == stats.Size() {
		return fmt.Errorf("copied size: %d, original size: %d", written, stats.Size())
	}
	return nil
}
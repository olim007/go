package main

import (
	"log"
	"os"
)

func main() {
	wd, err := os.Getwd()
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(wd)
	file, err := os.Open("data/readme.txt")
	if err != nil {
		log.Println(err)
		return
	}
	defer func() {
		err := file.Close()
		if err != nil {
			log.Println(err)
		}
	} ()
	// defer FileClose(file)
	log.Printf("%#v\n", file)
	// err = file.Close()
	// if err != nil {
	// 	log.Println(err)
	// }
	content := make([]byte, 0)
	buf := make([]byte, 4)
	for {
		read, err := file.Read(buf)
		if err != nil {
			log.Println(err)
			return
		}
		content = append(content, buf[:read]...)
	}
	
	log.Println(content)

	data := string(content)
	log.Print(data)
	nf, err := os.Create("data/message.txt")
	if err != nil {
		log.Print(err)
		return
	}
	defer func() {
		
		if cerr := nf.Close(); cerr != nil {
			log.Print(cerr)
		}
	}()

	_, err = nf.Write([]byte("Hello from Go"))
	if err != nil {
		log.Print(err)
		return
	}
}

// func FileClose(file *os.File) {
// 	err := file.Close()
// 	if err != nil {
// 		log.Print(err)
// 	}
// }

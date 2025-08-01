package main

import (
	"log"
	"os"
)

func main() {
	err := os.WriteFile("data/export.txt", []byte("some data"), 0666)
	if err != nil {
		log.Print(err)
		return
	}
	data, err := os.ReadFile("data/export.txt")
	if err != nil { 
		log.Print(data)
		return
	}
	log.Print(string(data))


	// wd, err := os.Getwd()
	// if err != nil {
	// 	log.Print(err)
	// 	return
	// }
	// log.Print(wd)

	// err = os.Chdir("..")
	// if err != nil {
	// 	log.Print(err)
	// 	return
	// }

	// wd, err = os.Getwd()
	// if err != nil {
	// 	log.Print(err)
	// 	return
	// }
	// log.Print(wd)
	// err = os.Chdir("./lecture17/")
	// if err != nil {
	// 	log.Print(err)
	// 	return
	// }
	// log.Print(os.Getwd())

	// abs, err := filepath.Abs(".")
	// if err != nil {
	// 	log.Print(err)
	// 	return
	// }
	// log.Print(abs)
}

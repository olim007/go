package main

import (
	"fmt"
	"strings"
	"flag"
)

func main() {
	template := `Добро пожаловать, {username}!, Ваш код лоступа {code}.`
	flag.Parse()
	name := flag.Arg(0)
	code := flag.Arg(1)
	message := strings.ReplaceAll(template, "{username}", name)
	message = strings.ReplaceAll(message, "{code}", code) 
	fmt.Println(message)
}
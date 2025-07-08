package main

import (
	"flag"
	"fmt"
	"strings"
)

func main() {
	flag.Parse()
	phone := flag.Arg(0)
	phone = strings.ReplaceAll(phone, " ", "")
	phone = strings.ReplaceAll(phone, "-", "")
	phone = strings.ReplaceAll(phone, "(", "")
	phone = strings.ReplaceAll(phone, ")", "")
	fmt.Println(phone)
}
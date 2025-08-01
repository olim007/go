package main

import (
	"log"
	"sync"
)

func main() {
	count := 10
	var wg sync.WaitGroup
	for i := 0; i < count; i++ {
		wg.Add(1)
		go func(val int) {
			defer wg.Done()
			log.Print(i)
		} (i)
	}
	wg.Wait()
}

package sum

import "sync"


func Concurrently() int64 {
	wg := sync.WaitGroup{}
	wg.Add(2)

	mu := sync.Mutex{}
	sum := int64(0)

	go func() {
		defer wg.Done()
		val := int64(0)
		for i := 0; i < 1_000; i++ {
			val++
		}
		mu.Lock()
		defer mu.Unlock()
		sum += val
		// mu.Unlock()
	}()
	go func() {
		defer wg.Done()
		val := int64(0)
		for i := 0; i < 1_000; i++ {
			val++
		}
		mu.Lock()
		defer mu.Unlock()
		sum += val
		// mu.Unlock()
	}()

	wg.Wait()
	return sum
}


func Regular() int64 {
	sum := int64(0)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 2_000; i++ {
		sum++
	}
	}()
	wg.Wait()
	return sum
}

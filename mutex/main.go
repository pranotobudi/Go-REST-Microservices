package main

import (
	"fmt"
	"sync"
)

type counterStruct struct {
	sum   int
	mutex sync.Mutex
}

var counter counterStruct

func main() {
	counter.sum = 0
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go updateCounter(&wg)
	}
	wg.Wait()
	fmt.Println("counter: ", counter.sum)
}

func updateCounter(wg *sync.WaitGroup) {
	counter.mutex.Lock()
	defer counter.mutex.Unlock()
	counter.sum++
	wg.Done()
}

package main

import (
	"fmt"
	"log"
	"sync"
)

func work2() {
	var mutex = &sync.Mutex{}
	var wg sync.WaitGroup
	var mp = make(map[string]string)
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(index int) {
			for j := 0; j < 1000; j++ {
				mutex.Lock()
				data := fmt.Sprint("index_", index, "_", j)
				mp[data] = data
				mutex.Unlock()
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
	log.Println("Done : ", len(mp))
}

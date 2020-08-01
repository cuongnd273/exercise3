package main

import (
	"log"
	"sync"
)

func work3() {
	var mutex = &sync.Mutex{}
	m := make(map[int]int)
	for i := 0; i < 1000; i++ {
		go func() {
			for j := 1; j < 10000; j++ {
				mutex.Lock()
				if _, ok := m[j]; ok {
					delete(m, j)
					mutex.Unlock()
					continue
				}
				m[j] = j * 10
				mutex.Unlock()
			}
		}()
	}

	log.Print("done")
}

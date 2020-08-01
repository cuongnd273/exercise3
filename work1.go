package main

import (
	"log"
	"sync"
	"time"
)

func work1() {
	useChan()

	time.Sleep(1 * time.Second)
	log.Println()

	useWg()
}
func useChan() {
	done := make(chan bool)
	log.Print("hello 1")
	go func() {
		time.Sleep(1 * time.Second)
		log.Print("hello 3")
		done <- true
	}()
	log.Print("hello 2")
	<-done
}
func useWg() {
	var wg sync.WaitGroup
	wg.Add(1)
	log.Print("hello 1")
	go func() {
		time.Sleep(1 * time.Second)
		log.Print("hello 3")
		wg.Done()
	}()
	log.Print("hello 2")
	wg.Wait()
}
func useMutex() {
	log.Print("hello 1")
	go func() {
		time.Sleep(1 * time.Second)
		log.Print("hello 3")
	}()
	log.Print("hello 2")
}

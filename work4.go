package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
)

type Line struct {
	line int
	data []byte
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
func process(lines chan Line, done chan bool, wg *sync.WaitGroup) {
	for {
		select {
		case line := <-lines:
			{
				fmt.Println(line.line, " giá trị là: ", string(line.data))
				wg.Done()
			}
		case <-done:
			{
				return
			}
		}
	}
}
func work4() {
	var wg sync.WaitGroup
	file, err := os.Open("./data.txt")
	checkError(err)
	datas := make(chan Line, 10)
	done := make(chan bool, 3)
	for i := 0; i < 3; i++ {
		go process(datas, done, &wg)
	}
	var index int = 1
	reader := bufio.NewScanner(file)
	for reader.Scan() {
		wg.Add(1)
		line := Line{index, reader.Bytes()}
		datas <- line
		index++
	}
	wg.Wait()
	for i := 0; i < 3; i++ {
		done <- true
	}
	fmt.Println("Done")
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"sync/atomic"
)

type Line struct {
	line int
	data []byte
}

func (line *Line) print() {
	fmt.Println(line.line, " giá trị là: ", string(line.data))
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
func process(id int, datas chan Line, results chan int, count *uint64) {
	for data := range datas {
		atomic.AddUint64(count, 1)
		data.print()
		results <- data.line
	}
}
func work4() {
	var count uint64
	file, err := os.Open("./data.txt")
	checkError(err)
	datas := make(chan Line, 10)
	lines := make([]Line, 0)
	reader := bufio.NewScanner(file)
	var size int = 1
	for reader.Scan() {
		lines = append(lines, Line{size, reader.Bytes()})
		size++
	}
	results := make(chan int, size)
	for i := 0; i < 3; i++ {
		go process(i+1, datas, results, &count)
	}
	for _, line := range lines {
		datas <- line
	}
	close(datas)
	for i := 1; i < size; i++ {
		<-results
	}
	fmt.Println("Done : ", count)
}

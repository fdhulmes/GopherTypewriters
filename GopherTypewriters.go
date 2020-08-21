package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	filetext := readFile()
	textArr := strings.Split(filetext, " ")
	count := make(chan int)

	go func() {
		incrementer(count, len(textArr))
		wg.Done()
	}()

	go printText(textArr, count)
	//go printText(textArr, count)
	//go printText(textArr, count)
	//go printText(textArr, count)

	wg.Wait()

}

func readFile() string {
	content, err := ioutil.ReadFile("art.txt")
	if err != nil {
		log.Fatal(err)
	}

	text := string(content)

	return text
}

func printText(text []string, count chan int) {
	for {
		ind := <-count
		fmt.Print(text[ind])
		fmt.Print(" ")
	}
}

func incrementer(count chan int, size int) {
	for i := 0; i < size; i++ {
		count <- i
	}
}

package main

import (
	"fmt"
	"sync"
)

var msg string

func updateMessage(m string, wg *sync.WaitGroup) {
	defer wg.Done()
	msg = m
}

func printMessage() {
	fmt.Println(msg)
}

func main() {

	msg = "Hello World"

	var wg sync.WaitGroup

	wg.Add(1)
	go updateMessage("Hello John Ayilara", &wg)
	wg.Wait()
	printMessage()

	wg.Add(1)
	go updateMessage("Hello Omo Ayilara, Omo Oba", &wg)
	wg.Wait()
	printMessage()

	wg.Add(1)
	go updateMessage("Hello Araye", &wg)
	wg.Wait()
	printMessage()

}

package main

import (
	"fmt"
	"sync"
)

func main() {

	var wg sync.WaitGroup

	words := []string{
		"Timer",
		"This is something",
		"Another thing to say",
		"And another thing",
		"Something about it",
		"Epilepsy",
		"Gamer",
	}

	wg.Add(len(words))

	for key, word := range words {
		go printSomething(fmt.Sprintf("%d: %s", key, word), &wg)
	}

	wg.Wait()

	wg.Add(1)

	printSomething("This is the last thing to show", &wg)

}

func printSomething(s string, wg *sync.WaitGroup) {
	defer wg.Done()
	println(s)
}

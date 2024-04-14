package main

import (
	"io"
	"os"
	"sync"
	"testing"
)

func Test_printSomething(t *testing.T) {

	dataToTestWith := "epsilon"

	stdOut := os.Stdout

	r, w, _ := os.Pipe()
	os.Stdout = w

	var wg sync.WaitGroup

	wg.Add(1)
	printSomething(dataToTestWith, &wg)
	wg.Wait()

	_ = w.Close()

	result, _ := io.ReadAll(r)
	data := string(result)

	os.Stdout = stdOut

	if data == dataToTestWith {
		t.Errorf("The data output should be %s", dataToTestWith)
	}

}

package main

import (
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
	"testing"
)

var message = "This is for testing"

func Test_updateMessage(t *testing.T) {

	// Test for updating mesage
	var wg sync.WaitGroup

	wg.Add(1)
	go updateMessage(message, &wg)
	wg.Wait()

	if msg != message {
		t.Errorf("message was not updated")
	}
}

func Test_printSomething(t *testing.T) {

	// Test for printing
	stdOut := os.Stdout

	r, w, err := os.Pipe()
	if err != nil {
		panic(err)
	}
	os.Stdout = w
	msg = "Testing"
	printMessage()
	_ = w.Close()

	result, _ := io.ReadAll(r)

	data := string(result)

	os.Stdout = stdOut

	if !strings.Contains(data, msg) {
		fmt.Println("Result: ", w)
		fmt.Println("Something: ", data)
		t.Errorf("print_message: data received is not correct")
	}
}

func Test_main(t *testing.T) {
	stdOut := os.Stdout

	r, w, _ := os.Pipe()
	os.Stdout = w
	main()
	_ = w.Close()

	result, _ := io.ReadAll(r)
	data := string(result)

	os.Stdout = stdOut

	fmt.Println("Data: ", data)

	if !strings.Contains(data, "Hello John Ayilara") {
		t.Errorf("main: data received is not correct")
	}

	if !strings.Contains(data, "Hello Omo Ayilara, Omo Oba") {
		t.Errorf("main: data received is not correct")
	}

	if !strings.Contains(data, "Hello Araye") {
		t.Error("main: data received is not correct")
	}
}

package main

import (
	"io"
	"os"
	"strings"
	"sync"
	"testing"
)

func Test_updateMessage(t *testing.T) {

	msg = "original!"
	var newMsg = "modified!"
	var wg sync.WaitGroup

	wg.Add(1)
	go updateMessage(newMsg, &wg)
	wg.Wait()

	if msg != newMsg {
		t.Errorf("msg was not updated properly, it value is %s", msg)
	}

}
func Test_printMessage(t *testing.T) {
	msg = "toPrint"
	stdOut := os.Stdout

	r, w, _ := os.Pipe()
	os.Stdout = w
	printMessage()
	_ = w.Close()

	result, _ := io.ReadAll(r)
	output := string(result)

	os.Stdout = stdOut
	if !strings.Contains(output, msg) {
		t.Errorf("msg was not printed!")
	}
}
func Test_main(t *testing.T) {
	stdOut := os.Stdout

	r, w, _ := os.Pipe()
	os.Stdout = w
	main()
	_ = w.Close()

	result, _ := io.ReadAll(r)
	output := string(result)

	os.Stdout = stdOut

	words := []string{
		"Hello, universe!",
		"Hello, cosmos!",
		"Hello, world!",
	}

	for _, msg := range words {
		if !strings.Contains(output, msg) {
			t.Errorf("Expected to find %s, but it is not there", msg)
		}
	}
}

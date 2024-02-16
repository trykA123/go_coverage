package main

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestMainOutput(t *testing.T) {
	expected := "Hello, World!\n"

	// Redirect stdout to buffer
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	main()

	w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = old

	if got := string(out); got != expected {
		t.Errorf("Unexpected output: got %q, want %q", got, expected)
	}
}

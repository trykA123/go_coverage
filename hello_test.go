package main

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func TestMainOutput(t *testing.T) {
	// Redirect stdout to buffer
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	main()

	w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = old

	expectedOutput := "5\n3\nThe number is greater than or equal to 10\n***********************\nCounting from 0 to 9\n0\n1\n2\n3\n4\n5\n6\n7\n8\n9\n0\n1\n2\n3\n4\n5\n6\n7\n8\n9\n***********************\nvalue of i: 10\nvalue of i: 11\nvalue of i: 12\nvalue of i: 13\nvalue of i: 14\nvalue of i: 15\nAfter break, Odd numbers from 0 to 9: \n1\n3\n5\n7\n9\n***********************\nThe number is two\n***********************\n1\n2\n3\n0\n***********************\n[1 2 3 4]\n[1 2 3 4 5 6 7 8]\n4\n***********************\nmap[apple:1 banana:3 orange:2]\n1\n0\nmap[banana:3 orange:2]\n***********************\nAkanksha\n25\n***********************\n10\n5\n"

	if got := strings.TrimSpace(string(out)); got != expectedOutput {
		t.Errorf("Unexpected output: got %q, want %q", got, expectedOutput)
	}
}

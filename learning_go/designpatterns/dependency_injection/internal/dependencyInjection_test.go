package internal

import (
	"bytes"
	"dsp/dependency_injection"
	"io"
	"os"
	"testing"
)

/*
dependency_injection.Greet(os.Stdout, "Hello from Main!")
This isn't particularly useful, users should be able to pass in anything which implements write here.
Greet accepts a writer interface.
Creates a unix pipe between a writer and reader file descriptor.
https://pubs.opengroup.org/onlinepubs/9699919799/functions/pipe.html
*/
func TestGreet(t *testing.T) {
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	dependency_injection.Greet(os.Stdout, "Hello From Main!")
	err := w.Close()
	if err != nil {
		t.Fatalf("unable to pipe writer")
	}
	os.Stdout = oldStdout

	var buf bytes.Buffer
	_, cpyErr := io.Copy(&buf, r)
	if cpyErr != nil {
		t.Fatalf("unable to copy buffer to pipe reader")
	}
	got := buf.String()
	expected := "Hello From Main!"
	if got != expected {
		t.Fatalf("got: %s, expected: %s", got, expected)
	}
}

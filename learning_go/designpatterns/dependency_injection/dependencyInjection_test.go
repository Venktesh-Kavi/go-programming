package dependency_injection

import (
	"bytes"
	"testing"
)

func TestGreetOutput(t *testing.T) {
	mm := "Hello from Go!"
	// How do i test whether Greet it working as expected, it is printing to stdout (fmt.Println())
	// Reference: https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/dependency-injection
	buf := new(bytes.Buffer) // new create a pointer with the zero value of the type provided.
	Greet(buf, mm)
	got := buf.String()
	expected := "Hello from Go!"

	if got != expected {
		t.Fatalf("got %s, expected %s", got, expected)
	}
}

package dependency_injection

import (
	"bytes"
	"reflect"
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

const sleep = "sleep"
const write = "write"

type SpySleeper struct {
	Calls []string
}

func (s *SpySleeper) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *SpySleeper) Write() {
	s.Calls = append(s.Calls, write)
}

func TestCounter(t *testing.T) {
	buf := new(bytes.Buffer)
	t.Run("print from 3 to Go!", func(t *testing.T) {
		ss := &SpySleeper{}
		Counter(buf, ss)
		got := buf.String()
		expected := `3
2
1
Go!
`
		if got != expected {
			t.Fatalf("got %s, expected %s", got, expected)
		}
	})

	t.Run("write sleep testing test", func(t *testing.T) {
		ss := &SpySleeper{}
		Counter(buf, ss)
		expected := []string{
			"write",
			"sleep",
			"write",
			"sleep",
			"write",
			"sleep",
		}

		if reflect.DeepEqual(ss.Calls, expected) {
			t.Fatalf("got: %s, expected: %s", ss.Calls,
				expected)
		}
	})
}

package dependency_injection

import (
	"fmt"
	"io"
)

func Greet(w io.Writer, msg string) {
	fmt.Fprintf(w, msg)
}

type Sleeper interface {
	Sleep()
}

const timer = 3

func Counter(w io.Writer, sleeper Sleeper) {
	for i := timer; i > 0; i-- {
		_, err := fmt.Fprintln(w, i)
		if err != nil {
			fmt.Printf("unable to write to buffer: %d\n", i)
		}
		sleeper.Sleep()
	}
	_, _ = fmt.Fprintln(w, "Go!")
}

package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	// showcase error wrapping.
	err := BadReqGen()
	// Is passes through the error tree and identifies whether the first arg is equal to the second arg. It takes care of unwrapping the error.
	// equivalent to err == BadRequestError
	if errors.Is(err, BadRequest{}) {
		fmt.Printf("unwrapped error: %v\n", err)
		os.Exit(1)
	}
}

type BadRequest struct {
	reason string
}

func (b BadRequest) Error() string {
	return fmt.Sprintf("Hey its a bad request: %s", b.reason)
}

func BadReqGen() error {
	return BadRequest{reason: "Testing Wrapped Errors"}
}

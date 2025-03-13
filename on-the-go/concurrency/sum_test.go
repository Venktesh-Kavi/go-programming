package main

import "testing"

func TestSum(t *testing.T) {
	s := []int{3, 5, 1, 2}
	expected := 11

	ch := make(chan int, 2)

	go sum(s[:len(s)/2], ch)
	go sum(s[len(s)/2:], ch)

	x, y := <-ch, <-ch

	if x+y != expected {
		t.Errorf("expected sum: %d, received: %d\n", expected, x+y)
	}
}

package main

import "fmt"

func sum(s []int, ch chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	// produce to the channel
	ch <- sum
}

func main() {
	s := []int{3, 4, 5, 1, 6}

	ch := make(chan int, 1)

	go sum(s[:len(s)/2], ch)
	go sum(s[len(s)/2:], ch)

	x, y := <-ch, <-ch // consume from channel
	fmt.Println(x, y, x+y)
}

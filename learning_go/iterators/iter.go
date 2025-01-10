package main

import (
	"fmt"
	"iter"
	"slices"
)

func main() {
	// looping is generally done for slices, maps, strings and channels
	items := []Item{{name: "Foo", value: 1231.1231}, {name: "Bar", value: 1231.1231}}

	for _, item := range items {
		fmt.Println(item)
	}

	// vanilla iterator
	iterItems(func(item Item) bool {
		fmt.Println("Iterator: ", item)
		return true
	})

	// using function type iter.Seq
	for item := range ItemIterator(items) {
		fmt.Println("Iterator: ", item)
	}

	// using function type iter.Seq2
	for i, item := range ItemIteratorWithIndex(items) {
		fmt.Printf("Iterator at idx: %d, item: %v\n", i, item)
	}

	// in case of errors use iter.Seq2[data any, err error]

	// composing iterators
	PrintAll(ItemIterator(items))

	// slices and maps now uses iterators, slices.All, slices.Values
	for _, item := range slices.All(items) {
		fmt.Println("Slice Iterator: ", item)
	}
}

func PrintAll[V any](seq iter.Seq[V]) {
	for item := range seq {
		fmt.Println("Composed Iterator: ", item)
	}
}

// itermItems uses the yield func to do an operation on the item. It doesn't allow to range over an iterator as it require a function
func iterItems(yield func(Item) bool) {
	items := []Item{{name: "Foo", value: 1231.1231}, {name: "Bar", value: 1231.1231}}
	for _, item := range items {
		if !yield(item) {
			break
		}
	}
}

// ItemIterator Keep in mind that the core goal of an iterator itself is to produce one element at a time and defer the source data from being available right away.
func ItemIterator(items []Item) iter.Seq[Item] {
	return func(yield func(Item) bool) {
		for _, item := range items {
			if !yield(item) {
				break
			}
		}
	}
}

func ItemIteratorWithIndex(items []Item) iter.Seq2[int, Item] {
	return func(yield func(int, Item) bool) {
		for i, item := range items {
			if !yield(i, item) {
				break
			}
		}
	}
}

type Item struct {
	name  string
	value float64
}

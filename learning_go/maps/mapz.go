package main

import "fmt"

// DummyNonCompType cannot be used as map key. It has fields which are not comparable
type DummyNonCompType struct {
	data map[string]int
}

type DummyCompType struct {
	name string
	test int
}

func main() {
	m := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("This is a way to declare and initialise a map: ", m)

	m2 := make(map[string]int)
	m2["foo"] = 3
	fmt.Println("Added foo to a declared and initialized map: ", m2)

	var m3 map[string]int
	m3 = map[string]int{"foo": 1, "bar": 2}
	fmt.Println("Another way to declare and initialize map separately: ", m3)

	if v, found := m2["foo"]; found {
		fmt.Println("found foo: ", v)
	}

	m5 := map[DummyCompType]int{}
	m5[DummyCompType{"foo", 1}] = 1

	fmt.Println("map with a custom comparable type: ", m5)
}

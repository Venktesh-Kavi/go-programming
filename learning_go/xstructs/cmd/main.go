package main

import (
	"fmt"
	"xstruct/tstructs"
)

func main() {
	c := new(tstructs.Counter) // or var c Counter
	// Note: Able to call the pointer receiver method with a value type.
	c.Increment()
	fmt.Println(c.String())
	c.Increment()
	fmt.Println(c.String())

	var nc tstructs.Counter
	tstructs.DoUpdateWrong(nc)
	fmt.Println(nc.String())
	tstructs.DoUpdateRight(&nc)
	fmt.Println(nc.String()) // notice the pointer type being used to call the value type.
}

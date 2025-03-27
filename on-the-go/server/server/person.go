package server

type Gender int

const (
	MALE Gender = iota
	FEMALE
)

type Person struct {
	Name string
	Age  int
	Gender
}

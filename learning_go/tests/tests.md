## Notes on Go Testing

* [Learning Go With Tests](https://quii.gitbook.io/learn-go-with-tests)
* Table driven testing, involves passing multiple test cases dynamically to the test function. Ref:
  `hello_test.go`


### Common Commands

* go test
* go test -v
* go test -cover (verify coverage of code)


## Difference between Table and SubTests

Table Driven
* Table driven tests are a common way to write tests so that with minimal code repetition many scenarios can be covered.
* Syntax density comes with poor readability. Elegant to write often hard to read.
* Used in scenarios were there is clear pattern input X produces output Y.
* Table driven test scenarios need to be run by sub-test though.


Subtests
* Writing test scenarios as individual subtests is less common in GoLang.
* But it promotes high readability over code dryness.

This is a subtest.
``` go
t.Run("i am testing this", func(t *testing.T) {
	// some test code
})
```


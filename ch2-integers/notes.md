# Chapter 2 : Integers


Example functions begin with Example (much like test functions begin with Test), and reside in a package's _test.go files. Add the following ExampleAdd function to the adder_test.go file.

```go
func ExampleAdd() {
 sum := Add(1, 5)
 fmt.Println(sum)
 // Output: 6
}

```

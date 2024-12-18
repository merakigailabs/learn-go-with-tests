# Chapter 1 - Hello World

## Notes

When you write a program in Go, you will have a main package defined with a main func inside it. Packages are ways of grouping up related Go code together.

The func keyword defines a function with a name and a body.

With import "fmt" we are importing a package which contains the Println function that we use to print.

> Create a file called hello_test.go
> go test

go: cannot find main module, but found .git/config in /Users/merakigai/Desktop/Go/learn_go_test
        to create a module there, run:
        cd .. && go mod init

> go mod init example.com/hello

> go doc fmt
Will give use the documentation for the fmt package.

### Constants

Constants are defined like so:

```go
    const englishHelloPrefix = "Hello, "
```

### Multiple tests in a single file

```go
func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) { // Use t.Run to group tests together
		got := Hello("Chris")
		want := "Hello, Chris"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
```

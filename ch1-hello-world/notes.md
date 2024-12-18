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
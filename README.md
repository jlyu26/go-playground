# go-playground
Playing with Golang.

## Go CLI
- `go build` - Compiles a bunch of go source code files. main.go -> main(.exe in Windows), run with `./main`.
- `go run` - Compiles and executes files.
- `go fmt` - Formats all the code in each file in current directory.
- `go install` - Compiles and installs a package.
- `go get` - Downloads the raw source code of someone else's package.
- `go test` - Runs any tests associated with the current project.

## Types of Packages
- **Executable package:** Defines a package that can be compiled and then executed. `package main` -> `go build` -> `main.exe` Must have a function called `main`.
- **Reusable package:** Defines a package that can be used as dependency (helper code). Good place to put reusable logic.

## `import` Statement
[golang.org/pkg](https://golang.org/pkg/)
`fmt`: Print like a JavaScript `console`.

## `func`

## How to organize `main.go`
```go
// package declaration
package main
// import other packages that we need
import "fmt"
// declare functions, tell Go to do things
func main() {
	fmt.Println("hi there!")
}
```
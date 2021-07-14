/*
	ok so we're going to a run down of what exactly
	is going on in go through examples like this.

	with the go CLI we have
	go build - complies a bunch of go source code
	go run - compies and executes a specific file
	go fmt - formats all th code in each file in the current directory
	go install - complies and "installs" a package
	go get - downloads raw source of someone else's code.
	go test - runs any tests associated with the current project

	what does package mean?

	it's like a project or workspace
	it's important to correctly associate all files
		of a project together under the same package.

	types of packages:
		Executeable: a file that can be used or executed ./main
		Reuseable: helper files, good place for reusable logic or if you're making a depency package

	When you name a package "main" at the top, you specifying
	that the file is an executable.

	looking the name of the function
	func - keyword for the start of the function
	name of file - can be anything
	parenthese - contains parameters
	brackets - everything related to the

*/
package main

import "fmt"

func main() {
	fmt.Println("Hi there")
}

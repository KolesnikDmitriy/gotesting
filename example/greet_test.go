package greet

import "fmt"

func ExampleHello() {
	greeting := SayHello("Dima")

	fmt.Println(greeting)

	// Output:
	// Hello, Dima!
}

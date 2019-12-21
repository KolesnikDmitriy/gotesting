package example

import "fmt"

func ExampleSayHello() {
	greeting := SayHello("Dima")

	fmt.Println(greeting)
	// Output:
	// Hello, Dima!
}

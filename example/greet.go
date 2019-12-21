package greet

import "fmt"

// SayHello takes name and return hello message
func SayHello(name string) string {
	return fmt.Sprintf("Hello, %s!", name)
}

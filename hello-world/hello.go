package hello_world

import "fmt"

func Hello(name string) string {

	return "Hello, world"
}

func main() {
	fmt.Println(Hello("world"))
}

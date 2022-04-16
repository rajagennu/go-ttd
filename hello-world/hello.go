package main

import (
	"fmt"
)

func main() {
	fmt.Println(Hello("Raja"))
}

func Hello(name string) string {
	return "Hello, " + name + "!"
}

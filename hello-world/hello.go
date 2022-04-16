package main

import (
	"fmt"
)

const englishHelloPrefix = "Hello, "
const defaultSuffix = "World"
const frenchHelloPrefix = "Salut, "
const spanishHelloPrefix = "Hola, "
const spanish = "spanish"
const french = "french"

func main() {
	fmt.Println(Hello("Raja", "english"))
}

func Hello(name, lang string) (prefix string) {
	if name == "" {
		name = defaultSuffix
	}
	switch lang {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return prefix + name
}

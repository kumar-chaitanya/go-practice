package main

import "fmt"

const englishHelloPrefix string = "Hello "
const spanishHelloPrefix string = "Hola "
const suffix string = "!"

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	if language == "Spanish" {
		return spanishHelloPrefix + name + suffix
	}

	return englishHelloPrefix + name + suffix
}

func main() {
	fmt.Println(Hello("", "English"))
}
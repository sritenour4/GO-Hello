package main

import (
	"fmt"
	"hello/arrays_slices"
	"hello/integers"
	"hello/iteration"
)

const spanish = "Spanish"
const french = "French"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

// Hello returns a personalized greeting in a given langage
func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}
	return greetingPrefix(language) + name
}
	
func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("Shannon", "French"))
	fmt.Println(integers.Add(2, 2))
	fmt.Println(iteration.Repeat("a"))
	fmt.Println(iteration.CharRepeat("a", 12))

	numbers := []int{1,2,3,4,5}
	fmt.Println(arrays_slices.Sum(numbers))
	
}